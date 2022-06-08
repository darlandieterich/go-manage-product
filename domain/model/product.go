package model

import (
	"errors"

	"github.com/google/uuid"
)

var (
	ErrProductNotFound     = errors.New("produto não encontrado")
	ErrProductExistentCode = errors.New("código já existente")
	ErrProductCode         = errors.New("produto com codigo inválido")
	ErrProductName         = errors.New("produto com nome inválido")
	ErrProductPrice        = errors.New("o preço de não deve ser inferior ao preço por")
)

type Product struct {
	General
	ID        uuid.UUID `gorm:"index,primaryKey" json:"id"`
	Code      string    `gorm:"index" json:"code"`
	Name      string    `json:"name"`
	PriceFrom float32   `json:"price_from"`
	PriceTo   float32   `json:"price_to"`
	Stock     Stock     `gorm:"foreignKey:ID" json:"stock"`
}

func NewProduct(code, name string,
	stockTotal, stockCute uint,
	priceFrom, priceTo float32) (product *Product, err error) {

	uuid := uuid.New()

	product = &Product{
		ID:        uuid,
		Code:      code,
		Name:      name,
		PriceFrom: priceFrom,
		PriceTo:   priceTo,
		Stock: Stock{
			Total:     stockTotal,
			Cute:      stockCute,
			Available: stockTotal - stockCute,
		},
	}

	if err = ProductValidation(product); err != nil {
		return nil, err
	}

	return product, nil
}

func ProductValidation(product *Product) error {
	if product.Code == "" {
		return ErrProductCode
	}

	if product.Name == "" {
		return ErrProductName
	}

	if product.Stock.Total == 0 {
		return ErrStockTotal
	}

	if product.Stock.Cute > product.Stock.Total {
		return ErrStockCute
	}

	if product.PriceFrom < product.PriceTo {
		return ErrProductPrice
	}

	return nil
}
