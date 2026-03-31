package main

import (
	"fmt"
	"pipeline/filter"
	"pipeline/model"
)

func main() {
	// Mock data Source
	orders := []model.Order{
		{ID: 1, RawAmount: 50},
		{ID: 2, RawAmount: 200},
		{ID: 3, RawAmount: -10}, // Error Order
	}

	// Create channel input of pipeline
	inputChan := make(chan model.Order)

	// Push data from source
	go func() {
		for _, o := range orders {
			inputChan <- o
		}
		close(inputChan)
	}()

	// Data pipeline: Input -> Validate -> Discount -> Tax -> Output
	validChan := filter.Validate(inputChan)
	discountChan := filter.ApplyDiscount(validChan)
	finalChan := filter.ApplyTax(discountChan)

	// Sink
	fmt.Println("Result")
	for order := range finalChan {
		if order.IsValid {
			fmt.Printf("Order #%d: Final amount  = %.2f\n", order.ID, order.FinalAmount)
		} else {
			fmt.Printf("Order #%d: Ignore (Valid: %v)\n", order.ID, order.IsValid)
		}
	}
}
