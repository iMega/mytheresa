package domain

import (
	"context"
)

const (
	RootNodeKey = "root"
	CategoryKey = "category"
	ProductKey  = "product"
	Base10      = 10
	Bit64       = 64
)

// Node contains SKUs of products.
type Node []string

// Product is a item of catalog any shop.
type Product struct {
	Name     string
	SKU      string
	Category string
	Price    Money
}

func (p *Product) GetKey() Key {
	return Key(ProductKey + p.SKU)
}

// Money represents an amount of money with its currency type.
type Money struct {
	Currency string
	Units    uint64
}

// Shop is an interface and is a behavior store.
type Shop interface {
	Get(context.Context, Request) ([5]Offer, error)
	Add(context.Context, Product) error
}

// Request helps filter products.
type Request struct {
	// Can be filtered by category.
	Category string

	// Can be filtered by priceLessThan.
	PriceLessThan uint64
}

// Offer - the current price of the item.
type Offer struct {
	Product  Product
	Final    Money
	Discount string
}

// Discount.
type Discount struct {
	Price Money
	Value string
}

// Discounter will calc discount.
type Discounter interface {
	Calc(Product) Discount
}
