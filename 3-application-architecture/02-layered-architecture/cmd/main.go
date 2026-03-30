package main

import (
	"layered-arch/internal/controller"
	"layered-arch/internal/repository"
	"layered-arch/internal/service"
	"log"
	"net/http"
)

func main() {
	repo := &repository.MySQLUserRepo{}
	svc := service.NewUserService(repo)
	handler := controller.NewUserHandler(svc)

	http.HandleFunc("/user", handler.GetUser)

	log.Println("Layered Arch Server running on :8080...")
	_ = http.ListenAndServe(":8080", nil)
}
