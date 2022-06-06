package application

import (
	product "product_manager/domain/repository"
	"product_manager/infra/repo"

	"gorm.io/gorm"
)

type ProductService struct {
	products product.ProductRepository
}

func NewProductService(conn *gorm.DB) *ProductService {
	return &ProductService{
		products: repo.NewTasRepository(conn),
	}
}
