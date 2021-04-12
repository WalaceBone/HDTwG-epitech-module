package store

import (
	"HDTwG/model"
	"context"
)

type Options struct {
	IP   string `json:"network"`
	Lang string `json:"lang"`
}

type Translations struct {
	TranslationFR []model.Translation
	TranslationEN []model.Translation
	TranslationES []model.Translation
}

type Store interface {
	Get(ctx context.Context, opts Options) ([]model.Translation, error)
	Put(ctx context.Context, translations Translations, ip []model.Location) error
}