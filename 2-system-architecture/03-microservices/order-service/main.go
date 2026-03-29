package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type UserResponse struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func main() {
	http.HandleFunc("/create-order", func(w http.ResponseWriter, r *http.Request) {
		userID := r.URL.Query().Get("user_id")

		resp, err := http.Get("http://localhost:8081/get-user-internal?id=" + userID)
		if err != nil || resp.StatusCode != 200 {
			fmt.Fprintf(w, "Err: auth user from User Service!")
			return
		}

		var u UserResponse
		json.NewDecoder(resp.Body).Decode(&u)

		fmt.Fprintf(w, "Placed order successfully for User: %s", u.Name)
	})

	println("Order Service (Micro) running on :8082")
	http.ListenAndServe(":8082", nil)
}
