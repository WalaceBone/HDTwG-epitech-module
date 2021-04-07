package main

import (
	"HDTwG/internal/network"
	"HDTwG/internal/network/http"
	"HDTwG/internal/store"
	"context"
	"log"

	"github.com/gin-gonic/gin"
)

func initRoute(router *gin.Engine, client *store.Client, redisClient *store.ClientNoSQL) {
	router.GET("/location", http.GetLocation(network.Get(redisClient)))
	router.PUT("/locations", http.PutLocation(network.Put(redisClient)))
}

var ctx = context.Background()

func main() {

	clt := store.NewClient()
	rclt := store.NewNSQLClient()

	if err := clt.Init(); err != nil {
		log.Fatal("error while init postgres client")
	}
	if err := rclt.Init(); err != nil {
		log.Fatal("error while init redis client")
	}
	router := gin.Default()

	initRoute(router, clt, rclt)

	err := router.Run(":8081")
	if err != nil {
		log.Fatal("error while running the router")
	}
}
