package shop

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/imega/mytheresa/domain"
)

// Node contains all SKUs.
type Node struct {
	Key     domain.Key
	Storage domain.Storage
}

func (node *Node) GetSKUs(ctx context.Context) ([]string, error) {
	data, err := node.Storage.Get(ctx, node.Key)
	if err != nil && !errors.Is(err, domain.ErrKeyDoesNotExists) {
		return nil, fmt.Errorf("failed to get data from storage, %w", err)
	}

	skus := domain.Node{}
	if data == nil {
		return skus, nil
	}

	if err := json.Unmarshal(data, &skus); err != nil {
		return nil, fmt.Errorf("failed to unmarshal skus, %w", err)
	}

	return skus, nil
}

func (node *Node) AddSKU(ctx context.Context, sku string) error {
	skus, err := node.GetSKUs(ctx)
	if err != nil {
		return fmt.Errorf("failed to get skus, %w", err)
	}

	skus = append(skus, sku)

	data, err := json.Marshal(skus)
	if err != nil {
		return fmt.Errorf("failed to marshal skus, %w", err)
	}

	if err := node.Storage.Set(ctx, node.Key, data); err != nil {
		return fmt.Errorf("failed to store skus, %w", err)
	}

	return nil
}
