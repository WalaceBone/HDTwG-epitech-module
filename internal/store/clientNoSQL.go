package store

import (
	"HDTwG/model"
	"context"

	"github.com/go-redis/redis/v8"
)

type ClientNoSQL struct {
	rdb *redis.Client
}

func NewNSQLClient() *ClientNoSQL {
	return &ClientNoSQL{}
}

func (c *ClientNoSQL) Init() error {
	c.rdb = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	return nil
}

func (c *ClientNoSQL) Get(ctx context.Context, opts Options) (model.Translation, error) {
	return model.Translation{}, nil
}

func (c *ClientNoSQL) Put(ctx context.Context, translations Translations, ip []model.Location) error {

	return nil
}
