// Copyright (c) 2016 Wothing Co., Ltd. All rights reserved.

package worpc

import (
	"runtime"

	"github.com/wothing/log"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
)

const (
	MAXSTACKSIZE = 4096
)

func Recovery(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
	tid := GetTidFromContext(ctx)

	// recovery func
	defer func() {
		if r := recover(); r != nil {
			// log stack
			stack := make([]byte, MAXSTACKSIZE)
			stack = stack[:runtime.Stack(stack, false)]
			log.Terrorf(tid, "panic grpc invoke: %s, err=%v, stack:\n%s", info.FullMethod, r, string(stack))

			// if panic, set error to err, in order that client and sense it.
			err = grpc.Errorf(codes.Unknown, "internal panic error: %v", r)
		}
	}()

	return handler(ctx, req)
}
