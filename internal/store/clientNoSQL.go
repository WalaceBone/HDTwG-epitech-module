package store

import (
	"HDTwG/model"
	"context"
	"fmt"
	"log"
	"sync"

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
		PoolSize: 500,
	})

	return nil
}

func (c *ClientNoSQL) Get(ctx context.Context, opts Options) (model.Translation, error) {

	val, err := c.rdb.Do(ctx, "HGET", "album:1", "title").Result()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(val)
	return model.Translation{}, nil
}

func (c *ClientNoSQL) Put(ctx context.Context, translations Translations, ip []model.Location) error {

	wg := sync.WaitGroup{}
	wg.Add(500)
	portion := len(ip) / 500

	for i := 0; i < 500; i++ {
		fmt.Println(i)
		go func(i int, portion int, ip []model.Location) {
			defer wg.Done()

			for j := i * portion; j < (i+1)*portion; j++ {
				err := c.rdb.Do(ctx, "HSET", "location", ip[j].Address, ip[j].UUID)
				if err != nil {
					fmt.Println(err)
				}
			}
		}(i, portion, ip)
	}

	wg.Wait()
	return nil
}
