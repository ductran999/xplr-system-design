package main

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool { // By pass CORS for developing
		return true
	},
}

type Message struct {
	Username string `json:"username"`
	Text     string `json:"text"`
}

var clients = make(map[*websocket.Conn]bool) // Map store client connection
var broadcast = make(chan Message)           // Channel receive message from user

func main() {
	http.HandleFunc("/ws", handleConnections)

	go handleMessages()

	log.Println("Go WebSocket Server running ws://localhost:8080/ws")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handleConnections(w http.ResponseWriter, r *http.Request) {
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer ws.Close()

	clients[ws] = true
	log.Printf("New client join! Total: %d\n", len(clients))

	// Infinity loop to get client message
	for {
		var msg Message
		err := ws.ReadJSON(&msg)
		if err != nil {
			log.Printf("Read message failed (User left): %v", err)
			delete(clients, ws) // remove tracking user already left
			break
		}

		log.Printf("Receive message %s: %s\n", msg.Username, msg.Text)

		// send message to broadcast channel
		broadcast <- msg
	}
}

func handleMessages() {
	for {
		// Wait for broadcast message
		msg := <-broadcast

		for client := range clients {
			err := client.WriteJSON(msg)
			if err != nil {
				log.Printf("Send message error: %v", err)
				client.Close()
				delete(clients, client) // who got network remove them from map
			}
		}
	}
}
