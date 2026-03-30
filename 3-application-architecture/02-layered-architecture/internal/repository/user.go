package repository

import (
	"errors"
	"layered-arch/internal/domain"
)

type MySQLUserRepo struct{}

func (r *MySQLUserRepo) GetByID(id int) (*domain.User, error) {
	if id == 1 {
		return &domain.User{ID: 1, Name: "Alice Layered"}, nil
	}

	return nil, errors.New("user not found")
}
