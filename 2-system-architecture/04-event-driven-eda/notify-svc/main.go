package main

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/nats-io/nats.go"
	"github.com/nats-io/nats.go/jetstream"
)

type OrderEvent struct {
	ID   string `json:"id"`
	User string `json:"user"`
}

func main() {
	nc, _ := nats.Connect("nats://localhost:4222")
	js, _ := jetstream.New(nc)
	ctx := context.Background()

	stream, _ := js.Stream(ctx, "ORDERS")

	consumer, _ := stream.CreateOrUpdateConsumer(ctx, jetstream.ConsumerConfig{
		Durable:   "EmailNotificationApp",
		AckPolicy: jetstream.AckExplicitPolicy,
	})

	fmt.Println("[Notification Service] wait event orders.created...")

	// 4. Pull mode
	iter, _ := consumer.Messages()
	for {
		msg, err := iter.Next()
		if err != nil {
			break
		}

		var order OrderEvent
		json.Unmarshal(msg.Data(), &order)

		fmt.Printf("[Notification Service] <- Got order event %s. Sent Email to %s...\n", order.ID, order.User)

		// Xác nhận đã xử lý xong (Ack)
		msg.Ack()
		fmt.Println("[Notification Service] -> Ack after send email success.")
	}
}
