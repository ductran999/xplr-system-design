package main

import (
	"database/sql"
	"encoding/json"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
)

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func main() {
	db, _ := sql.Open("sqlite3", "./user.db")
	db.Exec("CREATE TABLE IF NOT EXISTS users (id INTEGER PRIMARY KEY, name TEXT)")
	db.Exec("INSERT OR IGNORE INTO users (id, name) VALUES (1, 'Alice Micro')")

	http.HandleFunc("/get-user-internal", func(w http.ResponseWriter, r *http.Request) {
		id := r.URL.Query().Get("id")
		var u User
		err := db.QueryRow("SELECT id, name FROM users WHERE id = ?", id).Scan(&u.ID, &u.Name)
		if err != nil {
			http.Error(w, "User Not Found", 404)
			return
		}
		json.NewEncoder(w).Encode(u)
	})

	println("User Service (Micro) running on :8081")
	http.ListenAndServe(":8081", nil)
}
