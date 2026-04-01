package main

import (
	"hexagonal/internal/adapter/input"
	"hexagonal/internal/adapter/output"
	"hexagonal/internal/core/services"
	"log"
	"net/http"
)

func main() {
	repo := output.InMemoryRepo{}
	svc := services.NewOrderService(&repo)
	hdl := input.NewOrderHandler(svc)

	http.HandleFunc("/order", hdl.PostOrder)

	log.Println("Server is running on :8080")
	log.Println("curl http://localhost:8080/order?amount=150.5")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}
