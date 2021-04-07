package store

import (
	"HDTwG/model"
	"context"
	"encoding/json"
	"fmt"
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

func (c *ClientNoSQL) Get(ctx context.Context, opts Options) ([]model.Translation, error) {

	val, err := c.rdb.Do(ctx, "HGET", "location", opts.IP).Result()
	if err != nil {
		return nil, fmt.Errorf("Not found IP")
	}
	fmt.Printf("this is val = %v\n", val)
	return nil, nil
}

func (c *ClientNoSQL) Put(ctx context.Context, translations Translations, ip []model.Location) error {

	var err *redis.Cmd
	wg := sync.WaitGroup{}
	wg.Add(3000)
	nb_iplocations := len(ip) / 500
	nb_translationsfr := len(translations.TranslationFR) / 500
	nb_translationsen := len(translations.TranslationEN) / 500
	nb_translationses := len(translations.TranslationES) / 500

	for i := 0; i < 500; i++ {
		go func(i int, nb_iplocations int, ip []model.Location) error {
			defer wg.Done()

			for j := i * nb_iplocations; j < (i+1)*nb_iplocations; j++ {
				err = c.rdb.Do(ctx, "HSET", "location", ip[j].Address, ip[j].UUID)
				if err != nil {
					fmt.Println(err)
					return err.Err()
				}

			}
			return nil
		}(i, nb_iplocations, ip)
		go func(i int, nb_translationsfr int, translations Translations) error {
			defer wg.Done()
			translationsFR := translations.TranslationFR

			for j := i * nb_translationsfr; j < (i+1)*nb_translationsfr; j++ {
				mTransFR, _ := json.Marshal(translationsFR[j])
				err = c.rdb.Do(ctx, "HSET", "translationFR", translationsFR[j].UUID, mTransFR)
				if err != nil {
					fmt.Println(err)
					return err.Err()
				}

			}
			return nil
		}(i, nb_translationsfr, translations)
		go func(i int, nb_translationses int, translations Translations) error {
			defer wg.Done()
			translationsES := translations.TranslationES

			for j := i * nb_translationses; j < (i+1)*nb_translationses; j++ {
				mTransES, _ := json.Marshal(translationsES[j])
				err = c.rdb.Do(ctx, "HSET", "translationES", translationsES[j].UUID, mTransES)
				if err != nil {
					fmt.Println(err)
					return err.Err()
				}

			}
			return nil
		}(i, nb_translationses, translations)
		go func(i int, nb_translationsen int, translations Translations) error {
			defer wg.Done()
			translationsEN := translations.TranslationEN

			for j := i * nb_translationsen; j < (i+1)*nb_translationsfr; j++ {
				mTransEN, _ := json.Marshal(translationsEN[j])
				err = c.rdb.Do(ctx, "HSET", "translationEN", translationsEN[j].UUID, mTransEN)
				if err != nil {
					fmt.Println(err)
					return err.Err()
				}

			}
			return nil
		}(i, nb_translationsen, translations)
	}

	wg.Wait()
	return nil
}
