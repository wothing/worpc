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

// Recovery interceptor to handle grpc panic
func Recovery(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
	// recovery func
	defer func() {
		if r := recover(); r != nil {
			// log stack
			stack := make([]byte, MAXSTACKSIZE)
			stack = stack[:runtime.Stack(stack, false)]
			log.CtxErrorf(ctx, "panic grpc invoke: %s, err=%v, stack:\n%s", info.FullMethod, r, string(stack))

			// if panic, set custom error to 'err', in order that client and sense it.
			err = grpc.Errorf(codes.Internal, "panic error: %v", r)
		}
	}()

	return handler(ctx, req)
}
