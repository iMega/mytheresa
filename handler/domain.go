package handler

type Offer struct {
	SKU      string `json:"sku"`
	Name     string `json:"name"`
	Category string `json:"category"`
	Price    Price
}

type Price struct {
	Original int    `json:"original"`
	Final    int    `json:"final"`
	Discount string `json:"discount_percentage"`
	Currency string `json:"currency"`
}
