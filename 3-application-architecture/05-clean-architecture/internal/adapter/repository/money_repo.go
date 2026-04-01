package repository

import (
	"clean/internal/entity"
	"errors"
)

type InMemoryAccountRepo struct {
	db map[int]*entity.Account
}

func NewInMemoryAccountRepo() *InMemoryAccountRepo {
	return &InMemoryAccountRepo{
		db: map[int]*entity.Account{
			1: {ID: 1, Balance: 1000},
			2: {ID: 2, Balance: 500},
		},
	}
}

func (r *InMemoryAccountRepo) GetByID(id int) (*entity.Account, error) {
	if a, ok := r.db[id]; ok {
		return a, nil
	}
	return nil, errors.New("account not found")
}

func (r *InMemoryAccountRepo) Update(a *entity.Account) error {
	r.db[a.ID] = a
	return nil
}
