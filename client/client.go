package main

import (
	"context"
	"log"

	"github.com/yoshiyuki-140/RPC/chat"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	var conn *grpc.ClientConn
	// grpc.Dialは古いのでNewClientに書き換えた
	// conn, err := grpc.Dial(":9000", grpc.WithTransportCredentials(insecure.NewCredentials()))
	conn, err := grpc.NewClient(":9000", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}
	defer conn.Close()

	c := chat.NewChatServiceClient(conn)

	resp, err := c.SayHello(context.Background(), &chat.Message{
		Body: "Hello From Client!",
	})
	if err != nil {
		log.Fatalf("Error when calling SayHello: %s", err)
	}
	log.Printf("Responce from server: %s", resp.Body)
}
