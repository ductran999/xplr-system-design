package main

import (
	"context"
	userv1 "grpc-demo/gen/pb/user/v1"
	"log"
	"net"

	"google.golang.org/grpc"
)

type server struct {
	userv1.UnimplementedUserServiceServer
}

func (sv *server) GetUser(ctx context.Context, req *userv1.GetUserRequest) (*userv1.GetUserResponse, error) {
	if req.GetId() == 1 {
		return &userv1.GetUserResponse{
			Id:   1,
			Name: "Alice",
			Role: "Admin",
		}, nil
	}

	return &userv1.GetUserResponse{
		Id:   req.GetId(),
		Name: "Guest",
		Role: "User",
	}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("open port failed: %v", err)
	}

	grpcServer := grpc.NewServer()
	userv1.RegisterUserServiceServer(grpcServer, &server{})

	log.Println("Microservice A (gRPC Server) is listening on 50051...")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Server down: %v", err)
	}
}
