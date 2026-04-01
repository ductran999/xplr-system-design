package plugin

import "fmt"

type StripePlugin struct{}

func (s *StripePlugin) Name() string { return "Stripe" }

func (s *StripePlugin) Process(amount float64) bool {
	fmt.Printf("[Plugin Stripe]: Processing %.2f$ using credit card ...\n", amount)
	return true
}
