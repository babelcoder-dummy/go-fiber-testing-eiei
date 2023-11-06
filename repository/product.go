package repository

import (
	"github.com/babelcoder-enterprise-courses/go-fiber-testing/model"
	"gorm.io/gorm"
)

type Product struct {
	DB *gorm.DB
}

func (p *Product) Create(product *model.Product) error {
	return p.DB.Create(product).Error
}

func (p *Product) FindAll(term string) []model.Product {
	var products []model.Product
	query := p.DB.Order("id desc")

	if term != "" {
		query = query.Where("LOWER(name) LIKE LOWER(?)", "%"+term+"%")
	}

	query.Find(&products)

	return products
}

func (p *Product) FindByID(id uint) (*model.Product, error) {
	product := new(model.Product)
	err := p.DB.First(product, id).Error

	return product, err
}

func (p *Product) Update(product *model.Product) error {
	return p.DB.Save(product).Error
}

func (p *Product) Delete(id uint) {
	p.DB.Unscoped().Delete(&model.Product{}, id)
}
