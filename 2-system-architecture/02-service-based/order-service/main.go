package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, _ := sql.Open("sqlite3", "./shared.db")

	http.HandleFunc("/create-order", func(w http.ResponseWriter, r *http.Request) {
		userID := r.URL.Query().Get("user_id")

		var name string
		err := db.QueryRow("SELECT name FROM users WHERE id = ?", userID).Scan(&name)

		if err != nil {
			fmt.Fprintf(w, "Error: user %s not found in shared database!", userID)
			return
		}

		fmt.Fprintf(w, "Place order successfully!: %s", name)
	})

	log.Println("Order Service (SBA) listening on 8082...")

	_ = http.ListenAndServe("localhost:8082", nil)
}
