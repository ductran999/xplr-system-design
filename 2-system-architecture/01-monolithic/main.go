package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func InventoryService_Deduct(productID int) bool {
	log.Printf("[Inventory Service] Checking stock for product %d...", productID)
	time.Sleep(500 * time.Millisecond) // Simulate DB query
	log.Printf("[Inventory Service] Stock deducted successfully!")
	return true
}

func PaymentService_Process(userID int, amount float64) bool {
	log.Printf("[Payment] Charging %.2f$ from User %d...", amount, userID)
	time.Sleep(500 * time.Millisecond) // Simulate bank API call
	log.Printf("[Payment] Payment successful!")
	return true
}

func NotificationService_SendEmail(userID int) {
	log.Printf("[Email] Sending receipt to User %d...", userID)
	time.Sleep(200 * time.Millisecond)
	log.Printf("[Email] Email sent successfully!")
}

func handleCheckout(c *gin.Context) {
	userID := 1
	productID := 101
	price := 50.0

	log.Println("--- STARTING CHECKOUT PROCESS ---")

	// Tight Coupling: Direct function calls in memory
	// If Inventory fails -> the whole flow breaks
	if !InventoryService_Deduct(productID) {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Out of stock",
		})
		return
	}

	// If Payment fails -> return error immediately
	if !PaymentService_Process(userID, price) {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Payment failed",
		})
		return
	}

	// Send email
	NotificationService_SendEmail(userID)
	log.Println("--- CHECKOUT SUCCESSFUL ---")
	c.JSON(http.StatusOK, gin.H{
		"message": "Your order has been placed successfully!",
	})
}

func main() {
	router := gin.Default()
	router.Use(cors.Default())
	router.GET("/checkout", handleCheckout)

	log.Fatal(router.Run("localhost:8080"))
}
