package service

import (
	"mime/multipart"
	"os"

	"github.com/babelcoder-enterprise-courses/go-fiber-testing/dto"
	"github.com/babelcoder-enterprise-courses/go-fiber-testing/model"
	"github.com/babelcoder-enterprise-courses/go-fiber-testing/repository"
	"github.com/jinzhu/copier"
)

type Product struct {
	Repository repository.Product
	Storage    Storage
}

func (p *Product) Create(form *dto.CreateProductForm, image *multipart.FileHeader) (*model.Product, error) {
	product := new(model.Product)
	copier.Copy(product, form)

	imagePath, err := p.Storage.Save(image)
	if err != nil {
		return nil, err
	}

	product.Image = imagePath
	err = p.Repository.Create(product)

	return product, err
}

func (p *Product) FindAll(term string) []model.Product {
	return p.Repository.FindAll(term)
}

func (p *Product) FindOne(id uint) (*model.Product, error) {
	return p.Repository.FindByID(id)
}

func (p *Product) Update(id uint, image *multipart.FileHeader, form *dto.UpdateProductForm) (*model.Product, error) {
	product, err := p.Repository.FindByID(id)
	if err != nil {
		return nil, err
	}

	if image == nil {
		copier.CopyWithOption(product, form, copier.Option{IgnoreEmpty: true})
		err = p.Repository.Update(product)
		if err != nil {
			return nil, err
		}

		return product, err
	}

	oldImagePath := product.Image
	copier.CopyWithOption(product, form, copier.Option{IgnoreEmpty: true})

	imagePath, err := p.Storage.Save(image)
	if err != nil {
		return nil, err
	}

	product.Image = imagePath

	err = p.Repository.Update(product)
	if err != nil {
		os.Remove(imagePath)

		return nil, err
	}

	p.Storage.Remove(oldImagePath)
	return product, nil
}

func (p *Product) Delete(id uint) {
	p.Repository.Delete(id)
}
