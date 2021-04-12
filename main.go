package main

import (
	"HDTwG/internal/network"
	"HDTwG/internal/network/http"
	"HDTwG/internal/store"
	"context"
	"log"

	"github.com/gin-gonic/gin"
)

func initRoute(router *gin.Engine, client *store.Client, redisClient *store.ClientNoSQL, cacheClient *store.CacheClient) {
	router.GET("/location", http.GetLocation(network.Get(
		cacheClient,
		client,
		//redisClient,
	)))
	router.PUT("/locations", http.PutLocation(network.Put(
		cacheClient,
		client,
		//redisClient,
	)))
}

var ctx = context.Background()

func main() {

	clt, err := store.NewSQLClient()
	if err != nil {
		log.Fatal(err)
	}

	rclt := store.NewNSQLClient()
	cch := store.NewCacheClient()

	if err := rclt.Init(); err != nil {
		log.Fatal("error while init redis client")
	}
	if err := cch.Init(); err != nil {
		log.Fatal("error while init cache client")
	}
	router := gin.Default()

	initRoute(router, clt, rclt, cch)

	err = router.Run(":8081")
	if err != nil {
		log.Fatal("error while running the router")
	}
}
