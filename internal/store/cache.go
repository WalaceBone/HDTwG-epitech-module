package store

import "HDTwG/model"

type Cache struct {
	locations     []model.Location
	translationFR []model.Translation
	translationEN []model.TranslationEN
	translationES []model.TranslationES
}
