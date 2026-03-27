package usecase

import "context"

type MonitoringUseCase interface {
	SendHeartbeat(ctx context.Context) error
}

type monitoringUC struct{}

func NewMonitoringUseCase() *monitoringUC {
	return &monitoringUC{}
}

func (uc *monitoringUC) SendHeartbeat(ctx context.Context) error {
	return nil
}
