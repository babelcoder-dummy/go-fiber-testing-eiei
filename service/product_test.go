package service_test

import (
	"testing"

	mock_repository "github.com/babelcoder-enterprise-courses/go-fiber-testing/mocks/repository"
	"github.com/babelcoder-enterprise-courses/go-fiber-testing/model"
	"github.com/babelcoder-enterprise-courses/go-fiber-testing/service"
	"github.com/stretchr/testify/suite"
	"gorm.io/gorm"
)

type ProductTestSuite struct {
	suite.Suite
}

func (s *ProductTestSuite) TestFindAll() {
	repo := mock_repository.NewMockProducter(s.T())
	service := service.Product{Repository: repo}
	products := []model.Product{
		{Model: gorm.Model{ID: 1}, Name: "Test 1", Slug: "test-1", Desc: "Desc 1"},
		{Model: gorm.Model{ID: 2}, Name: "Test 2", Slug: "test-2", Desc: "Desc 2"},
	}
	term := "Test"

	repo.EXPECT().FindAll(term).Return(products)

	s.Equal(products, service.FindAll(term))
}

func TestProductTestSuite(t *testing.T) {
	suite.Run(t, new(ProductTestSuite))
}
