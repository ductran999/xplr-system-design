package ports

import "hexagonal/internal/core/domain"

// Driven Port (Output): Core use this to interact with external system
type OrderRepository interface {
	Save(order domain.Order) error
}

// Driving Port (Input): Outer will call this
type OrderService interface {
	CreateOrder(amount float64) error
}
