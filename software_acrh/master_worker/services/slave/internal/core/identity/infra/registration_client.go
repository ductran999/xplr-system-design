package identityinfra

import (
	"context"
	agentv1 "master-slave/api/gen/pb/agent/v1"
	identity "master-slave/services/slave/internal/core/identity/entity"
)

type registrationClient struct {
	grpcClient agentv1.AgentServiceClient
}

func NewRegistrationClient(grpcClient agentv1.AgentServiceClient) *registrationClient {
	if grpcClient == nil {
		panic("infra: registration client requires grpc client")
	}

	return &registrationClient{
		grpcClient: grpcClient,
	}
}

func (rc *registrationClient) Register(ctx context.Context, agent identity.Agent) (*identity.Agent, error) {
	in := agentv1.RegisterRequest{
		RegistrationToken: agent.RegistrationToken,
		AgentVersion:      agent.Version,
		Metadata: &agentv1.AgentMetadata{
			Namespace:  agent.Metadata.Namespace,
			NodeName:   agent.Metadata.NodeName,
			PodName:    agent.Metadata.PodName,
			Hostname:   agent.Metadata.Hostname,
			K8SVersion: agent.Metadata.K8SVersion,
		},
	}

	resp, err := rc.grpcClient.Register(ctx, &in)
	if err != nil {
		return nil, err
	}

	agent.AccessKey = resp.AgentIdentity
	agent.ClusterID = resp.ClusterId

	return &agent, nil
}
