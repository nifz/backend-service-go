package dtos

import "time"

type ProductInput struct {
	CodeProduct string `json:"code_product"`
	Name        string `json:"product_name"`
	Image       string `json:"product_image"`
	Description string `json:"product_description"`
	Price       int    `json:"product_price"`
	Quantity    int    `json:"product_quantity"`
}

type ProductResponse struct {
	CodeProduct string     `json:"code_product"`
	Name        string     `json:"product_name"`
	Image       string     `json:"product_image"`
	Description string     `json:"product_description"`
	Price       int        `json:"product_price"`
	Quantity    int        `json:"product_quantity"`
	CreatedAt   *time.Time `json:"created_at,omitempty"`
	UpdatedAt   *time.Time `json:"updated_at,omitempty"`
}
