package model

import "errors"

var (
	ErrStockTotal = errors.New("o total deve ser maior que zero")
	ErrStockCute  = errors.New("o total deve ser maior que o corte")
)

type Stock struct {
	General
	ID        int     `gorm:"primaryKey" json:"id"`
	Total     int     `json:"total"`
	Cute      int     `json:"cute"`
	Available int     `json:"available"`
	Product   Product `gorm:"foreignKey:ID" json:"product"`
}
