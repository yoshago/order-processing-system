package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/yoshago/order-processing-system/internal/api"
	"github.com/yoshago/order-processing-system/internal/db"
)

func main() {
	// Initialize database connection
	database, err := db.InitDB()
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// Initialize Gin router
	r := gin.Default()

	// Initialize API routes
	api.InitRoutes(r, database)

	// Start the server
	if err := r.Run(":8080"); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
