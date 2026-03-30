package interactor

import (
	"errors"
	"viper/entity"
)

type UserInteractor struct{}

func (i *UserInteractor) FetchUserByID(id int) (entity.User, error) {
	if id == 1 {
		return entity.User{ID: 1, Name: "Alice VIPER"}, nil
	}

	return entity.User{}, errors.New("user not found")
}
