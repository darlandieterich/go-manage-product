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

func (s *ProductService) UpdateProduct(ctx context.Context, uuid uuid.UUID, params *ProductParam) (err error) {
	product, err := s.products.FindById(ctx, uuid)

	if err != nil {
		return err
	}

	product.Name = params.Name
	product.PriceFrom = params.PriceFrom
	product.PriceTo = params.PriceTo
	product.Stock.Total = params.StockTotal
	product.Stock.Cute = params.StockCute
	product.Stock.Available = params.StockTotal - params.StockCute

	if err = model.ProductValidation(product); err != nil {
		return err
	}

	if err = s.products.Update(ctx, product); err != nil {
		return err
	}

	return nil
}

func (s *ProductService) GetProduct(ctx context.Context, code string) (product *model.Product, err error) {
	if code == "" {
		return nil, model.ErrProductCode
	}

	return s.products.FindByCode(ctx, code)
}
