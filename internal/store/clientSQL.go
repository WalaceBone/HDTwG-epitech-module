package store

import (
	"HDTwG/model"
	"context"
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

	err := c.db.AutoMigrate(&model.Translation{}, &model.Location{})
	if err != nil {
		return err
	}
	return nil
}

func (c *Client) Get(ctx context.Context, opts Options) (model.Translation, error) {
	return model.Translation{}, nil
}

func (c *Client) Put(ctx context.Context) error {

	//load conf file
	//look for download
	//parse
	//update

	return nil
}
