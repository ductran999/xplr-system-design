package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type Message struct {
	Status    string `json:"status"`
	Text      string `json:"message"`
	Timestamp string `json:"timestamp"`
}

type Broker struct {
	clients        map[chan Message]bool // store clients waiting
	newClients     chan chan Message     // new client register via this channel
	defunctClients chan chan Message     // channel receive signal when client timeout
	messages       chan Message          //
}

func NewBroker() *Broker {
	return &Broker{
		clients:        make(map[chan Message]bool),
		newClients:     make(chan chan Message),
		defunctClients: make(chan chan Message),
		messages:       make(chan Message),
	}
}

func (b *Broker) Start() {
	for {
		select {
		case s := <-b.newClients:
			b.clients[s] = true
			log.Printf("[Broker] Add new 1 Client. Total waiting client: %d\n", len(b.clients))

		case s := <-b.defunctClients:
			delete(b.clients, s)
			close(s)
			log.Printf("[Broker] 1 Client disconnected. Total waiting client: %d\n", len(b.clients))

		//
		case msg := <-b.messages:
			for clientMessageChan := range b.clients {
				clientMessageChan <- msg
			}
		}
	}
}

func main() {
	broker := NewBroker()

	go broker.Start()

	router := gin.Default()
	router.Use(cors.Default())
	router.GET("/messages", func(c *gin.Context) {
		messageChan := make(chan Message)
		broker.newClients <- messageChan

		defer func() {
			broker.defunctClients <- messageChan
		}()

		select {
		case <-c.Done():
			return // User cancel request

		case msg := <-messageChan:
			c.JSON(http.StatusOK, gin.H{
				"data": msg,
			})
		}
	})

	router.POST("/messages", func(ctx *gin.Context) {
		var body struct {
			Text string `json:"text"`
		}

		if err := ctx.ShouldBindBodyWithJSON(&body); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"code":    "bad_request",
				"message": "invalid body",
			})
			return
		}

		log.Println("[Server] receive message:", body.Text)

		// Send message to broadcast channel
		msg := Message{
			Status:    "success",
			Text:      body.Text,
			Timestamp: time.Now().Format("15:04:05"),
		}
		broker.messages <- msg
		ctx.JSON(http.StatusOK, gin.H{
			"message": "send message to broadcast channel successfully!",
		})
	})

	_ = router.Run("localhost:8080")
}
