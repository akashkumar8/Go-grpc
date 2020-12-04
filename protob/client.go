package main

import (
	"context"
	"log"

	"google.golang.org/grpc"

	communicate "github.com/akashkumar8/Grpc"
)

func main() {
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(":9000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connect: %s", err)
	}

	defer conn.Close()

	c := communicate.NewCommunicateServiceClient(conn)

	message := communicate.Message{
		Body: "Hello From the client!",
	}

	response, err := c.HeyyFolks(context.Background(), &message)
	if err != nil {
		log.Fatalf("Error when calling HeyyFolks: %s", err)
	}

	log.Printf("Response From Server: %s", response.Body)

}
