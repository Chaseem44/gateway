package collectd

import (
	"fmt"
	"net"
	"net/http"
	"sync/atomic"
	"time"

	"github.com/gorilla/mux"

	"strings"

	"encoding/json"

	"github.com/signalfx/golib/datapoint"
	"github.com/signalfx/golib/datapoint/dpsink"
	"github.com/signalfx/golib/errors"
	"github.com/signalfx/golib/event"
	"github.com/signalfx/golib/pointer"
	"github.com/signalfx/golib/sfxclient"
	"github.com/signalfx/golib/web"
	"github.com/signalfx/metricproxy/protocol"
	"golang.org/x/net/context"
)

// ListenerServer will listen for collectd datapoint connections
type ListenerServer struct {
	name     string
	listener net.Listener
	server   http.Server
	decoder  JSONDecoder
}

var _ protocol.Listener = &ListenerServer{}

// Close the socket currently open for collectd JSON connections
func (s *ListenerServer) Close() error {
	return s.listener.Close()
}

// Datapoints returns JSON decoder datapoints
func (s *ListenerServer) Datapoints() []*datapoint.Datapoint {
	return s.decoder.Datapoints()
}

// JSONDecoder can decode collectd's native JSON datapoint format
type JSONDecoder struct {
	SendTo      dpsink.Sink
	DefaultDims map[string]string

	TotalErrors    int64
	TotalBlankDims int64
}

const sfxDimQueryParamPrefix string = "sfxdim_"

// ServeHTTPC decodes datapoints for the connection and sends them to the decoder's sink
func (decoder *JSONDecoder) ServeHTTPC(ctx context.Context, rw http.ResponseWriter, req *http.Request) {
	err := decoder.Read(ctx, req)
	if err != nil {
		atomic.AddInt64(&decoder.TotalErrors, 1)
		rw.WriteHeader(http.StatusBadRequest)
		rw.Write([]byte(fmt.Sprintf("Unable to decode json: %s", err.Error())))
		return
	}
	rw.Write([]byte(`"OK"`))
}

func newDataPoints(f *JSONWriteFormat, defaultDims map[string]string) []*datapoint.Datapoint {
	dps := make([]*datapoint.Datapoint, 0, len(f.Dsnames))
	for i := range f.Dsnames {
		if i < len(f.Dstypes) && i < len(f.Values) && f.Values[i] != nil {
			dps = append(dps, NewDatapoint(f, uint(i), defaultDims))
		}
	}
	return dps
}

func newEvent(f *JSONWriteFormat, defaultDims map[string]string) *event.Event {
	if f.Time != nil && f.Severity != nil && f.Message != nil {
		return NewEvent(f, defaultDims)
	}
	return nil
}

func (decoder *JSONDecoder) Read(ctx context.Context, req *http.Request) error {
	defaultDims := decoder.defaultDims(req)
	var d JSONWriteBody
	err := json.NewDecoder(req.Body).Decode(&d)
	if err != nil {
		return err
	}
	es := make([]*event.Event, 0, len(d)*2)
	dps := make([]*datapoint.Datapoint, 0, len(d)*2)
	for _, f := range d {
		if e := newEvent(f, defaultDims); e == nil {
			dps = append(dps, newDataPoints(f, defaultDims)...)
		} else {
			es = append(es, e)
		}
	}

	var e1, e2 error
	if len(dps) > 0 {
		e1 = decoder.SendTo.AddDatapoints(ctx, dps)
	}
	if len(es) > 0 {
		e2 = decoder.SendTo.AddEvents(ctx, es)
	}
	return errors.NewMultiErr([]error{e1, e2})
}

func (decoder *JSONDecoder) defaultDims(req *http.Request) map[string]string {
	params := req.URL.Query()
	defaultDims := make(map[string]string, len(decoder.DefaultDims))
	for key := range params {
		if strings.HasPrefix(key, sfxDimQueryParamPrefix) {
			value := params.Get(key)
			if len(value) == 0 {
				atomic.AddInt64(&decoder.TotalBlankDims, 1)
				continue
			}
			key = key[len(sfxDimQueryParamPrefix):]
			defaultDims[key] = value
		}
	}
	for k, v := range decoder.DefaultDims {
		defaultDims[k] = v
	}
	return defaultDims
}

// Datapoints about this decoder, including how many datapoints it decoded
func (decoder *JSONDecoder) Datapoints() []*datapoint.Datapoint {
	return []*datapoint.Datapoint{
		sfxclient.Cumulative("total_blank_dims", nil, atomic.LoadInt64(&decoder.TotalBlankDims)),
		sfxclient.Cumulative("invalid_collectd_json", nil, atomic.LoadInt64(&decoder.TotalErrors)),
	}
}

// ListenerConfig controls optional parameters for collectd listeners
type ListenerConfig struct {
	Name              *string
	ListenAddr        *string
	ListenPath        *string
	Timeout           *time.Duration
	DefaultDimensions map[string]string
	StartingContext   context.Context
}

var defaultListenerConfig = &ListenerConfig{
	Name:            pointer.String("collectd"),
	ListenAddr:      pointer.String("127.0.0.1:8081"),
	ListenPath:      pointer.String("/post-collectd"),
	Timeout:         pointer.Duration(time.Second * 30),
	StartingContext: context.Background(),
}

// NewListener servers http collectd requests
func NewListener(sink dpsink.Sink, passedConf *ListenerConfig) (*ListenerServer, error) {
	conf := pointer.FillDefaultFrom(passedConf, defaultListenerConfig).(*ListenerConfig)

	listener, err := net.Listen("tcp", *conf.ListenAddr)
	if err != nil {
		return nil, err
	}

	r := mux.NewRouter()
	listenServer := ListenerServer{
		listener: listener,
		server: http.Server{
			Handler:      r,
			Addr:         listener.Addr().String(),
			ReadTimeout:  *conf.Timeout,
			WriteTimeout: *conf.Timeout,
		},
		decoder: JSONDecoder{
			SendTo:      sink,
			DefaultDims: passedConf.DefaultDimensions,
		},
	}
	httpHandler := web.NewHandler(conf.StartingContext, &listenServer.decoder)
	r.Path(*conf.ListenPath).Headers("Content-type", "application/json").Handler(httpHandler)

	go listenServer.server.Serve(listener)
	return &listenServer, nil
}
