package repository_test

import (
	"database/sql/driver"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/babelcoder-enterprise-courses/go-fiber-testing/model"
	"github.com/babelcoder-enterprise-courses/go-fiber-testing/repository"
	"github.com/stretchr/testify/suite"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type RepositoryTestSuite struct {
	suite.Suite
}

func (s *RepositoryTestSuite) TestFindAll() {
	tests := []struct {
		name    string
		term    string
		rows    *sqlmock.Rows
		args    []driver.Value
		sql     string
		results []model.Product
	}{
		{
			name: "without term",
			term: "",
			rows: sqlmock.NewRows([]string{"id", "name", "slug", "desc", "price", "image"}).
				AddRow(1, "Test 1", "test-1", "Desc 1", 10, "https://image.com").
				AddRow(2, "Test 2", "test-2", "Desc 2", 20, "https://image.com").
				AddRow(3, "Test 3", "test-3", "Desc 3", 30, "https://image.com"),
			sql:  `SELECT \* FROM "products" WHERE "products"."deleted_at" IS NULL ORDER BY id desc`,
			args: nil,
			results: []model.Product{
				{Model: gorm.Model{ID: 1}, Name: "Test 1", Slug: "test-1", Desc: "Desc 1", Price: 10, Image: "https://image.com"},
				{Model: gorm.Model{ID: 2}, Name: "Test 2", Slug: "test-2", Desc: "Desc 2", Price: 20, Image: "https://image.com"},
				{Model: gorm.Model{ID: 3}, Name: "Test 3", Slug: "test-3", Desc: "Desc 3", Price: 30, Image: "https://image.com"},
			},
		},
		{
			name: "with term",
			term: "2",
			rows: sqlmock.NewRows([]string{"id", "name", "slug", "desc", "price", "image"}).
				AddRow(2, "Test 2", "test-2", "Desc 2", 20, "https://image.com"),
			sql:  `SELECT \* FROM "products" WHERE LOWER\(name\) LIKE LOWER\(\$1\) AND "products"."deleted_at" IS NULL ORDER BY id desc`,
			args: nil,
			results: []model.Product{
				{Model: gorm.Model{ID: 2}, Name: "Test 2", Slug: "test-2", Desc: "Desc 2", Price: 20, Image: "https://image.com"},
			},
		},
	}

	for _, tc := range tests {
		db, mock, err := sqlmock.New()
		if err != nil {
			s.T().Fatalf("an error '%s' was not expecteds when opening a stub database connection", err)
		}
		defer db.Close()

		mock.
			ExpectQuery(tc.sql).
			WithArgs(tc.args...).
			WillReturnRows(tc.rows)

		gDB, _ := gorm.Open(postgres.New(postgres.Config{Conn: db}))
		repo := repository.Product{DB: gDB}
		products := repo.FindAll(tc.term)

		err = mock.ExpectationsWereMet()
		s.NoError(err)
		s.Len(products, len(tc.results))

		for i := range products {
			s.Equal(tc.results[i].ID, products[i].ID)
			s.Equal(tc.results[i].Name, products[i].Name)
			s.Equal(tc.results[i].Desc, products[i].Desc)
			s.Equal(tc.results[i].Slug, products[i].Slug)
			s.Equal(tc.results[i].Image, products[i].Image)
		}
	}

}

func TestRepositoryTestSuite(t *testing.T) {
	suite.Run(t, new(RepositoryTestSuite))
}
