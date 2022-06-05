package application

import (
	stock "product_manager/domain/repository"

	"gorm.io/gorm"
)

type StockRepository struct {
	stock stock.StockRepository
}

func NewStockRepository(conn *gorm.DB) *StockRepository {
	return &StockRepository{
		stock: nil,
	}
}
