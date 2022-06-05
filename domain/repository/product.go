package repository

import (
	"product_manager/domain/model"
)

type ProductRepository interface {
	Create(product *model.Product) error
	Update(product *model.Product) error
	Delete(product *model.Product) error
	FindByCode(code string) (*model.Product, error)
	FindAll() ([]*model.Product, error)
}
