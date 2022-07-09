package shop

import (
	"context"

	"github.com/imega/mytheresa/domain"
)

type Shop struct{}

func (shop *Shop) Get(ctx context.Context, req domain.Request) [5]domain.Offer {
	result := [5]domain.Offer{}

	return result
}
