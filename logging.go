// Copyright (c) 2016 Wothing Co., Ltd. All rights reserved.

package worpc

import (
	"time"

	"github.com/wothing/log"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

// Logging interceptor for grpc
func Logging(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
	start := time.Now()

	log.CtxInfof(ctx, "calling %s, req=%s", info.FullMethod, marshal(req))
	resp, err = handler(ctx, req)
	log.CtxInfof(ctx, "finished %s, took=%v, resp=%v, err=%v", info.FullMethod, time.Since(start), resp, err)

	return resp, err
}
