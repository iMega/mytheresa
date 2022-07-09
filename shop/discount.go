package shop

import "github.com/imega/mytheresa/domain"

type Discount struct{}

func (d *Discount) Calc(product domain.Product) domain.Discount {
	if product.Category == "boots" {
		return domain.Discount{
			Price: domain.Money{
				Units:    product.Price.Units - (product.Price.Units * 30 / 100),
				Currency: product.Price.Currency,
			},
			Value: "30%",
		}
	}

	if product.SKU == "000003" {
		return domain.Discount{
			Price: domain.Money{
				Units:    product.Price.Units - (product.Price.Units * 15 / 100),
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
