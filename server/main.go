package main

import (
	"log"
	"net"

	"github.com/raihan-faza/go-grpc-practice/userpb"
	"google.golang.org/grpc"
)

type server struct {
	userpb.UnimplementedUserServiceServer
}

func (s *server) GetUser(userpb.GetUserRequest) userpb.GetUserResponse {
	return userpb.GetUserResponse{}
}
func main() {
	listener, err := net.Listen("tcp", ":1337")
	if err != nil {
		panic(err)
	}
	grpcServer := grpc.NewServer()
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
