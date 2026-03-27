package main

import (
	"log"
	"master-slave/services/master/internal/app"
)

func main() {
	if err := app.Run(); err != nil {
		log.Fatalln("Server crashed:", err)
	}
}
