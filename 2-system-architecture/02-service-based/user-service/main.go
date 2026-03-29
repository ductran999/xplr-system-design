package main

import (
	"database/sql"
	"fmt"
	"log/slog"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, _ := sql.Open("sqlite3", "./shared.db")
	db.Exec("CREATE TABLE IF NOT EXISTS users (id INTEGER PRIMARY KEY, name TEXT)")
	http.HandleFunc("/get-user", func(w http.ResponseWriter, r *http.Request) {
		id := r.URL.Query().Get("id")
		var name string
		err := db.QueryRow("SELECT name FROM users WHERE id = ?", id).Scan(&name)
		if err != nil {
			fmt.Fprintf(w, "User not found")
			return
		}

		fmt.Fprintf(w, "User Name: %s", name)
	})

	slog.Info("User Service (SBA) on port: 8081...")
	_ = http.ListenAndServe("localhost:8081", nil)
}
