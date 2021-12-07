package main

import (
	"fmt"
	"io"
	"net"
	"os"
	"sync"
	"time"

	"google.golang.org/grpc"
	"grpc.com/grpc/protobuff"
)

type serve struct {
	protobuff.UnimplementedStreamingServer
}

func (se *serve) ServerSideStreaming(in *protobuff.Input, res protobuff.Streaming_ServerSideStreamingServer) error {
	var wg sync.WaitGroup
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func() {
			result := protobuff.Output{Result: fmt.Sprintf("Hi %v", in.Name)}
			res.Send(&result)
			wg.Done()
		}()
		fmt.Println(i + 1)
		time.Sleep(1 * time.Second)
	}
	wg.Wait()
	return nil
}
func (se *serve) ClientSideStreaming(in protobuff.Streaming_ClientSideStreamingServer) error {
	var sum int32 = 0
	for {
		input, err := in.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}
		fmt.Println(input)
		sum += input.Num
		time.Sleep(1 * time.Second)
	}

	in.SendAndClose(&protobuff.AddOutput{Sum: sum})
	return nil
}
func main() {
	s := grpc.NewServer()
	protobuff.RegisterStreamingServer(s, &serve{})
	list, err := net.Listen("tcp", ":8081")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	if err := s.Serve(list); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

}
