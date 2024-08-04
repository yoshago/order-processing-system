package processor

import (
	"time"

	"github.com/yoshago/order-processing-system/internal/models"
	"gorm.io/gorm"
)

func Worker(orders <-chan models.Order, results chan<- models.Order, db *gorm.DB) {
	for order := range orders {
		// Simulate order processing
		time.Sleep(time.Second * 2) // Simulate time-consuming task
		order.Status = "completed"
		order.Result = "processed"

		// Update order status in DB
		db.Model(&models.Order{}).Where("id = ?", order.ID).Updates(order)

		// Send result to results channel
		results <- order
	}
}
