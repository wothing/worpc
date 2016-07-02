// Copyright (c) 2016 Wothing Co., Ltd. All rights reserved.

// grpc interceptor chain builder & middlewares. Now, we are only using unary rpc,
// this package only support unary interceptor.
package worpc

import (
	"google.golang.org/grpc"
)

func NewServer() *grpc.Server {
	return grpc.NewServer(grpc.UnaryInterceptor(UnaryInterceptorChain(Recovery, Logging)))
}
