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

func (shop *Shop) Get(ctx context.Context, req domain.Request) [5]domain.Offer {
	result := [5]domain.Offer{}

	return result
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
