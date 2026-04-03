package domain

import "errors"

var (
	ErrInvalidRegistrationKey = errors.New("invalid registration key")
)

type Agent struct {
	ClusterID   string
	AccessToken string
}
