package application

import (
	product "product_manager/domain/repository"

	"gorm.io/gorm"
)

type ProductService struct {
	products product.ProductRepository
}

func NewProductService(conn *gorm.DB) *ProductService {
	return &ProductService{
		products: nil,
	}
}
