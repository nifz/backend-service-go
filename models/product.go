package models

import (
	"time"

	"gorm.io/gorm"
)

type Product struct {
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `json:"deleted_at"`
	Code        string         `gorm:"primaryKey" json:"product_code"`
	Name        string         `json:"product_name"`
	Image       string         `json:"product_image"`
	Description string         `gorm:"type:text" json:"product_description"`
	Price       int            `json:"product_price"`
	Quantity    int            `json:"product_quantity"`
}
