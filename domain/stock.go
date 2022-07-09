package domain

import (
	"context"
)

// Catalog looks like any regular shop.
type Catalog []Category

const RootNodeKey = "root"

// RootNode contains all SKUs of products.
type RootNode []string

// Category is a item of catalog and it can contains any products.
type Category struct {
	Name     string
	Products []Product
}

const ProductKey = "product"

// Product is a item of catalog any shop.
type Product struct {
	Name  string
	SKU   string
	Price Money
}

func (p *Product) GetKey() Key {
	return Key(ProductKey + p.SKU)
}

func (p *Product) MarshalJSON() ([]byte, error) {
	return []byte{}, nil
}

func (p *Product) UnmarshalJSON(b []byte) error {
	return nil
}

// Money represents an amount of money with its currency type.
type Money struct {
	Currency string
	Units    uint64
}

// Shop is an interface and is a behavior store.
type Shop interface {
	Get(context.Context, Request) [5]Offer
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
