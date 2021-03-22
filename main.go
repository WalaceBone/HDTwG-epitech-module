package main

import (
	"HDTwG/internal/network"
	"HDTwG/internal/network/http"
	"HDTwG/internal/store"
	"github.com/gin-gonic/gin"
	"log"
)

func initRoute(router *gin.Engine, client *store.Client) {
	router.GET("/location", http.GetLocation(network.Get(client)))
}

func main() {

	clt := store.NewClient()

	if err := clt.Init(); err != nil {
		log.Fatal("error while init client")
	}
	router := gin.Default()

	initRoute(router, clt)

	err := router.Run(":8081")
	if err != nil {
		log.Fatal("error while running the router")
	}
}
