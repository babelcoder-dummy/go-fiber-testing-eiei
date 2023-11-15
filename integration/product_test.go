package integration_test

import (
	"encoding/json"
	"io"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/babelcoder-enterprise-courses/go-fiber-testing/config"
	"github.com/babelcoder-enterprise-courses/go-fiber-testing/dto"
	"github.com/babelcoder-enterprise-courses/go-fiber-testing/server"
	"github.com/go-testfixtures/testfixtures/v3"
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/suite"
	"gopkg.in/yaml.v3"
)

type ProductTestSuite struct {
	suite.Suite
	app      *fiber.App
	fixtures *testfixtures.Loader
	data     []dto.CreateProductResponse
}

func (s *ProductTestSuite) SetupSuite() {
	s.T().Setenv("APP_ENV", "test")
	s.setupServer()
	s.loadFixtures()
	s.loadData()
}

func (s *ProductTestSuite) SetupTest() {
	if err := s.fixtures.Load(); err != nil {
		s.T().Fatal("cannot load fixtures to the database")
	}
}

func (s *ProductTestSuite) setupServer() {
	testServer := server.New()
	s.app = testServer.App
	testServer.Setup()
}

func (s *ProductTestSuite) loadFixtures() {
	db, _ := config.DB.DB()
	fixtures, err := testfixtures.New(
		testfixtures.Database(db),
		testfixtures.Dialect("postgres"),
		testfixtures.Paths("fixtures/products.yml"),
	)
	if err != nil {
		s.T().Fatal("cannot load fixtures files")
	}
	s.fixtures = fixtures
}

func (s *ProductTestSuite) loadData() {
	productsData, err := os.ReadFile("fixtures/products.yml")
	if err != nil {
		s.T().Fatal("cannot load fixtures data")
	}

	yaml.Unmarshal(productsData, &s.data)
}

func (s *ProductTestSuite) TestGetProducts() {
	req := httptest.NewRequest("GET", "/v1/products", nil)
	res, err := s.app.Test(req)
	body, _ := io.ReadAll(res.Body)
	expected, _ := json.Marshal(s.data)

	s.Equal(200, res.StatusCode)
	s.NoError(err)
	s.Equal("application/json", res.Header.Get("Content-Type"))
	s.Equal(string(expected), string(body))
}

func TestProductTestSuite(t *testing.T) {
	suite.Run(t, new(ProductTestSuite))
}
