package filter

import "pipeline/model"

func Validate(in <-chan model.Order) <-chan model.Order {
	out := make(chan model.Order)
	go func() {
		for order := range in {
			if order.RawAmount > 0 {
				order.IsValid = true
			}
			out <- order
		}
		close(out)
	}()

	return out
}
