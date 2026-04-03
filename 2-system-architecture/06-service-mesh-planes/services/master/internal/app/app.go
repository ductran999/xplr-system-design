package app

import (
	"fmt"
	"log"
	agentv1 "master-slave/api/gen/pb/agent/v1"
	"master-slave/services/master/internal/delivery/grpc/handler"
	"master-slave/services/master/internal/delivery/grpc/interceptor"
	"master-slave/services/master/internal/usecase"
	"net"

	"google.golang.org/grpc"
)

func Run() error {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		return fmt.Errorf("failed to listen: %w", err)
	}

	grpcServer := grpc.NewServer(
		grpc.UnaryInterceptor(interceptor.AuthInterceptor()),
	)

	// connManager := NewConnectionManager()
	registerUC := usecase.NewRegisterUC()
	hdl := handler.NewHandler(registerUC)
	agentv1.RegisterAgentServiceServer(grpcServer, hdl)

	log.Println("Control Plane listening on :50051")
	if err := grpcServer.Serve(lis); err != nil {
		return fmt.Errorf("failed to serve: %w", err)
	}

	return nil
}
