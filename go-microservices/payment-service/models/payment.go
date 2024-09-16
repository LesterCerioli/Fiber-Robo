package models

import "gorm.io/gorm"

type Payment struct {
	gorm.Model
	Amount      float64 `json:"amount"`
	Description string  `json:"description"`
	CustomerID  uint    `json:"customer_id"`
}
