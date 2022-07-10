package shop

import (
	"context"
	"testing"

	"github.com/imega/mytheresa/domain"
	"github.com/imega/mytheresa/storage"
	"github.com/stretchr/testify/assert"
)

func TestShop_Get(t *testing.T) {
	type args struct {
		ctx context.Context
		req domain.Request
	}

	tests := []struct {
		name               string
		args               args
		discounter         domain.Discounter
		item000003Category string
		want               [5]domain.Offer
		wantErr            bool
	}{
		{
			name: "optimistic, get all products",
			args: args{ctx: context.Background()},
			want: [5]domain.Offer{
				{
					Product: domain.Product{
						SKU:      "000001",
						Name:     "BV Lean leather ankle boots",
						Category: "boots",
						Price:    domain.Money{Currency: "EUR", Units: 89000},
					},
					Final: domain.Money{Currency: "EUR", Units: 89000},
				},
				{
					Product: domain.Product{
						SKU:      "000002",
						Name:     "BV Lean leather ankle boots",
						Category: "boots",
						Price:    domain.Money{Currency: "EUR", Units: 99000},
					},
					Final: domain.Money{Currency: "EUR", Units: 99000},
				},
				{
					Product: domain.Product{
						SKU:      "000003",
						Name:     "Ashlington leather ankle boots",
						Category: "boots",
						Price:    domain.Money{Currency: "EUR", Units: 71000},
					},
					Final: domain.Money{Currency: "EUR", Units: 71000},
				},
				{
					Product: domain.Product{
						SKU:      "000004",
						Name:     "Naima embellished suede sandals",
						Category: "sandals",
						Price:    domain.Money{Currency: "EUR", Units: 79500},
					},
					Final: domain.Money{Currency: "EUR", Units: 79500},
				},
				{
					Product: domain.Product{
						SKU:      "000005",
						Name:     "Nathane leather sneakers",
						Category: "sneakers",
						Price:    domain.Money{Currency: "EUR", Units: 59000},
					},
					Final: domain.Money{Currency: "EUR", Units: 59000},
				},
			},
		},
		{
			name:       "get all products and apply the discount",
			args:       args{ctx: context.Background()},
			discounter: New(getRawLP()),
			want: [5]domain.Offer{
				{
					Product: domain.Product{
						SKU:      "000001",
						Name:     "BV Lean leather ankle boots",
						Category: "boots",
						Price:    domain.Money{Currency: "EUR", Units: 89000},
					},
					Final:    domain.Money{Currency: "EUR", Units: 62300},
					Discount: "30%",
				},
				{
					Product: domain.Product{
						SKU:      "000002",
						Name:     "BV Lean leather ankle boots",
						Category: "boots",
						Price:    domain.Money{Currency: "EUR", Units: 99000},
					},
					Final:    domain.Money{Currency: "EUR", Units: 69300},
					Discount: "30%",
				},
				{
					Product: domain.Product{
						SKU:      "000003",
						Name:     "Ashlington leather ankle boots",
						Category: "boots",
						Price:    domain.Money{Currency: "EUR", Units: 71000},
					},
					Final:    domain.Money{Currency: "EUR", Units: 49700},
					Discount: "30%",
				},
				{
					Product: domain.Product{
						SKU:      "000004",
						Name:     "Naima embellished suede sandals",
						Category: "sandals",
						Price:    domain.Money{Currency: "EUR", Units: 79500},
					},
					Final: domain.Money{Currency: "EUR", Units: 79500},
				},
				{
					Product: domain.Product{
						SKU:      "000005",
						Name:     "Nathane leather sneakers",
						Category: "sneakers",
						Price:    domain.Money{Currency: "EUR", Units: 59000},
					},
					Final: domain.Money{Currency: "EUR", Units: 59000},
				},
			},
		},
		{
			name: "get products filtered by category",
			args: args{
				ctx: context.Background(),
				req: domain.Request{Category: "sneakers"},
			},
			want: [5]domain.Offer{
				{
					Product: domain.Product{
						SKU:      "000005",
						Name:     "Nathane leather sneakers",
						Category: "sneakers",
						Price:    domain.Money{Currency: "EUR", Units: 59000},
					},
					Final: domain.Money{Currency: "EUR", Units: 59000},
				},
			},
		},
		{
			name: "discounts collide, the biggest discount must be applied",
			args: args{
				ctx: context.Background(),
				req: domain.Request{
					Category:      "boots",
					PriceLessThan: 72000,
				},
			},
			discounter: New(getRawLP()),
			want: [5]domain.Offer{
				{
					Product: domain.Product{
						SKU:      "000003",
						Name:     "Ashlington leather ankle boots",
						Category: "boots",
						Price:    domain.Money{Currency: "EUR", Units: 71000},
					},
					Final:    domain.Money{Currency: "EUR", Units: 49700},
					Discount: "30%",
				},
			},
		},
		{
			name: "products in the boots category have a 30% discount",
			args: args{
				ctx: context.Background(),
				req: domain.Request{Category: "boots"},
			},
			discounter: New(getRawLP()),
			want: [5]domain.Offer{
				{
					Product: domain.Product{
						SKU:      "000001",
						Name:     "BV Lean leather ankle boots",
						Category: "boots",
						Price:    domain.Money{Currency: "EUR", Units: 89000},
					},
					Final:    domain.Money{Currency: "EUR", Units: 62300},
					Discount: "30%",
				},
				{
					Product: domain.Product{
						SKU:      "000002",
						Name:     "BV Lean leather ankle boots",
						Category: "boots",
						Price:    domain.Money{Currency: "EUR", Units: 99000},
					},
					Final:    domain.Money{Currency: "EUR", Units: 69300},
					Discount: "30%",
				},
				{
					Product: domain.Product{
						SKU:      "000003",
						Name:     "Ashlington leather ankle boots",
						Category: "boots",
						Price:    domain.Money{Currency: "EUR", Units: 71000},
					},
					Final:    domain.Money{Currency: "EUR", Units: 49700},
					Discount: "30%",
				},
			},
		},
		{
			name: `The product with sku=000007 has a 15% discount`,
			args: args{
				ctx: context.Background(),
				req: domain.Request{Category: "otherCategory"},
			},
			discounter:         New(getRawLP()),
			item000003Category: "otherCategory",
			want: [5]domain.Offer{
				{
					Product: domain.Product{
						SKU:      "000003",
						Name:     "Ashlington leather ankle boots",
						Category: "otherCategory",
						Price:    domain.Money{Currency: "EUR", Units: 71000},
					},
					Final:    domain.Money{Currency: "EUR", Units: 60350},
					Discount: "15%",
				},
			},
		},
		{
			name: "get products filtered by priceLessThan 800 euro",
			args: args{
				ctx: context.Background(),
				req: domain.Request{
					PriceLessThan: 80000,
				},
			},
			discounter: New(getRawLP()),
			want: [5]domain.Offer{
				{
					Product: domain.Product{
						SKU:      "000003",
						Name:     "Ashlington leather ankle boots",
						Category: "boots",
						Price:    domain.Money{Currency: "EUR", Units: 71000},
					},
					Final:    domain.Money{Currency: "EUR", Units: 49700},
					Discount: "30%",
				},
				{
					Product: domain.Product{
						SKU:      "000004",
						Name:     "Naima embellished suede sandals",
						Category: "sandals",
						Price:    domain.Money{Currency: "EUR", Units: 79500},
					},
					Final: domain.Money{Currency: "EUR", Units: 79500},
				},
				{
					Product: domain.Product{
						SKU:      "000005",
						Name:     "Nathane leather sneakers",
						Category: "sneakers",
						Price:    domain.Money{Currency: "EUR", Units: 59000},
					},
					Final: domain.Money{Currency: "EUR", Units: 59000},
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			shop := shopInit(tt.discounter, tt.item000003Category)

			got, err := shop.Get(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("Shop.Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			assert.Equal(t, tt.want, got)
		})
	}
}

func shopInit(discounter domain.Discounter, item000003Category string) *Shop {
	ctx := context.Background()
	store := storage.New()
	shop := &Shop{Storage: store, Discounter: discounter}

	shop.Add(ctx, domain.Product{
		SKU:      "000001",
		Name:     "BV Lean leather ankle boots",
		Category: "boots",
		Price:    domain.Money{Currency: "EUR", Units: 89000},
	})

	shop.Add(ctx, domain.Product{
		SKU:      "000002",
		Name:     "BV Lean leather ankle boots",
		Category: "boots",
		Price:    domain.Money{Currency: "EUR", Units: 99000},
	})

	category := "boots"
	if item000003Category != "" {
		category = item000003Category
	}

	shop.Add(ctx, domain.Product{
		SKU:      "000003",
		Name:     "Ashlington leather ankle boots",
		Category: category,
		Price:    domain.Money{Currency: "EUR", Units: 71000},
	})

	shop.Add(ctx, domain.Product{
		SKU:      "000004",
		Name:     "Naima embellished suede sandals",
		Category: "sandals",
		Price:    domain.Money{Currency: "EUR", Units: 79500},
	})

	shop.Add(ctx, domain.Product{
		SKU:      "000005",
		Name:     "Nathane leather sneakers",
		Category: "sneakers",
		Price:    domain.Money{Currency: "EUR", Units: 59000},
	})

	shop.Add(ctx, domain.Product{
		SKU:      "000006",
		Name:     "product over limit",
		Category: "limit",
		Price:    domain.Money{Currency: "EUR", Units: 99999999},
	})

	return shop
}
