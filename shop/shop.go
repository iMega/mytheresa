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

	key := domain.RootNodeKey
	if req.Category != "" {
		key = domain.CategoryKey + req.Category
	}

	node := &Node{Storage: shop.Storage, Key: domain.Key(key)}

	skus, err := node.GetSKUs(ctx)
	if err != nil {
		return result, fmt.Errorf("failed to get skus, %w", err)
	}

	if len(skus) == 0 {
		return result, nil
	}

	for idx, sku := range skus {
		if isLimitOffers(idx) {
			break
		}

		data, err := shop.Storage.Get(ctx, domain.Key(domain.ProductKey+sku))
		if err != nil {
			return result, fmt.Errorf("failed to get product, %w", err)
		}

		var product domain.Product
		if err := json.Unmarshal(data, &product); err != nil {
			return result, fmt.Errorf("failed to unmarshal product, %w", err)
		}

		result[idx] = domain.Offer{
			Product: product,
			Final:   product.Price,
		}
	}

	return result, nil
}

func (shop *Shop) Add(ctx context.Context, product domain.Product) error {
	rootNode := &Node{Storage: shop.Storage, Key: domain.RootNodeKey}
	if err := rootNode.AddSKU(ctx, product.SKU); err != nil {
		return fmt.Errorf("failed to add sku to rootNode, %w", err)
	}

	category := &Node{
		Storage: shop.Storage,
		Key:     domain.Key(domain.CategoryKey + product.Category),
	}
	if err := category.AddSKU(ctx, product.SKU); err != nil {
		return fmt.Errorf("failed to add sku to category, %w", err)
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

const limitOffers = 5

func isLimitOffers(idx int) bool {
	return idx == limitOffers
}
