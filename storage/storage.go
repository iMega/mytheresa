package storage

import (
	"context"
	"fmt"

	"github.com/imega/mytheresa/domain"
)

// Storage is a simple storage and bases on map.
type Storage struct {
	Store map[string][]byte
}

func New() *Storage {
	return &Storage{
		Store: make(map[string][]byte),
	}
}

func (storage *Storage) Get(
	ctx context.Context,
	key domain.Key,
) (domain.Value, error) {
	data, ok := storage.Store[string(key)]
	if !ok {
		return nil, fmt.Errorf(
			"product does not exists, %w",
			domain.ErrKeyDoesNotExists,
		)
	}

	return data, nil
}

func (storage *Storage) Set(
	ctx context.Context,
	key domain.Key,
	value domain.Value,
) error {
	storage.Store[string(key)] = value

	return nil
}
