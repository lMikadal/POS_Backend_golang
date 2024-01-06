package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/lMikadal/POS_Backend_golang.git/models"
	"github.com/lMikadal/POS_Backend_golang.git/routers"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}
}

var migrate = []interface{}{
	&models.User{},
	&models.Role{},
	&models.Goods{},
	&models.GoodsPrice{},
	&models.Tag{},
	&models.Order{},
	&models.Status{},
	&models.Stock{},
}

func initDB() (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%v user=%v password=%v dbname=%v port=%v", os.Getenv("DB_HOST"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"), os.Getenv("DB_PORT"))

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	// db.AutoMigrate(migrate...)

	return db, nil
}

func main() {
	db, err := initDB()
	if err != nil {
		log.Fatal(err)
	}

	router := gin.Default()
	api1 := router.Group("/api/v1")

	routers.SetCollectionRoutes(api1, db)

	router.Run(":8080")
}
