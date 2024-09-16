package models

import "gorm.io/gorm"

type Stock struct {
	gorm.Model
	Symbol string  `json:"symbol"`
	Price  float64 `json:"price"`
}
