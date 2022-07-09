package domain

import "context"

// Storage is an interface to store any entities.
type Storage interface {
	Get(context.Context, Key) (Value, error)
	Set(context.Context, Key, Value) error
}

// Key is unique identifier any entity.
type Key string

// Value is data of entity.
type Value []byte
