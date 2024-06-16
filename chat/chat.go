package chat

import (
	"context"
	"log"
)

type Server struct {
	UnimplementedChatServiceServer
}

func (s Server) SayHello(ctx context.Context, in *Message) (*Message, error) {
	log.Printf("Receive message body rom client: %s", in.Body)
	return &Message{Body: "Hello From the server"}, nil
}
