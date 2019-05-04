package main

import (
	"fmt"
	"github.com/mactsouk/protobuf"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

var port = ":8080"

func SayIt(ctx context.Context, m message_service.MessageServiceClient, text string) (*message_service.Response, error) {
	request := &message_service.Request{
		Text:    text,
		Subtext: "New Message!",
	}

	r, err := m.SayIt(ctx, request)

	if err != nil {
		return nil, err
	}

	return r, nil
}

func main() {
	conn, err := grpc.Dial(port, grpc.WithInsecure())
	if err != nil {
		fmt.Println("Dial:", err)
		return
	}

	client := message_service.NewMessageServiceClient(conn)

	r, err := SayIt(context.Background(), client, "My Message!")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Response Text:", r.Text)
	fmt.Println("Response SubText:", r.Subtext)
}
