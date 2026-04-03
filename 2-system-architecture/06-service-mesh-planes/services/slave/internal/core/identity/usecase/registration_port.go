package identityuc

import (
	"context"
	identity "master-slave/services/slave/internal/core/identity/entity"
)

type RegistrationGateway interface {
	Register(ctx context.Context, agent identity.Agent) (*identity.Agent, error)
}

type RuntimeInfoProvider interface {
	GetMetadata(ctx context.Context) (identity.AgentMetadata, error)
}
