package shop

import (
	"context"

	"github.com/imega/mytheresa/domain"
)

type Shop struct {
	Storage domain.Storage
}

func (shop *Shop) Get(ctx context.Context, req domain.Request) [5]domain.Offer {
	result := [5]domain.Offer{}

	return result
}

func (shop *Shop) Add(ctx context.Context, product domain.Product) error {
	return nil
}
