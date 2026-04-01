package plugin

import "fmt"

type PayPalPlugin struct{}

func (p *PayPalPlugin) Name() string { return "PayPal" }
func (p *PayPalPlugin) Process(amount float64) bool {
	fmt.Printf("[Plugin PayPal]: Process %.2f$...\n", amount)
	return true
}
