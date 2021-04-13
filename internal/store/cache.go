package store

import (
	"HDTwG/model"
	"context"
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


//Tres boulgour ce qui se passe ici mais on en parle pas
func (c *CacheClient) Insert(ctx context.Context, ip model.Location, translation []model.Translation) error {

	var key string

	key = ip.Address + ip.UUID

	c.cache.Set(key, &translation, TimeToLive)
	return nil
}

func (c *CacheClient) Get(ctx context.Context, opts Options) ([]model.Translation, error) {
	result, found := c.cache.Get(opts.IP)
	if !found {
		return nil, nil
	}
	return result.([]model.Translation), nil
}

func (c *CacheClient) Put(ctx context.Context, translations Translations, ip []model.Location) error {
	c.Insert(ctx, ip[0], translations.TranslationFR)
	return nil
}

//TODO
// Cache Read
// Cache Write
// bigcache
// choisir comportement du cache
// choisir
