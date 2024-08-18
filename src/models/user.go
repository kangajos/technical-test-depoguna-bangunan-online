package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	ID       int64  `gorm:"primaryKey"`
	Name     string `json:"name"`
	UserName string `json:"userName"`
	Password string `json:"password"`

	Orders []Order `gorm:"foreignKey:UserID" json:"orders"` // Correct foreignKey tag
}

func (User) TableName() string {
	return "users"
}
