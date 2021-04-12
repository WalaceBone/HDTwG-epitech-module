package network

import (
	Stores "HDTwG/internal/store"
	"HDTwG/model"
	"context"
	"log"
	"reflect"
)

type GetCmd func(ctx context.Context, options Stores.Options) ([]model.Translation, error)

func Get(stores ...Stores.Store) GetCmd {
	return func(ctx context.Context, options Stores.Options) ([]model.Translation, error) {
		var translation []model.Translation

		for _, store := range stores {
			log.Print(reflect.TypeOf(store))
			var err error
			translation, err = store.Get(ctx, options)
			if err != nil {
				//TODO error models
				if err != model.ErrTranslationNotFound {
					log.Print(err)
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
			// if translation != []model.Translation{} {
			// 	break
			// }
		}
		// if translation == []model.Translation{} {
		// 	return []model.Translation{}, err
		// }

		return translation, nil
	}
}
