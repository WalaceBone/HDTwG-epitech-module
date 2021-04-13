package store

import (
	"HDTwG/model"
	"context"
	"fmt"

	"github.com/jmoiron/sqlx"
	// Used for the postgres driver
	_ "github.com/lib/pq"
)

type Client struct {
	db *sqlx.DB
}

func NewSQLClient() (*Client, error) {

	db, err := sqlx.Connect("postgres", fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		"localhost", 5432, "user", "password", "db"))
	if err != nil {
		return &Client{}, err
	}

	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS translation (
		uuid VARCHAR(255) PRIMARY KEY,
		continent VARCHAR(255),
		country VARCHAR(255),
		region VARCHAR(255),
		department VARCHAR(255),
		city VARCHAR(255)
	)`)
	if err != nil {
		return &Client{}, err
	}

	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS translation_en (
		uuid VARCHAR(255) PRIMARY KEY,
		continent VARCHAR(255),
		country VARCHAR(255),
		region VARCHAR(255),
		department VARCHAR(255),
		city VARCHAR(255)
	)`)
	if err != nil {
		return &Client{}, err
	}

	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS translation_es (
		uuid VARCHAR(255) PRIMARY KEY,
		continent VARCHAR(255),
		country VARCHAR(255),
		region VARCHAR(255),
		department VARCHAR(255),
		city VARCHAR(255)
	)`)
	if err != nil {
		return &Client{}, err
	}

	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS location (
		address VARCHAR(255),
		uuid VARCHAR(255)
	)`)
	if err != nil {
		return &Client{}, err
	}

	return &Client{db}, nil
}

func (c *Client) Get(ctx context.Context, opts Options) ([]model.Translation, error) {
	var translations []model.Translation
	var translation model.Translation
	var location model.Location

	err := c.db.Get(&location, "SELECT location.* FROM location WHERE address=$1", opts.IP)
	if err != nil {
		return []model.Translation{}, err
	}
	if opts.Lang == "English" || opts.Lang == "" {
		c.db.Get(&translation, "SELECT translation_en.* FROM translation_en WHERE uuid=$1", location.UUID)
		translations = append(translations, translation)
	}
	if opts.Lang == "French" || opts.Lang == "" {
		c.db.Get(&translation, "SELECT translation.* FROM translation WHERE uuid=$1", location.UUID)
		translations = append(translations, translation)
	}
	if opts.Lang == "Spanish" || opts.Lang == "" {
		c.db.Get(&translation, "SELECT translation_es.* FROM translation_es WHERE uuid=$1", location.UUID)
		translations = append(translations, translation)
	}
	fmt.Println(translations)
	return translations, nil
}

func (c *Client) Put(ctx context.Context, translations Translations, locations []model.Location) error {
	go func() {
		c.db.MustExec("DELETE FROM location")
		c.db.MustExec("copy location from '/ressources/IP-locations/IP-locations.csv' DELIMITER ',' CSV HEADER")
	}()
	go func() {
		c.db.MustExec("DELETE FROM translation")
		c.db.MustExec("copy translation from '/ressources/IP-locations/Locations-FR.csv' DELIMITER ';' CSV HEADER")
	}()
	go func() {
		c.db.MustExec("DELETE FROM translation_es")
		c.db.MustExec("copy translation_es from '/ressources/IP-locations/Locations-ES.csv' DELIMITER ';' CSV HEADER")
	}()
	go func() {
		c.db.MustExec("DELETE FROM translation_en")
		c.db.MustExec("copy translation_en from '/ressources/IP-locations/Locations-EN.csv' DELIMITER ';' CSV HEADER")
	}()
	return nil
}
