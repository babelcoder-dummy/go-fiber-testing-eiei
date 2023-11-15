package controller_test

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http/httptest"
	"testing"

	"github.com/babelcoder-enterprise-courses/go-fiber-testing/controller"
	"github.com/babelcoder-enterprise-courses/go-fiber-testing/dto"
	mock_service "github.com/babelcoder-enterprise-courses/go-fiber-testing/mocks/service"
	"github.com/babelcoder-enterprise-courses/go-fiber-testing/model"
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/suite"
	"gorm.io/gorm"
)

type ProductTestSuite struct {
	suite.Suite
}

func (s *ProductTestSuite) TestFindAll() {
	tests := []struct {
		name        string
		products    []model.Product
		fromService []model.Product
		term        string
		expected    []dto.ProductResponse
	}{
		{
			name: "existing products without term",
			products: []model.Product{
				{Model: gorm.Model{ID: 1}, Slug: "test-1", Name: "Test 1", Desc: "Desc 1", Price: 10, Image: "https://image.com"},
				{Model: gorm.Model{ID: 2}, Slug: "test-2", Name: "Test 2", Desc: "Desc 2", Price: 20, Image: "https://image.com"},
				{Model: gorm.Model{ID: 3}, Slug: "test-3", Name: "Test 3", Desc: "Desc 3", Price: 30, Image: "https://image.com"},
			},
			fromService: []model.Product{
				{Model: gorm.Model{ID: 1}, Slug: "test-1", Name: "Test 1", Desc: "Desc 1", Price: 10, Image: "https://image.com"},
				{Model: gorm.Model{ID: 2}, Slug: "test-2", Name: "Test 2", Desc: "Desc 2", Price: 20, Image: "https://image.com"},
				{Model: gorm.Model{ID: 3}, Slug: "test-3", Name: "Test 3", Desc: "Desc 3", Price: 30, Image: "https://image.com"},
			},
			term: "",
			expected: []dto.ProductResponse{
				{CreateProductResponse: dto.CreateProductResponse{ID: 1, Slug: "test-1", Name: "Test 1", Desc: "Desc 1", Price: 10, Image: "https://image.com"}},
				{CreateProductResponse: dto.CreateProductResponse{ID: 2, Slug: "test-2", Name: "Test 2", Desc: "Desc 2", Price: 20, Image: "https://image.com"}},
				{CreateProductResponse: dto.CreateProductResponse{ID: 3, Slug: "test-3", Name: "Test 3", Desc: "Desc 3", Price: 30, Image: "https://image.com"}},
			},
		},
		{
			name:        "no products without term",
			products:    []model.Product{},
			fromService: []model.Product{},
			term:        "",
			expected:    []dto.ProductResponse{},
		},
		{
			name:        "no products with term",
			products:    []model.Product{},
			fromService: []model.Product{},
			term:        "test",
			expected:    []dto.ProductResponse{},
		},
		{
			name: "existing products with term",
			products: []model.Product{
				{Model: gorm.Model{ID: 1}, Slug: "test-1", Name: "Test 1", Desc: "Desc 1", Price: 10, Image: "https://image.com"},
				{Model: gorm.Model{ID: 2}, Slug: "test-2", Name: "Test 2", Desc: "Desc 2", Price: 20, Image: "https://image.com"},
				{Model: gorm.Model{ID: 3}, Slug: "test-3", Name: "Test 3", Desc: "Desc 3", Price: 30, Image: "https://image.com"},
			},
			fromService: []model.Product{
				{Model: gorm.Model{ID: 2}, Slug: "test-2", Name: "Test 2", Desc: "Desc 2", Price: 20, Image: "https://image.com"},
			},
			term: "2",
			expected: []dto.ProductResponse{
				{CreateProductResponse: dto.CreateProductResponse{ID: 2, Slug: "test-2", Name: "Test 2", Desc: "Desc 2", Price: 20, Image: "https://image.com"}},
			},
		},
	}

	for _, tc := range tests {
		s.T().Run(tc.name, func(t *testing.T) {
			app := fiber.New()
			service := mock_service.NewMockProducter(t)
			controller := controller.Product{Service: service}

			service.EXPECT().FindAll(tc.term).Return(tc.fromService)

			app.Get("/products", controller.FindAll)
			res, err := app.Test(httptest.NewRequest("GET", fmt.Sprintf("/products?term=%s", tc.term), nil))
			body, _ := io.ReadAll(res.Body)
			expected, _ := json.Marshal(tc.expected)

			s.NoError(err)
			s.Equal(string(expected), string(body))
		})
	}
}

func TestProductTestSuite(t *testing.T) {
	suite.Run(t, new(ProductTestSuite))
}
