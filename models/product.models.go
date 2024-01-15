package models

type Product struct {
	Code        string `json:"code"`
	Name        string `json:"name"`
	Quantity    int    `json:"quantity"`
	Description string `json:"description"`
	IsActive    int    `json:"is_active"`
}

func (Product) TableName() string {
	return "product"
}

type CreateProduct struct {
	Code        string `json:"code"`
	Name        string `json:"name"`
	Quantity    int    `json:"quantity"`
	Description string `json:"description"`
	IsActive    bool   `json:"is_active"`
}

type UpdateProduct struct {
	Code        string `json:"code"`
	Name        string `json:"name"`
	Quantity    int    `json:"quantity"`
	Description string `json:"description"`
	IsActive    bool   `json:"is_active"`
}

type CreateProducts []*CreateProduct
type UpdateProducts []*UpdateProduct
type Products []*Product
