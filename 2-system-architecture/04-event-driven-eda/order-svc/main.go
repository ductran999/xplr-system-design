package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/nats-io/nats.go"
	"github.com/nats-io/nats.go/jetstream"
)

type Order struct {
	ID     string  `json:"id"`
	User   string  `json:"user"`
	Amount float64 `json:"amount"`
}

func main() {
	nc, _ := nats.Connect("nats://localhost:4222")
	js, _ := jetstream.New(nc)
	ctx := context.Background()

	js.CreateStream(ctx, jetstream.StreamConfig{
		Name:     "ORDERS",
		Subjects: []string{"orders.>"},
	})

	http.HandleFunc("/place-order", func(w http.ResponseWriter, r *http.Request) {
		order := Order{
			ID:     fmt.Sprintf("ORD-%d", time.Now().Unix()),
			User:   "Danny",
			Amount: 250.5,
		}

		fmt.Printf("[Order Service] Store %s in DB\n", order.ID)

		orderData, _ := json.Marshal(order)
		_, err := js.Publish(ctx, "orders.created", orderData)
		if err != nil {
			http.Error(w, "Error event", 500)
			return
		}

		fmt.Printf("[Order Service] -> Publish event: orders.created for order %s\n", order.ID)
		w.Write([]byte("Place order successfully! Email will be sent soon."))
	})

	log.Println("Order Service is running on :8081...")
	_ = http.ListenAndServe("localhost:8081", nil)
}
