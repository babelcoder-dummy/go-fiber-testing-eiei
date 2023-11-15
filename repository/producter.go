package repository

import "github.com/babelcoder-enterprise-courses/go-fiber-testing/model"

type Producter interface {
	Create(*model.Product) error
	FindAll(string) []model.Product
	FindByID(uint) (*model.Product, error)
	Update(*model.Product) error
	Delete(uint)
}
