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
	Stock     Stock     `json:"foreign_key:product_id"`
}

func NewProduct(code, name string,
	stockTotal, stockCute uint,
	priceFrom, priceTo float32) (*Product, error) {
	if code == "" {
		return nil, ErrProductCode
	}

	if name == "" {
		return nil, ErrProductName
	}

	if stockCute > stockTotal {
		return nil, ErrStockCute
	}

	if priceFrom < priceTo {
		return nil, ErrProductPrice
	}

	return &Product{
		ID:        uuid.New(),
		Code:      code,
		Name:      name,
		PriceFrom: priceFrom,
		PriceTo:   priceTo,
		Stock: Stock{
			Total:     stockTotal,
			Cute:      stockCute,
			Available: stockTotal - stockCute,
		},
	}, nil
}
