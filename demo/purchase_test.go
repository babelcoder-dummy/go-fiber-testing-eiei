package demo_test

import (
	"testing"

	"github.com/babelcoder-enterprise-courses/go-fiber-testing/demo"
	mock_demo "github.com/babelcoder-enterprise-courses/go-fiber-testing/mocks/demo"
	"github.com/stretchr/testify/assert"
)

func TestPurchase(t *testing.T) {
	repo := mock_demo.NewMockRepository[[]demo.PurchaseItem, *demo.Summary](t)
	productItems := []demo.ProductItem{
		{ID: 1, Quantity: 10, Name: "Taylor", Price: 1_000},
		{ID: 2, Quantity: 20, Name: "Swift", Price: 2_000},
	}
	purchaseItems := []demo.PurchaseItem{
		{ID: 1, Quantity: 10, Price: 1_000},
		{ID: 2, Quantity: 20, Price: 2_000},
	}

	repo.EXPECT().Save(purchaseItems).Return(
		&demo.Summary{PurchaseDate: "14/11/2023", Status: "Success", TotalPrice: 50_000},
		nil,
	).Once()

	totalPrice, err := demo.Purchase(repo, productItems)

	repo.AssertExpectations(t)
	assert.Equal(t, 50_000.0, totalPrice)
	assert.NoError(t, err)
}
