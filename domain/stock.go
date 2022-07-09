package domain

// Catalog looks like any regular shop.
type Catalog []Category

// Category is a item of catalog and it can contains any products.
type Category struct {
	Name     string
	Products []Product
}

// Product is a item of catalog any shop.
type Product struct {
	Name  string
	SKU   string
	Price Money
}

// Money represents an amount of money with its currency type.
type Money struct {
	Currency string
	Units    uint64
}
