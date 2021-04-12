package main

import (
	"HDTwG/internal/network"
	"HDTwG/internal/network/http"
	"HDTwG/internal/store"
	"context"
	"log"

	"github.com/gin-gonic/gin"
)

func initRoute(router *gin.Engine, client store.Store, redisClient *store.ClientNoSQL) {
	router.GET("/location", http.GetLocation(network.Get(client)))
	router.PUT("/locations", http.PutLocation(network.Put(client)))
}

var ctx = context.Background()

func main() {

	var client store.Store
	client, err := store.NewSQLClient()
	if err != nil {
		log.Fatal(err)
	}

	rclt := store.NewNSQLClient()
	if err := rclt.Init(); err != nil {
		log.Fatal("error while init redis client")
	}
	router := gin.Default()

	initRoute(router, client, rclt)

	err = router.Run(":8081")
	if err != nil {
		log.Fatal("error while running the router")
	}
}
