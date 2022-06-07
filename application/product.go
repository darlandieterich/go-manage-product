package application

import (
	"context"
	"product_manager/domain/model"
	product "product_manager/domain/repository"
	"product_manager/infra/repo"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ProductService struct {
	products product.ProductRepository
}

type ProductParam struct {
	Code       string  `json:"code"`
	Name       string  `json:"name"`
	StockTotal uint    `json:"stock_total"`
	StockCute  uint    `json:"stock_cute"`
	PriceFrom  float32 `json:"price_from"`
	PriceTo    float32 `json:"price_to"`
}

func NewProductService(conn *gorm.DB) *ProductService {
	return &ProductService{
		products: repo.NewProductRepository(conn),
	}
}

func (s *ProductService) CreateProduct(ctx context.Context, params *ProductParam) (id uuid.UUID, err error) {
	product, err := model.NewProduct(
		params.Code,
		params.Name,
		params.StockTotal,
		params.StockCute,
		params.PriceFrom,
		params.PriceTo,
	)

	if err != nil {
		return uuid.Nil, err
	}

	id, err = s.products.Create(ctx, product)

	if err != nil {
		return uuid.Nil, err
	}

	return id, nil
}
