package shop

import "github.com/imega/mytheresa/domain"

const (
	discount30 = 30
	discount15 = 15
	hundred    = 100
)

type Discount struct{}

func (d *Discount) Calc(product domain.Product) domain.Discount {
	if product.Category == "boots" {
		return domain.Discount{
			Price: domain.Money{
				Units:    product.Price.Units - (product.Price.Units * discount30 / hundred),
				Currency: product.Price.Currency,
			},
			Value: "30%",
		}
	}

	if product.SKU == "000003" {
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
