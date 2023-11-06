package dto

import "mime/multipart"

type CreateProductForm struct {
	Name  string               `form:"name" validate:"required"`
	Desc  string               `form:"desc" validate:"required"`
	Price float64              `form:"price" validate:"required"`
	Image multipart.FileHeader `form:"image" validate:"required"`
}

type UpdateProductForm struct {
	Name  string               `form:"name"`
	Desc  string               `form:"desc"`
	Price float64              `form:"price"`
	Image multipart.FileHeader `form:"image" swaggerignore:"true"`
}

type CreateProductResponse struct {
	ID    uint    `json:"id"`
	Slug  string  `json:"slug"`
	Name  string  `json:"name"`
	Desc  string  `json:"desc"`
	Price float64 `json:"price"`
	Image string  `json:"image"`
}

type ProductResponse struct {
	CreateProductResponse
}
