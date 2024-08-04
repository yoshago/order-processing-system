package models

import (
	"time"
)

type Order struct {
	ID        string `gorm:"primaryKey"`
	ProductID string
	Quantity  int
	Status    string
	Result    string
	CreatedAt time.Time
	UpdatedAt time.Time
}
