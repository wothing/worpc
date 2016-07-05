package main

import (
	"fmt"

	"github.com/wothing/worpc/example/pb"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

func main() {
	conn, err := grpc.Dial("127.0.0.1:1701", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}

	client := pb.NewHelloServiceClient(conn)

	{
		// ctx := metadata.NewContext(context.Background(), metadata.Pairs("tid", "normal-hello-request"))
		resp, err := client.NormalHello(context.Background(), &pb.HelloRequest{Greeting: "world"})
		fmt.Printf("normal hello: resp=%#v, error=%v\n", resp, err)
	}
	{
		// ctx := metadata.NewContext(context.Background(), metadata.Pairs("tid", "normal-err-request"))
		resp, err := client.NormalHello(context.Background(), &pb.HelloRequest{Greeting: "world"})
		fmt.Printf("err hello: resp=%#v, error=%v\n", resp, err)
	}
	{
		ctx := metadata.NewContext(context.Background(), metadata.Pairs("tid", "normal-panic-request"))
		resp, err := client.PanicHello(ctx, &pb.HelloRequest{Greeting: "world"})
		fmt.Printf("panic hello: resp=%#v, error=%v\n", resp, err)
	}
}
