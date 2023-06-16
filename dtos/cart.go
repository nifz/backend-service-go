package dtos

import "time"

type CartInput struct {
	CodeProduct string `json:"product_code"`
	Quantity    int    `json:"cart_quantity"`
}

type CartInputUpdate struct {
	Quantity int `json:"cart_quantity"`
}

type CartResponse struct {
	ID        int             `json:"id"`
	Product   ProductResponse `json:"product"`
	Quantity  int             `json:"cart_quantity"`
	CreatedAt *time.Time      `json:"created_at,omitempty"`
	UpdatedAt *time.Time      `json:"updated_at,omitempty"`
}
