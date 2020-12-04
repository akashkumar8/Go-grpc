package main 

import (
	"fmt"
	"net"

	"google.golang.org/grpc"
	"github.com/akashkumar8/Grpc"

)

func main() {
	lis, err := net.Listen("tcp", ":9000")
	if err := nil {
		log.Fatal("Failed to listen on port 9000: %v", err)
	}

	s := communicate.Server{}

	grpcServer := grpc.NewServer()

	communicate.RegisterCommunicateServiceServer(grpcServer, &s)
	if err := grpcServer.serve(lis); err != nil {
		log.Fatal("Failed to serve gRPC server over port 9000: %v",err)
	}
}