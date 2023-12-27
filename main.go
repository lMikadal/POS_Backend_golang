package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/lMikadal/POS_Backend_golang.git/routers"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}
}

func main() {

	router := gin.Default()
	api := router.Group("/api/v1")

	routers.SetCollectionRoutes(api)

	router.Run(":8080")

}
