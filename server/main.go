package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/raihan-faza/go-grpc-practice/data"
	"github.com/raihan-faza/go-grpc-practice/userpb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type server struct {
	userpb.UnimplementedUserServiceServer
}

func (s *server) GetUser(ctx context.Context, req *userpb.GetUserRequest) (*userpb.GetUserResponse, error) {
	user, ok := data.MappedDummyData[int(req.UserID)]
	if !ok {
		return nil, fmt.Errorf("user doesnt exist.")
	}
	return &userpb.GetUserResponse{
		UserID:   user.UserID,
		UserName: user.UserName,
		UserAge:  user.UserAge,
	}, nil
}

func main() {
	for _, user := range data.DummyData {
		data.MappedDummyData[int(user.UserID)] = user
	}
	listener, err := net.Listen("tcp", ":1337")
	if err != nil {
		panic(err)
	}
	grpcServer := grpc.NewServer()
	userpb.RegisterUserServiceServer(grpcServer, &server{})
	reflection.Register(grpcServer)
	fmt.Print("grpc running on port 1337")
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
