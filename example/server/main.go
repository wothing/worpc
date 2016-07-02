package main

import (
	"flag"
	"fmt"
	"net"

	"github.com/wothing/log"
	"github.com/wothing/worpc"
	"github.com/wothing/worpc/example/pb"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
)

var (
	port = flag.Int("port", 1701, "listening port")
)

func main() {
	flag.Parse()

	lis, err := net.Listen("tcp", fmt.Sprintf("0.0.0.0:%d", *port))
	if err != nil {
		panic(err)
	}

	log.Printf("starting hello service at %d", *port)
	s := worpc.NewServer()
	pb.RegisterHelloServiceServer(s, &helloServer{})
	s.Serve(lis)
}

type helloServer struct {
}

func (helloServer) NormalHello(ctx context.Context, req *pb.HelloRequest) (*pb.HelloResponse, error) {
	tid := worpc.GetTidFromContext(ctx)

	log.Tinfof(tid, "in normal hello.")
	return &pb.HelloResponse{Reply: "Hello, " + req.Greeting}, nil
}

func (helloServer) PanicHello(ctx context.Context, req *pb.HelloRequest) (*pb.HelloResponse, error) {
	tid := worpc.GetTidFromContext(ctx)

	log.Tinfof(tid, "in panic hello.")
	panic(fmt.Errorf("nothing"))

	return &pb.HelloResponse{Reply: "Hello, " + req.Greeting}, nil
}

func (helloServer) ErrHello(ctx context.Context, req *pb.HelloRequest) (*pb.HelloResponse, error) {
	tid := worpc.GetTidFromContext(ctx)
	log.Tinfof(tid, "in error hello.")
	return nil, grpc.Errorf(codes.Canceled, "just try to error")
}
