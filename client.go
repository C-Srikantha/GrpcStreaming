package main

import (
	"context"
	"fmt"
	"io"
	"math/rand"
	"os"

	"google.golang.org/grpc"
	"grpc.com/grpc/protobuff"
)

type in struct{}

func main() {
	conn, err := grpc.Dial(":8081", grpc.WithInsecure())
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	c := protobuff.NewStreamingClient(conn)
	//Example for ServerSideStreaming
	res, err := c.ServerSideStreaming(context.Background(), &protobuff.Input{Name: "Srikantha"})
	if err != nil {
		fmt.Println(err)
		return
	}
	for {
		output, err := res.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(output.Result)
	}
	//Example for ClientSideStreaming
	res1, err := c.ClientSideStreaming(context.Background())
	for {
		i := rand.Intn(10)
		if i > 8 {
			break
		}
		res1.Send(&protobuff.InputNum{Num: int32(i)})
	}
	output, err := res1.CloseAndRecv()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(output.Sum)

}
