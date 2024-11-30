package model

import (
	"time"

	db "github.com/ij4l/go-htmx/external/database/sqlc"
)

type Product struct {
	ProductID int32     `json:"product_id"`
	Name      string    `json:"name"`
	Category  string    `json:"category"`
	Price     string    `json:"price"`
	ImageUrl  string    `json:"image_url"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type ProductANTD struct {
	Page         int
	NextPage     bool
	PrevPage     bool
	TotalShow    int
	TotalProduct int
	Products      []Product
}

func New(product db.Product) Product {
	return Product{
		ProductID: product.ProductID,
		Name:      product.Name,
		Category:  product.Category.String,
		Price:     product.Price.Int.String(),
		ImageUrl:  product.ImageUrl.String,
		CreatedAt: product.CreatedAt.Time,
		UpdatedAt: product.UpdatedAt.Time,
	}
}

func NewList(db_products []db.Product) (products []Product) {
	for _, product := range db_products {
		products = append(products, New(product))
	}
	return
}
