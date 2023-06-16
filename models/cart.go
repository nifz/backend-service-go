package models

import "gorm.io/gorm"

type Cart struct {
	gorm.Model
	CodeProduct string  `gorm:"primaryKey" json:"product_code"`
	Product     Product `gorm:"foreignKey:CodeProduct;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"product"`
	Quantity    int     `json:"cart_quantity"`
}
