// Copyright 2019 The go-language-server Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Code generated by GoJay. DO NOT EDIT.

package jsonrpc2

import (
	"github.com/francoispqt/gojay"
)

// UnmarshalJSONObject implements gojay's UnmarshalerJSONObject
func (v *ID) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	if k == "" {
		if v.Name != "" {
			return dec.String(&v.Name)
		}
		return dec.Int64(&v.Number)
	}
	return nil
}

// NKeys returns the number of keys to unmarshal
func (v *ID) NKeys() int { return 1 }

// MarshalJSONObject implements gojay's MarshalerJSONObject
func (v *ID) MarshalJSONObject(enc *gojay.Encoder) {
	if v.Name != "" {
		enc.StringKey("id", v.Name)
		return
	}
	enc.Int64Key("id", v.Number)
}

// IsNil returns wether the structure is nil value or not
func (id *ID) IsNil() bool { return id == nil }

// UnmarshalJSONObject implements gojay's UnmarshalerJSONObject
func (v *Request) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	switch k {
	case "jsonrpc":
		return dec.String(&v.JSONRPC)
	case "id":
		if v.ID == nil {
			v.ID = &ID{}
		}
		return dec.Decode(v.ID)
	case "method":
		return dec.String(&v.Method)
	case "params":
		if v.Params.EmbeddedJSON == nil {
			v.Params.EmbeddedJSON = &gojay.EmbeddedJSON{}
		}
		return dec.EmbeddedJSON(v.Params.EmbeddedJSON)
	}
	return nil
}

// NKeys returns the number of keys to unmarshal
func (v *Request) NKeys() int { return 4 }

// MarshalJSONObject implements gojay's MarshalerJSONObject
func (v *Request) MarshalJSONObject(enc *gojay.Encoder) {
	enc.StringKey("jsonrpc", v.JSONRPC)
	enc.Encode(v.ID)
	enc.StringKey("method", v.Method)
	enc.AddObjectKeyOmitEmpty("params", v.Params)
}

// IsNil returns wether the structure is nil value or not
func (v *Request) IsNil() bool { return v == nil }

// UnmarshalJSONObject implements gojay's UnmarshalerJSONObject
func (v *Response) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	switch k {
	case "jsonrpc":
		return dec.String(&v.JSONRPC)
	case "id":
		if v.ID == nil {
			v.ID = &ID{}
		}
		return dec.Decode(v.ID)
	case "error":
		if v.Error == nil {
			v.Error = &Error{}
		}
		return dec.Object(v.Error)
	case "result":
		if v.Result.EmbeddedJSON == nil {
			v.Result.EmbeddedJSON = &gojay.EmbeddedJSON{}
		}
		return dec.EmbeddedJSON(v.Result.EmbeddedJSON)
	}
	return nil
}

// NKeys returns the number of keys to unmarshal
func (v *Response) NKeys() int { return 4 }

// MarshalJSONObject implements gojay's MarshalerJSONObject
func (v *Response) MarshalJSONObject(enc *gojay.Encoder) {
	enc.StringKey("jsonrpc", v.JSONRPC)
	enc.Encode(v.ID)
	enc.ObjectKey("error", v.Error)
	enc.AddObjectKeyOmitEmpty("result", v.Result)
}

// IsNil returns wether the structure is nil value or not
func (v *Response) IsNil() bool { return v == nil }

// UnmarshalJSONObject implements gojay's UnmarshalerJSONObject
func (v *NotificationMessage) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	switch k {
	case "jsonrpc":
		return dec.String(&v.JSONRPC)
	case "method":
		return dec.String(&v.Method)
	case "params":
		if v.Params.EmbeddedJSON == nil {
			v.Params.EmbeddedJSON = &gojay.EmbeddedJSON{}
		}
		return dec.EmbeddedJSON(v.Params.EmbeddedJSON)
	}
	return nil
}

// NKeys returns the number of keys to unmarshal
func (v *NotificationMessage) NKeys() int { return 3 }

// MarshalJSONObject implements gojay's MarshalerJSONObject
func (v *NotificationMessage) MarshalJSONObject(enc *gojay.Encoder) {
	enc.StringKey("jsonrpc", v.JSONRPC)
	enc.StringKey("method", v.Method)
	enc.AddObjectKeyOmitEmpty("result", v.Params)
}

// IsNil returns wether the structure is nil value or not
func (v *NotificationMessage) IsNil() bool { return v == nil }

// IsNil returns wether the structure is nil value or not
func (v *RawMessage) IsNil() bool { return v == nil }

// UnmarshalJSONObject implements gojay's UnmarshalerJSONObject
func (v *RawMessage) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {
	if v.EmbeddedJSON == nil {
		v.EmbeddedJSON = &gojay.EmbeddedJSON{}
	}
	return dec.EmbeddedJSON(v.EmbeddedJSON)
}

// NKeys returns the number of keys to unmarshal
func (v *RawMessage) NKeys() int { return 1 }

// MarshalJSONObject implements gojay's MarshalerJSONObject
func (v *RawMessage) MarshalJSONObject(enc *gojay.Encoder) {
	enc.AddEmbeddedJSONKeyOmitEmpty("result", v.EmbeddedJSON)
}
