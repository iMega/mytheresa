package shop

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/imega/mytheresa/domain"
)

type Shop struct {
	Storage domain.Storage
}

func (shop *Shop) Get(
	ctx context.Context,
	req domain.Request,
) ([5]domain.Offer, error) {
	result := [5]domain.Offer{}

	var skus []string
	if req.Category != "" {
		skus = []string{}
	} else {
		rootNode := &RootNode{Storage: shop.Storage}
		res, err := rootNode.GetSKUs(ctx)
		if err != nil {
			return result, fmt.Errorf("failed to get skus, %w", err)
		}

		skus = res
	}

	if len(skus) == 0 {
		return result, nil
	}

	for i := 0; i < 5; i++ {
		data, err := shop.Storage.Get(ctx, domain.Key(domain.ProductKey+skus[i]))
		if err != nil {
			return result, fmt.Errorf("failed to get product, %w", err)
		}

		product := domain.Product{}
		if err := json.Unmarshal(data, &product); err != nil {
			return result, fmt.Errorf("failed to unmarshal product, %w", err)
		}

		result[i] = domain.Offer{
			Product: product,
			Final:   product.Price,
		}
	}

	return result, nil
}

func (shop *Shop) Add(ctx context.Context, product domain.Product) error {
	rootNode := &RootNode{Storage: shop.Storage}
	if err := rootNode.AddSKU(ctx, product.SKU); err != nil {
		return fmt.Errorf("failed to add sku to rootNode, %w", err)
	}

	data, err := json.Marshal(product)
	if err != nil {
		return fmt.Errorf("failed to marshal product, %w", err)
	}

	err = shop.Storage.Set(ctx, product.GetKey(), domain.Value(data))
	if err != nil {
		return fmt.Errorf("failed to store product, %w", err)
	}

	return nil
}
