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
	PriceFrom float32   `json:"preco_from"`
	PriceTo   float32   `json:"preco_to"`
}

func NewProduct(code, name string) (*Product, error) {
	if code == "" {
		return nil, ErrProductCode
	}

	if name == "" {
		return nil, ErrProductName
	}

	return &Product{
		Code: code,
		Name: name,
	}, nil
}
