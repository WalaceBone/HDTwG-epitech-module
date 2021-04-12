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
		PoolSize: 16000,
	})

	return nil
}

func (c *ClientNoSQL) Get(ctx context.Context, opts Options) ([]model.Translation, error) {

	var translation string
	//var Store []model.Translation

	switch opts.Lang {
	case "English":
		translation = "translationEN"
	case "French":
		translation = "translationFR"
	case "Spanish":
		translation = "translationES"
	case "":
		translation = ""

	}

	uuid, err := c.rdb.Do(ctx, "HGET", "location", opts.IP).Result()
	if err != nil {
		return nil, fmt.Errorf("No IP found")
	}
	if translation != "" {
		store, err := c.rdb.Do(ctx, "HGET", translation, uuid).Result()
		if err != nil {
			return nil, fmt.Errorf("No store found")
		}

		fmt.Printf("this is the stores : %v", store)

	} else if translation == "" {
		store_en, err := c.rdb.Do(ctx, "HGET", "translationEN", uuid).Result()
		if err != nil {
			return nil, fmt.Errorf("No store found")
		}
		store_fr, err := c.rdb.Do(ctx, "HGET", "translationFR", uuid).Result()
		if err != nil {
			return nil, fmt.Errorf("No store found")
		}
		store_es, err := c.rdb.Do(ctx, "HGET", "translationES", uuid).Result()
		if err != nil {
			return nil, fmt.Errorf("No store found")
		}
		fmt.Printf("France: %v\nUS: %v\nSpain: %v\n", store_fr, store_en, store_es)
	}

	return nil, nil
}

func (c *ClientNoSQL) Put(ctx context.Context, translations Translations, ip []model.Location) error {

	var log *redis.Cmd
	wg := sync.WaitGroup{}
	wg.Add(16000)
	nb_iplocations := len(ip) / 4000
	nb_translationsfr := len(translations.TranslationFR) / 4000
	nb_translationsen := len(translations.TranslationEN) / 4000
	nb_translationses := len(translations.TranslationES) / 4000

	for i := 0; i < 4000; i++ {
		go func(i int, nb_iplocations int, ip []model.Location) error {
			defer wg.Done()

			for j := i * nb_iplocations; j < (i+1)*nb_iplocations; j++ {
				log = c.rdb.Do(ctx, "HSET", "location", ip[j].Address, ip[j].UUID)
				fmt.Println(log)
			}
			return nil
		}(i, nb_iplocations, ip)
		go func(i int, nb_translationsfr int, translations Translations) error {
			defer wg.Done()
			translationsFR := translations.TranslationFR

			for j := i * nb_translationsfr; j < (i+1)*nb_translationsfr; j++ {
				mTransFR, _ := json.Marshal(translationsFR[j])
				log = c.rdb.Do(ctx, "HSET", "translationFR", translationsFR[j].UUID, mTransFR)
				fmt.Println(log)
			}
			return nil
		}(i, nb_translationsfr, translations)
		go func(i int, nb_translationses int, translations Translations) error {
			defer wg.Done()
			translationsES := translations.TranslationES

			for j := i * nb_translationses; j < (i+1)*nb_translationses; j++ {
				mTransES, _ := json.Marshal(translationsES[j])
				log = c.rdb.Do(ctx, "HSET", "translationES", translationsES[j].UUID, mTransES)
				fmt.Println(log)
			}
			return nil
		}(i, nb_translationses, translations)
		go func(i int, nb_translationsen int, translations Translations) error {
			defer wg.Done()
			translationsEN := translations.TranslationEN

			for j := i * nb_translationsen; j < (i+1)*nb_translationsfr; j++ {
				mTransEN, _ := json.Marshal(translationsEN[j])
				log = c.rdb.Do(ctx, "HSET", "translationEN", translationsEN[j].UUID, mTransEN)
				fmt.Println(log)
			}
			return nil
		}(i, nb_translationsen, translations)
	}

	wg.Wait()
	return nil
}
