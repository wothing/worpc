// Copyright (c) 2016 Wothing Co., Ltd. All rights reserved.

package worpc

import (
	"bytes"
	"fmt"
	"reflect"

	"github.com/golang/protobuf/jsonpb"
	"github.com/golang/protobuf/proto"
	"golang.org/x/net/context"
	"google.golang.org/grpc/metadata"
)

func GetTidFromContext(ctx context.Context) string {
	if md, ok := metadata.FromContext(ctx); ok {
		if md["tid"] != nil && len(md["tid"]) > 0 {
			return md["tid"][0]
		}
	}
	return "00000000-0000-0000-0000-000000000000"
}

var (
	js = &jsonpb.Marshaler{EnumsAsInts: true, EmitDefaults: true, OrigName: true}
)

// MarshalToString converts a protocol buffer object to JSON string.
func marshal(x interface{}) string {
	v := reflect.ValueOf(x)
	if !v.IsValid() || v.IsNil() {
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
