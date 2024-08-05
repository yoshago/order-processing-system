package main

import (
	"log"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/yoshago/order-processing-system/internal/api"
	"github.com/yoshago/order-processing-system/internal/db"
	"github.com/yoshago/order-processing-system/internal/processor"
)

func main() {
	// Initialize database connection
	database, err := db.InitDB()
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	//Get worker count from env
	workerCountStr := os.Getenv("WORKER_COUNT")
	workerCount, err := strconv.Atoi(workerCountStr)
	if err != nil {
		log.Fatalf("Invalid WORKER_COUNT value: %v", err)
	}

	// Initialize worker pool
	workerPool := processor.NewWorkerPool(database, workerCount) // Adjust the worker count as needed
	workerPool.Start()

	// Initialize Gin router
	r := gin.Default()

	// Initialize API routes
	api.InitRoutes(r, database, workerPool)

	// Start the server
	if err := r.Run(":8080"); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
