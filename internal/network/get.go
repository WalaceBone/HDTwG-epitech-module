package network

import (
	Stores "HDTwG/internal/store"
	"HDTwG/model"
	"context"
)

type GetCmd func(ctx context.Context, options Stores.Options) ([]model.Translation, error)

func Get(stores ...Stores.Store) GetCmd {
	return func(ctx context.Context, options Stores.Options) ([]model.Translation, error) {
		var translation []model.Translation
		var err error

		for _, store := range stores {
			translation, err = store.Get(ctx, options)
			if err != nil {
				if err != model.ErrTranslationNotFound {
					return []model.Translation{}, err
				}
			}
			if translation != nil {
				stores[0].Put(ctx, Stores.Translations{
					TranslationFR: nil,
					TranslationEN: nil,
					TranslationES: nil,
				}, []model.Location{
				})

				return translation, nil
			}
			if translation != nil {
				break
			}
		}
		if translation == nil {
			return []model.Translation{}, err
		}

		return translation, nil
	}
}
