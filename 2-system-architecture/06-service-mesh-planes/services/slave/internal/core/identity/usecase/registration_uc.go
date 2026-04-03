package identityuc

import (
	"context"
	"master-slave/services/slave/internal/config"
	identity "master-slave/services/slave/internal/core/identity/entity"
)

type RegisterClusterUseCase interface {
	Execute(ctx context.Context) (*identity.Agent, error)
}

type registerClusterUC struct {
	cfg *config.Config

	registrationClient  RegistrationGateway
	runtimeInfoProvider RuntimeInfoProvider
}

func NewRegisterClusterUC(
	cfg *config.Config,
	registrationClient RegistrationGateway,
	runtimeInfoProvider RuntimeInfoProvider,
) RegisterClusterUseCase {
	if cfg == nil {
		panic("nil agent config")
	}
	if registrationClient == nil {
		panic("nil registration client")
	}
	if runtimeInfoProvider == nil {
		panic("nil runtime info provider")
	}

	return &registerClusterUC{
		cfg:                 cfg,
		registrationClient:  registrationClient,
		runtimeInfoProvider: runtimeInfoProvider,
	}
}

func (uc *registerClusterUC) Execute(ctx context.Context) (*identity.Agent, error) {
	meta, err := uc.runtimeInfoProvider.GetMetadata(ctx)
	if err != nil {
		return nil, err
	}

	agent := identity.Agent{
		RegistrationToken: uc.cfg.RegistrationToken,
		Version:           uc.cfg.Version,
		Metadata:          meta,
	}

	return uc.registrationClient.Register(ctx, agent)
}
