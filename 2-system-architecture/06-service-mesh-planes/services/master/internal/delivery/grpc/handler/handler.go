package handler

import (
	"context"
	"errors"
	"log"
	agentv1 "master-slave/api/gen/pb/agent/v1"
	"master-slave/services/master/internal/delivery/grpc/interceptor"
	"master-slave/services/master/internal/domain"
	"master-slave/services/master/internal/usecase"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

type grpcHandler struct {
	agentv1.UnimplementedAgentServiceServer

	registerUC usecase.RegisterUseCase
}

func NewHandler(regUC usecase.RegisterUseCase) *grpcHandler {
	return &grpcHandler{
		registerUC: regUC,
	}
}

func (s *grpcHandler) Register(ctx context.Context, req *agentv1.RegisterRequest) (*agentv1.RegisterResponse, error) {
	res, err := s.registerUC.Register(ctx, req.RegistrationToken)
	if err != nil {
		if errors.Is(err, domain.ErrInvalidRegistrationKey) {
			return nil, status.Error(codes.Unauthenticated, err.Error())
		}
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &agentv1.RegisterResponse{
		ClusterId:     res.ClusterID,
		AgentIdentity: res.AccessToken,
		Message:       "Registration successful",
	}, nil
}

func (s *grpcHandler) SendHeartbeat(ctx context.Context, req *agentv1.SendHeartbeatRequest) (*agentv1.SendHeartbeatResponse, error) {
	token, ok := ctx.Value(interceptor.ContextTokenKey).(string)
	if !ok {
		return nil, status.Error(codes.Unauthenticated, "missing token in context")
	}

	log.Println("token:", token)

	return &agentv1.SendHeartbeatResponse{
		NextIntervalSeconds: 10,
	}, nil
}

func (s *grpcHandler) ConnectTunnel(stream agentv1.AgentService_ConnectTunnelServer) error {
	return nil
	// clusterID, _ := getClusterIDFromMetadata(stream.Context())

	// cmdChan := make(chan *agentv1.ConnectTunnelResponse, 100)
	// testCmd := &agentv1.ConnectTunnelResponse{
	// 	CommandId: "cmd-uuid-123",
	// 	Action:    "DEPLOY_MODEL",
	// }
	// cmdChan <- testCmd
	// s.connectionManager.Register(clusterID, cmdChan)
	// defer s.connectionManager.Unregister(clusterID)

	// go func() {
	// 	for cmd := range cmdChan {
	// 		if err := stream.Send(cmd); err != nil {
	// 			return
	// 		}
	// 	}
	// }()

	// for {
	// 	report, err := stream.Recv()
	// 	if err != nil {
	// 		return err
	// 	}
	// 	log.Printf("Received report from %s: %s", clusterID, report.Status)
	// }
}

func getClusterIDFromMetadata(ctx context.Context) (string, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return "", status.Errorf(codes.Unauthenticated, "metadata is not provided")
	}

	values := md.Get("x-cluster-id")
	if len(values) == 0 {
		return "", status.Errorf(codes.Unauthenticated, "cluster-id not found")
	}

	return values[0], nil
}
