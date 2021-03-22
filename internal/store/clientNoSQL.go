package store

import (
	"HDTwG/model"
	"context"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type ClientNoSQL struct {
	db *gorm.DB

}

func NewNSQLClient() *ClientNoSQL {
	return &ClientNoSQL{}
}

func (c *ClientNoSQL) Init() error {
//TODO Change later for nosql
	dsn := "host=localhost user=user password=password dbname=db port=5432 sslmode=disable TimeZone=Europe/Paris"
	c.db, _ = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	return nil
}

func (c *ClientNoSQL) Get(ctx context.Context, opts Options) ([]model.Translation, error) {
	return nil, nil
}

func (c *ClientNoSQL) Put(ctx context.Context) error {

	//load conf file
	//look for download
	//parse
	//update

	return nil
}
