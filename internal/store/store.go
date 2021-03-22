package store

import (
	"HDTwG/model"
	"context"
)

type Options struct {
	IP   string
	Lang string
}

type Store interface {
	Get(ctx context.Context, opts Options) (model.Translation, error)
	Put(ctx context.Context) error
}
