package service

import (
	"mime/multipart"

	"github.com/babelcoder-enterprise-courses/go-fiber-testing/dto"
	"github.com/babelcoder-enterprise-courses/go-fiber-testing/model"
)

type Producter interface {
	Create(*dto.CreateProductForm, *multipart.FileHeader) (*model.Product, error)
	FindAll(string) []model.Product
	FindOne(uint) (*model.Product, error)
	Update(uint, *multipart.FileHeader, *dto.UpdateProductForm) (*model.Product, error)
	Delete(uint)
}
