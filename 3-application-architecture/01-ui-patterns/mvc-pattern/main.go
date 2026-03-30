package main

import (
	"mvc/controller"
	"net/http"
)

func main() {
	http.HandleFunc("/user", controller.GetUser)

	println("MVC Server running on :8080...")
	http.ListenAndServe(":8080", nil)
}
