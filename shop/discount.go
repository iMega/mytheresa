package shop

import "github.com/imega/mytheresa/domain"

const (
	discount30 = 30
	discount15 = 15
	hundred    = 100
)

type Discount struct {
	WithDiscount30 bool
	WithDiscount15 bool
}

func (d *Discount) Calc(product domain.Product) domain.Discount {
	if d.WithDiscount30 && product.Category == "boots" {
		return domain.Discount{
			Price: domain.Money{
				Units:    product.Price.Units - (product.Price.Units * discount30 / hundred),
				Currency: product.Price.Currency,
			},
			Value: "30%",
		}
	}

	if d.WithDiscount15 && product.SKU == "000003" {
		return domain.Discount{
			Price: domain.Money{
				Units:    product.Price.Units - (product.Price.Units * discount15 / hundred),
				Currency: product.Price.Currency,
			},
			Value: "15%",
		}
	}

	return domain.Discount{
		Price: domain.Money{
			Units:    product.Price.Units,
			Currency: product.Price.Currency,
		},
		Value: "",
	}
}
