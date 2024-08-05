package processor

import (
	"github.com/yoshago/order-processing-system/internal/models"
	"gorm.io/gorm"
)

type WorkerPool struct {
	DB          *gorm.DB
	WorkerCount int
	OrderQueue  chan models.Order
	ResultQueue chan models.Order
}

func NewWorkerPool(db *gorm.DB, workerCount int) *WorkerPool {
	return &WorkerPool{
		DB:          db,
		WorkerCount: workerCount,
		OrderQueue:  make(chan models.Order, 100),
		ResultQueue: make(chan models.Order, 100),
	}
}

func (wp *WorkerPool) Start() {
	for i := 0; i < wp.WorkerCount; i++ {
		go Worker(wp.OrderQueue, wp.ResultQueue, wp.DB)
	}
}

func (wp *WorkerPool) AddOrder(order models.Order) {
	wp.OrderQueue <- order
}

func (wp *WorkerPool) GetResults() chan models.Order {
	return wp.ResultQueue
}
