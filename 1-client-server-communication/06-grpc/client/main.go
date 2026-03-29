package main

import (
	"context"
	userv1 "grpc-demo/gen/pb/user/v1"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Connect service A: %v", err)
	}
	defer conn.Close()

	client := userv1.NewUserServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	log.Println("Microservice B call Microservice A...")

	response, err := client.GetUser(ctx, &userv1.GetUserRequest{Id: 1})
	if err != nil {
		log.Fatalf("Call grpc failed: %v", err)
	}

	log.Printf("Result: ID=%d, Name=%s, Role=%s", response.GetId(), response.GetName(), response.GetRole())
}
