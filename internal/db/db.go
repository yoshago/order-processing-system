package db

import (
	"log"

	"github.com/yoshago/order-processing-system/internal/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDB() (*gorm.DB, error) {
	dsn := "host=localhost user=shaag password=752PGpanor dbname=orderdb port=5433 sslmode=disable TimeZone=Asia/Jerusalem"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	// Migrate the schema
	if err := db.AutoMigrate(&models.Order{}); err != nil {
		log.Fatal("Failed to migrate database:", err)
	}

	return db, nil
}
