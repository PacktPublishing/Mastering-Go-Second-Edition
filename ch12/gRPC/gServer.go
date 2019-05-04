package main

import (
	"fmt"
	//	"github.com/golang/protobuf/proto"
	"github.com/mactsouk/protobuf"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"net"
)

type MessageServer struct {
}

var port = ":8080"

func (MessageServer) SayIt(ctx context.Context, r *message_service.Request) (*message_service.Response, error) {
	fmt.Println("Request Text:", r.Text)
	fmt.Println("Request SubText:", r.Subtext)

	response := &message_service.Response{
		Text:    r.Text,
		Subtext: "Got it!",
	}

	return response, nil
}

func main() {

	server := grpc.NewServer()
	var messageServer MessageServer
	message_service.RegisterMessageServiceServer(server, messageServer)
	listen, err := net.Listen("tcp", port)
	if err != nil {
		fmt.Println(err)
		return
	}

	server.Serve(listen)
}
