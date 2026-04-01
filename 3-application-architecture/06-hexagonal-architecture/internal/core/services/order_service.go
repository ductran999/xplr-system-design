package services

import (
	"hexagonal/internal/core/domain"
	"hexagonal/internal/core/ports"
	"log"
)

type orderService struct {
	repo ports.OrderRepository
}

func NewOrderService(r ports.OrderRepository) ports.OrderService {
	return &orderService{repo: r}
}

func (s *orderService) CreateOrder(amount float64) error {
	order := domain.Order{Amount: amount}
	log.Printf("[Core Service] placing order %.2f...", amount)

	return s.repo.Save(order)
}
