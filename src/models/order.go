package models

import "gorm.io/gorm"

type Order struct {
	gorm.Model
	Name   string `json:"name"`
	Qty    int32  `json:"qty"`
	Price  int64  `json:"price"`
	UserID int64  `json:"userId"` // Change to UserID for GORM's convention

	// User User `gorm:"foreignKey:UserID"`
}

func (Order) TableName() string {
	return "orders"
}
