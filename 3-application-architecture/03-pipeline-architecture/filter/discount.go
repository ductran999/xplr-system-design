package filter

import "pipeline/model"

func ApplyDiscount(in <-chan model.Order) <-chan model.Order {
	out := make(chan model.Order)
	go func() {
		for order := range in {
			if order.IsValid && order.RawAmount > 100 {
				order.FinalAmount = order.RawAmount * 0.9 // Giảm 10%
			} else {
				order.FinalAmount = order.RawAmount
			}
			out <- order
		}
		close(out)
	}()

	return out
}
