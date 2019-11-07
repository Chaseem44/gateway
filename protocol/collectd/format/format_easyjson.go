// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package collectdformat

import (
	json "encoding/json"

	easyjson "github.com/mailru/easyjson"
	jlexer "github.com/mailru/easyjson/jlexer"
	jwriter "github.com/mailru/easyjson/jwriter"
)

// suppress unused package warning
var (
	_ *json.RawMessage
	_ *jlexer.Lexer
	_ *jwriter.Writer
	_ easyjson.Marshaler
)

func easyjson72863a49DecodeGithubComSignalfxMetricproxyProtocolCollectdFormat(in *jlexer.Lexer, out *JSONWriteFormat) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "dsnames":
			if in.IsNull() {
				in.Skip()
				out.Dsnames = nil
			} else {
				in.Delim('[')
				if out.Dsnames == nil {
					if !in.IsDelim(']') {
						out.Dsnames = make([]*string, 0, 8)
					} else {
						out.Dsnames = []*string{}
					}
				} else {
					out.Dsnames = (out.Dsnames)[:0]
				}
				for !in.IsDelim(']') {
					var v1 *string
					if in.IsNull() {
						in.Skip()
						v1 = nil
					} else {
						if v1 == nil {
							v1 = new(string)
						}
						*v1 = string(in.String())
					}
					out.Dsnames = append(out.Dsnames, v1)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "dstypes":
			if in.IsNull() {
				in.Skip()
				out.Dstypes = nil
			} else {
				in.Delim('[')
				if out.Dstypes == nil {
					if !in.IsDelim(']') {
						out.Dstypes = make([]*string, 0, 8)
					} else {
						out.Dstypes = []*string{}
					}
				} else {
					out.Dstypes = (out.Dstypes)[:0]
				}
				for !in.IsDelim(']') {
					var v2 *string
					if in.IsNull() {
						in.Skip()
						v2 = nil
					} else {
						if v2 == nil {
							v2 = new(string)
						}
						*v2 = string(in.String())
					}
					out.Dstypes = append(out.Dstypes, v2)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "host":
			if in.IsNull() {
				in.Skip()
				out.Host = nil
			} else {
				if out.Host == nil {
					out.Host = new(string)
				}
				*out.Host = string(in.String())
			}
		case "interval":
			if in.IsNull() {
				in.Skip()
				out.Interval = nil
			} else {
				if out.Interval == nil {
					out.Interval = new(float64)
				}
				*out.Interval = float64(in.Float64())
			}
		case "plugin":
			if in.IsNull() {
				in.Skip()
				out.Plugin = nil
			} else {
				if out.Plugin == nil {
					out.Plugin = new(string)
				}
				*out.Plugin = string(in.String())
			}
		case "plugin_instance":
			if in.IsNull() {
				in.Skip()
				out.PluginInstance = nil
			} else {
				if out.PluginInstance == nil {
					out.PluginInstance = new(string)
				}
				*out.PluginInstance = string(in.String())
			}
		case "time":
			if in.IsNull() {
				in.Skip()
				out.Time = nil
			} else {
				if out.Time == nil {
					out.Time = new(float64)
				}
				*out.Time = float64(in.Float64())
			}
		case "type":
			if in.IsNull() {
				in.Skip()
				out.TypeS = nil
			} else {
				if out.TypeS == nil {
					out.TypeS = new(string)
				}
				*out.TypeS = string(in.String())
			}
		case "type_instance":
			if in.IsNull() {
				in.Skip()
				out.TypeInstance = nil
			} else {
				if out.TypeInstance == nil {
					out.TypeInstance = new(string)
				}
				*out.TypeInstance = string(in.String())
			}
		case "values":
			if in.IsNull() {
				in.Skip()
				out.Values = nil
			} else {
				in.Delim('[')
				if out.Values == nil {
					if !in.IsDelim(']') {
						out.Values = make([]*float64, 0, 8)
					} else {
						out.Values = []*float64{}
					}
				} else {
					out.Values = (out.Values)[:0]
				}
				for !in.IsDelim(']') {
					var v3 *float64
					if in.IsNull() {
						in.Skip()
						v3 = nil
					} else {
						if v3 == nil {
							v3 = new(float64)
						}
						*v3 = float64(in.Float64())
					}
					out.Values = append(out.Values, v3)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "message":
			if in.IsNull() {
				in.Skip()
				out.Message = nil
			} else {
				if out.Message == nil {
					out.Message = new(string)
				}
				*out.Message = string(in.String())
			}
		case "meta":
			if in.IsNull() {
				in.Skip()
			} else {
				in.Delim('{')
				if !in.IsDelim('}') {
					out.Meta = make(map[string]interface{})
				} else {
					out.Meta = nil
				}
				for !in.IsDelim('}') {
					key := string(in.String())
					in.WantColon()
					var v4 interface{}
					if m, ok := v4.(easyjson.Unmarshaler); ok {
						m.UnmarshalEasyJSON(in)
					} else if m, ok := v4.(json.Unmarshaler); ok {
						_ = m.UnmarshalJSON(in.Raw())
					} else {
						v4 = in.Interface()
					}
					(out.Meta)[key] = v4
					in.WantComma()
				}
				in.Delim('}')
			}
		case "severity":
			if in.IsNull() {
				in.Skip()
				out.Severity = nil
			} else {
				if out.Severity == nil {
					out.Severity = new(string)
				}
				*out.Severity = string(in.String())
			}
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjson72863a49EncodeGithubComSignalfxMetricproxyProtocolCollectdFormat(out *jwriter.Writer, in JSONWriteFormat) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"dsnames\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		if in.Dsnames == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v5, v6 := range in.Dsnames {
				if v5 > 0 {
					out.RawByte(',')
				}
				if v6 == nil {
					out.RawString("null")
				} else {
					out.String(string(*v6))
				}
			}
			out.RawByte(']')
		}
	}
	{
		const prefix string = ",\"dstypes\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		if in.Dstypes == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v7, v8 := range in.Dstypes {
				if v7 > 0 {
					out.RawByte(',')
				}
				if v8 == nil {
					out.RawString("null")
				} else {
					out.String(string(*v8))
				}
			}
			out.RawByte(']')
		}
	}
	{
		const prefix string = ",\"host\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		if in.Host == nil {
			out.RawString("null")
		} else {
			out.String(string(*in.Host))
		}
	}
	{
		const prefix string = ",\"interval\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		if in.Interval == nil {
			out.RawString("null")
		} else {
			out.Float64(float64(*in.Interval))
		}
	}
	{
		const prefix string = ",\"plugin\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		if in.Plugin == nil {
			out.RawString("null")
		} else {
			out.String(string(*in.Plugin))
		}
	}
	{
		const prefix string = ",\"plugin_instance\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		if in.PluginInstance == nil {
			out.RawString("null")
		} else {
			out.String(string(*in.PluginInstance))
		}
	}
	{
		const prefix string = ",\"time\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		if in.Time == nil {
			out.RawString("null")
		} else {
			out.Float64(float64(*in.Time))
		}
	}
	{
		const prefix string = ",\"type\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		if in.TypeS == nil {
			out.RawString("null")
		} else {
			out.String(string(*in.TypeS))
		}
	}
	{
		const prefix string = ",\"type_instance\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		if in.TypeInstance == nil {
			out.RawString("null")
		} else {
			out.String(string(*in.TypeInstance))
		}
	}
	{
		const prefix string = ",\"values\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		if in.Values == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v9, v10 := range in.Values {
				if v9 > 0 {
					out.RawByte(',')
				}
				if v10 == nil {
					out.RawString("null")
				} else {
					out.Float64(float64(*v10))
				}
			}
			out.RawByte(']')
		}
	}
	{
		const prefix string = ",\"message\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		if in.Message == nil {
			out.RawString("null")
		} else {
			out.String(string(*in.Message))
		}
	}
	{
		const prefix string = ",\"meta\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		if in.Meta == nil && (out.Flags&jwriter.NilMapAsEmpty) == 0 {
			out.RawString(`null`)
		} else {
			out.RawByte('{')
			v11First := true
			for v11Name, v11Value := range in.Meta {
				if v11First {
					v11First = false
				} else {
					out.RawByte(',')
				}
				out.String(string(v11Name))
				out.RawByte(':')
				if m, ok := v11Value.(easyjson.Marshaler); ok {
					m.MarshalEasyJSON(out)
				} else if m, ok := v11Value.(json.Marshaler); ok {
					out.Raw(m.MarshalJSON())
				} else {
					out.Raw(json.Marshal(v11Value))
				}
			}
			out.RawByte('}')
		}
	}
	{
		const prefix string = ",\"severity\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		if in.Severity == nil {
			out.RawString("null")
		} else {
			out.String(string(*in.Severity))
		}
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v JSONWriteFormat) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson72863a49EncodeGithubComSignalfxMetricproxyProtocolCollectdFormat(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v JSONWriteFormat) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson72863a49EncodeGithubComSignalfxMetricproxyProtocolCollectdFormat(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *JSONWriteFormat) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson72863a49DecodeGithubComSignalfxMetricproxyProtocolCollectdFormat(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *JSONWriteFormat) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson72863a49DecodeGithubComSignalfxMetricproxyProtocolCollectdFormat(l, v)
}
func easyjson72863a49DecodeGithubComSignalfxMetricproxyProtocolCollectdFormat1(in *jlexer.Lexer, out *JSONWriteBody) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		in.Skip()
		*out = nil
	} else {
		in.Delim('[')
		if *out == nil {
			if !in.IsDelim(']') {
				*out = make(JSONWriteBody, 0, 8)
			} else {
				*out = JSONWriteBody{}
			}
		} else {
			*out = (*out)[:0]
		}
		for !in.IsDelim(']') {
			var v12 *JSONWriteFormat
			if in.IsNull() {
				in.Skip()
				v12 = nil
			} else {
				if v12 == nil {
					v12 = new(JSONWriteFormat)
				}
				(*v12).UnmarshalEasyJSON(in)
			}
			*out = append(*out, v12)
			in.WantComma()
		}
		in.Delim(']')
	}
	if isTopLevel {
		in.Consumed()
	}
}
func easyjson72863a49EncodeGithubComSignalfxMetricproxyProtocolCollectdFormat1(out *jwriter.Writer, in JSONWriteBody) {
	if in == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
		out.RawString("null")
	} else {
		out.RawByte('[')
		for v13, v14 := range in {
			if v13 > 0 {
				out.RawByte(',')
			}
			if v14 == nil {
				out.RawString("null")
			} else {
				(*v14).MarshalEasyJSON(out)
			}
		}
		out.RawByte(']')
	}
}

// MarshalJSON supports json.Marshaler interface
func (v JSONWriteBody) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson72863a49EncodeGithubComSignalfxMetricproxyProtocolCollectdFormat1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v JSONWriteBody) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson72863a49EncodeGithubComSignalfxMetricproxyProtocolCollectdFormat1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *JSONWriteBody) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson72863a49DecodeGithubComSignalfxMetricproxyProtocolCollectdFormat1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *JSONWriteBody) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson72863a49DecodeGithubComSignalfxMetricproxyProtocolCollectdFormat1(l, v)
}
