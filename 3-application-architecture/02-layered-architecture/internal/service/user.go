package service

import (
	"layered-arch/internal/domain"
	"strings"
)

type userService struct {
	repo domain.UserRepository
}

func NewUserService(r domain.UserRepository) domain.UserService {
	return &userService{repo: r}
}

func (s *userService) GetDisplayName(id int) (string, error) {
	user, err := s.repo.GetByID(id)
	if err != nil {
		return "", err
	}

	return "PRO USER: " + strings.ToUpper(user.Name), nil
}
