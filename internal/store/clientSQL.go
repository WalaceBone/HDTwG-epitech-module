package store

import (
	"HDTwG/model"
	"context"
	"fmt"
	"sync"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Client struct {
	db *gorm.DB
}

func NewClient() *Client {
	return &Client{}
}

func (c *Client) Init() error {

	dsn := "host=localhost user=user password=password dbname=db port=5432 sslmode=disable TimeZone=Europe/Paris"
	c.db, _ = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	err := c.db.AutoMigrate(&model.Translation{}, &model.TranslationES{}, &model.TranslationEN{}, &model.Location{})
	if err != nil {
		return err
	}
	return nil
}

func (c *Client) Get(ctx context.Context, opts Options) ([]model.Translation, error) {
	var translations []model.Translation
	var translation model.Translation
	var location model.Location

	if err := c.db.Model(&model.Location{}).Where(model.Location{Address: opts.IP}).Find(&location).Error; err != nil {
		return []model.Translation{}, err
	}
	if opts.Lang == "English" || opts.Lang == "" {
		if err := c.db.Model(&model.TranslationEN{}).Where(model.TranslationEN{UUID: location.UUID}).Find(&translation).Error; err != nil {
			return []model.Translation{}, err
		}
		translations = append(translations, translation)
	} else if opts.Lang == "French" || opts.Lang == "" {
		if err := c.db.Model(&model.Translation{}).Where(model.Translation{UUID: location.UUID}).Find(&translation).Error; err != nil {
			return []model.Translation{}, err
		}
		translations = append(translations, translation)
	} else if opts.Lang == "Spanish" || opts.Lang == "" {
		if err := c.db.Model(&model.TranslationES{}).Where(model.TranslationES{UUID: location.UUID}).Find(&translation).Error; err != nil {
			return []model.Translation{}, err
		}
		translations = append(translations, translation)
	}
	return translations, nil
}

func (c *Client) Put(ctx context.Context, translations Translations, locations []model.Location) error {
	wg := sync.WaitGroup{}
	wg.Add(50)
	portion := len(locations) / 50

	for i := 0; i < 50; i++ {
		go func(i int, portion int, locations []model.Location) {
			defer wg.Done()
			for j := i * portion; j < (i+1)*portion; j++ {
				err := c.db.Create(&locations[j]).Error
				if err != nil {
					fmt.Println(err)
				}
			}
		}(i, portion, locations)
	}

	for i := 0; i < 50; i++ {
		go func(i int, translations Translations, portion int) {
			defer wg.Done()
			for j := i * portion; j < (i+1)*portion; j++ {
				if err := c.db.Create(&model.Translation{
					UUID:       translations.TranslationFR[j].UUID,
					Continent:  translations.TranslationFR[j].Continent,
					Country:    translations.TranslationFR[j].Country,
					Region:     translations.TranslationFR[j].Region,
					Department: translations.TranslationFR[j].Department,
					City:       translations.TranslationFR[j].City,
				}).Error; err != nil {
					fmt.Println(err)
				}
				if err := c.db.Create(&model.TranslationEN{
					UUID:       translations.TranslationEN[j].UUID,
					Continent:  translations.TranslationEN[j].Continent,
					Country:    translations.TranslationEN[j].Country,
					Region:     translations.TranslationEN[j].Region,
					Department: translations.TranslationEN[j].Department,
					City:       translations.TranslationEN[j].City,
				}).Error; err != nil {
					fmt.Println(err)
				}
				if err := c.db.Create(&model.TranslationES{
					UUID:       translations.TranslationES[j].UUID,
					Continent:  translations.TranslationES[j].Continent,
					Country:    translations.TranslationES[j].Country,
					Region:     translations.TranslationES[j].Region,
					Department: translations.TranslationES[j].Department,
					City:       translations.TranslationES[j].City,
				}).Error; err != nil {
					fmt.Println(err)
				}
			}
		}(i, translations, len(translations.TranslationFR)/50)
	}

	wg.Wait()
	return nil
}
