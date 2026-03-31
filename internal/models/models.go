package models

import "gorm.io/gorm"

type Tv struct {
	gorm.Model
	Brand    string `json:"brand"`
	Price    int    `json:"price"`
	Quantity int    `json:"qty"`
}
