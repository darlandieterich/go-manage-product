package repository

import (
	"product_manager/domain/model"
)

type StockRepository interface {
	Create(stock *model.Stock) error
	Update(stock *model.Stock) error
	Delete(stock *model.Stock) error
	FindByProductID(productID int) (*model.Stock, error)
	FindAll() ([]*model.Stock, error)
}
