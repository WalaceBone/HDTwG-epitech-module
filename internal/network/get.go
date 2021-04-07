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
				//TODO error models
				/*if err != model.ErrNotFound {
					logrus.Error(err)
				}*/
			}
			// if translation != []model.Translation{} {
			// 	break
			// }
		}
		// if translation == []model.Translation{} {
		// 	return []model.Translation{}, err
		// }
		return translation, err
	}
}
