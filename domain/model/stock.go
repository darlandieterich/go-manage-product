package model

import (
	"errors"

	"github.com/google/uuid"
)

var (
	ErrStockTotal = errors.New("o total deve ser maior que zero")
	ErrStockCute  = errors.New("o total deve ser maior que o corte")
)

type Stock struct {
	General
	ID        uuid.UUID `gorm:"index,primaryKey" json:"id"`
	Total     uint      `json:"total"`
	Cute      uint      `json:"cute"`
	Available uint      `json:"available"`
	ProductID uuid.UUID `gorm:"index" json:"product_id"`
}
