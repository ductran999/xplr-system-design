package filter

import "pipeline/model"

func ApplyTax(in <-chan model.Order) <-chan model.Order {
	out := make(chan model.Order)
	go func() {
		for order := range in {
			if order.IsValid {
				order.FinalAmount = order.FinalAmount * 1.1 // Thuế 10%
			}
			out <- order
		}
		close(out)
	}()

	return out
}
