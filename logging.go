// Copyright (c) 2016 Wothing Co., Ltd. All rights reserved.

package worpc

import (
	"time"

	"github.com/wothing/log"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

func Logging(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
	tid := GetTidFromContext(ctx)
	start := time.Now()

	// pre-logging
	log.Tinfof(tid, "calling %s, req=%s", info.FullMethod, marshal(req))

	resp, err = handler(ctx, req)

	// post-logging
	log.Tinfof(tid, "finished %s, took=%v, resp=%s, err=%v", info.FullMethod, time.Since(start), marshal(resp), err)

	return resp, err
}
