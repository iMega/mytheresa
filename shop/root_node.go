package shop

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/imega/mytheresa/domain"
)

// RootNode contains all SKUs.
type RootNode struct {
	Storage domain.Storage
}

func (node *RootNode) GetSKUs(ctx context.Context) ([]string, error) {
	all, err := node.Storage.Get(ctx, domain.RootNodeKey)
	if err != nil && !errors.Is(err, domain.ErrKeyDoesNotExists) {
		return nil, fmt.Errorf("failed to get rootNode from storage, %w", err)
	}

	rootNode := domain.RootNode{}
	if all == nil {
		return rootNode, nil
	}

	if err := json.Unmarshal(all, &rootNode); err != nil {
		return nil, fmt.Errorf("failed to unmarshal rootNode, %w", err)
	}

	return rootNode, nil
}

func (node *RootNode) AddSKU(ctx context.Context, sku string) error {
	rootNode, err := node.GetSKUs(ctx)
	if err != nil {
		return fmt.Errorf("failed to get rootNode, %w", err)
	}

	rootNode = append(rootNode, sku)

	data, err := json.Marshal(rootNode)
	if err != nil {
		return fmt.Errorf("failed to marshal rootNode, %w", err)
	}

	err = node.Storage.Set(ctx, domain.Key(domain.RootNodeKey), data)
	if err != nil {
		return fmt.Errorf("failed to store rootNode, %w", err)
	}

	return nil
}
