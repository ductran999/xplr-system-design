package kernel

import (
	"fmt"
	"microkernel/plugin"
)

type PaymentEngine struct {
	plugins map[string]plugin.PaymentPlugin
}

func NewPaymentEngine() *PaymentEngine {
	return &PaymentEngine{
		plugins: make(map[string]plugin.PaymentPlugin),
	}
}

func (e *PaymentEngine) RegisterPlugin(p plugin.PaymentPlugin) {
	e.plugins[p.Name()] = p
	fmt.Printf("[Kernel]: Install Plugin succesfully: %s\n", p.Name())
}

func (e *PaymentEngine) ExecutePayment(method string, amount float64) {
	p, exists := e.plugins[method]
	if !exists {
		fmt.Printf("[Kernel Error]: Method %s not install yet!\n", method)
		return
	}

	if p.Process(amount) {
		fmt.Printf("[Kernel]: Payment successfully via %s\n", method)
	}
}
