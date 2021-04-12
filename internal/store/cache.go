package store

import (
	"HDTwG/model"
	"context"
	"log"
	"time"

	gocache "github.com/patrickmn/go-cache"
)

type Cache struct {
	locations     []model.Location
	translationFR []model.Translation
	translationEN []model.TranslationEN
	translationES []model.TranslationES
}

const TimeToLive = 40320

type Translation struct {
	TranslationFR model.Translation
	TranslationEN model.Translation
	TranslationES model.Translation
}

type CacheClient struct {
	cache *gocache.Cache
}

func NewCacheClient() *CacheClient {
	return &CacheClient{}
}

func (c *CacheClient) Init() error {
	c.cache = gocache.New(100*time.Minute, TimeToLive)
	return nil
}

func (c *CacheClient) Insert(ctx context.Context, ip model.Location, translation model.Translation) error {
	c.cache.Set(ip.Address, &translation, TimeToLive)
	return nil
}

func (c *CacheClient) Get(ctx context.Context, opts Options) ([]model.Translation, error) {

	log.Print(opts)
	result, found := c.cache.Get(opts.IP)
	if !found {
		return nil, nil
	}
	return result.([]model.Translation), nil
}

func (c *CacheClient) Put(ctx context.Context, translations Translations, ip []model.Location) error {
	return nil
}

//TODO
// Cache Read
// Cache Write
// bigcache
// choisir comportement du cache
// choisir
