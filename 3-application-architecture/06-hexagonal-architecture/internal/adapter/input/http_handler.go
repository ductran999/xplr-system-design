package input

import (
	"hexagonal/internal/core/ports"
	"net/http"
	"strconv"
)

type OrderHandler struct {
	svc ports.OrderService
}

func NewOrderHandler(s ports.OrderService) *OrderHandler {
	return &OrderHandler{svc: s}
}

func (h *OrderHandler) PostOrder(w http.ResponseWriter, r *http.Request) {
	amount, _ := strconv.ParseFloat(r.URL.Query().Get("amount"), 64)
	h.svc.CreateOrder(amount)

	w.Write([]byte("Order Success!"))
}
