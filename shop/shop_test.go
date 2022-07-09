package shop

import (
	"context"
	"reflect"
	"testing"

	"github.com/imega/mytheresa/domain"
)

func TestShop_Get(t *testing.T) {
	type args struct {
		ctx context.Context
		req domain.Request
	}

	tests := []struct {
		name string
		shop *Shop
		args args
		want [5]domain.Offer
	}{
		{
			name: "optimistic, get all products",
		},
		{
			name: "get all products and apply the discount",
		},
		{
			name: "get products filtered by category and apply the discount",
		},
		{
			name: "discounts collide, the biggest discount must be applied",
		},
		{
			name: "products in the boots category have a 30% discount",
		},
		{
			name: "The product with sku=000003 has a 15% discount",
		},
		{
			name: "get products filtered by priceLessThan 800 euro",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			shop := &Shop{}

			got := shop.Get(tt.args.ctx, tt.args.req)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Shop.Get() = %v, want %v", got, tt.want)
			}
		})
	}
}
