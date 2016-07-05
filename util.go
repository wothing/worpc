// Copyright (c) 2016 Wothing Co., Ltd. All rights reserved.

package worpc

import (
	"bytes"
	"fmt"
	"reflect"

	"github.com/golang/protobuf/jsonpb"
	"github.com/golang/protobuf/proto"
)

var (
	js = &jsonpb.Marshaler{EnumsAsInts: true, EmitDefaults: true, OrigName: true}
)

// MarshalToString converts a protocol buffer object to JSON string.
func marshal(x interface{}) string {
	if x == nil || reflect.ValueOf(x).IsNil() {
		return fmt.Sprintf("<nil>")
	}

	pb, ok := x.(proto.Message)
	if !ok {
		return fmt.Sprintf("Marshal to json error: not a proto message")
	}

	var buf bytes.Buffer
	if err := js.Marshal(&buf, pb); err != nil {
		return fmt.Sprintf("Marshal to json error: %s", err.Error())
	}
	return buf.String()
}
