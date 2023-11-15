package demo

import (
	"github.com/samber/lo"
)

type ProductItem struct {
	ID       uint
	Name     string
	Price    float64
	Quantity uint
}

type PurchaseItem struct {
	ID       uint
	Price    float64
	Quantity uint
}

type Summary struct {
	PurchaseDate string
	Status       string
	TotalPrice   float64
}

type Repository[T any, U any] interface {
	Save(T) (U, error)
}

type PurchaseRepository struct {
}

func (p *PurchaseRepository) Save(items []PurchaseItem) (*Summary, error) {
	// Save to the database
	return &Summary{}, nil
}

func Purchase(repo Repository[[]PurchaseItem, *Summary], items []ProductItem) (float64, error) {
	purchaseItems := lo.Map(items, func(item ProductItem, index int) PurchaseItem {
		return PurchaseItem{
			ID:       item.ID,
			Price:    item.Price,
			Quantity: item.Quantity,
		}
	})

	summary, err := repo.Save(purchaseItems)

	if err != nil {
		return 0.0, err
	}

	return summary.TotalPrice, nil
}
