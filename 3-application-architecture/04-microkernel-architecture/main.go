package main

import (
	"fmt"
	"microkernel/kernel"
	"microkernel/plugin"
)

func main() {
	engine := kernel.NewPaymentEngine()

	stripe := &plugin.StripePlugin{}
	paypal := &plugin.PayPalPlugin{}

	engine.RegisterPlugin(stripe)
	engine.RegisterPlugin(paypal)

	fmt.Println("--- Start transaction ---")
	engine.ExecutePayment("Stripe", 150.0)
	engine.ExecutePayment("PayPal", 80.5)

	engine.ExecutePayment("Momo", 50.0)
}
