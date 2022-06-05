package model

var (
//Err
)

type Stock struct {
	General
	ID        int     `gorm:"primaryKey" json:"id"`
	Total     int     `json:"total"`
	Cute      int     `json:"cute"`
	Available int     `json:"available"`
	Product   Product `gorm:"foreignKey:ID" json:"product"`
}
