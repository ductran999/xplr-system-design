package usecase

import (
	"context"
	"master-slave/services/master/internal/domain"
)

type RegisterUseCase interface {
	Register(ctx context.Context, key string) (*domain.Agent, error)
}

type registerUC struct{}

func NewRegisterUC() *registerUC {
	return &registerUC{}
}

func (uc *registerUC) Register(ctx context.Context, key string) (*domain.Agent, error) {
	if key != "OK" {
		return nil, domain.ErrInvalidRegistrationKey
	}

	return &domain.Agent{
		ClusterID:   "cluster-1",
		AccessToken: "access_key_ne",
	}, nil
}
