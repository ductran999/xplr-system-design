package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type Task struct {
	ID     string `json:"id"`
	Status string `json:"status"`
}

func main() {
	router := gin.Default()
	router.Use(cors.Default())

	statuses := []string{"pending", "in progress", "review", "done"}
	tasks := make(map[string]*Task)

	router.GET("/tasks/:id", func(c *gin.Context) {
		id := c.Param("id")

		if _, ok := tasks[id]; !ok {
			tasks[id] = &Task{
				ID:     id,
				Status: statuses[0],
			}
		} else {
			current := tasks[id].Status

			for i, s := range statuses {
				if s == current && i < len(statuses)-1 {
					tasks[id].Status = statuses[i+1]
					break
				}
			}
		}

		c.JSON(200, gin.H{
			"data": tasks[id],
		})
	})

	_ = router.Run("localhost:8080")
}
