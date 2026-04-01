package output

import (
	"fmt"
	"hexagonal/internal/core/domain"
)

type InMemoryRepo struct{}

func (r *InMemoryRepo) Save(order domain.Order) error {
	fmt.Printf("[Adapter Output] Place %.2f order successfully !\n", order.Amount)
	return nil
}
