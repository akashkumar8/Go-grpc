package communicate

import (
	"context"
	"log"
)

type Server struct {
}

func (s *Server) HeyyFolks(ctx context.Context, message *Message) (*Message, error) {
	log.Printf("Received message body from client: %s", message.Body)
	return &Message{Body: "Hello From the Server!"}, nil
}
