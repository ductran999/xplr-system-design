package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type User struct {
	Name    string
	Address string
}

func main() {
	users := []User{
		{
			Name:    "Daniel",
			Address: "Ha Noi",
		},
		{
			Name:    "Jenny",
			Address: "Sai Gon",
		},
	}

	router := gin.Default()
	router.Use(cors.Default())

	router.GET("/users", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"data": users,
		})
	})

	router.Run("localhost:8080")
}
