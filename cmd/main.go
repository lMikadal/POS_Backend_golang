package main

import (
	"errors"
	"fmt"
	"log"
	"log/slog"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	repository "github.com/lMikadal/POS_Backend_golang.git/repositories"
	"github.com/lMikadal/POS_Backend_golang.git/routers"
	"github.com/lMikadal/POS_Backend_golang.git/seeds"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func init() {
	// setting log
	var logHandler *slog.JSONHandler

	logHandler = slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelDebug,
	})

	// init env
	err := godotenv.Load()
	if err != nil {
		slog.Error("Error loading .env file", "error", err)
		os.Exit(1)
	}

	logger := slog.New(logHandler)
	slog.SetDefault(logger)
}

var migrate = []interface{}{
	&repository.Role{},
	&repository.User{},
	// &models.Goods{},
	// &models.Tag{},
	// &models.GoodsPrice{},
	// &models.Order{},
	// &models.Status{},
	// &models.Stock{},
}

func initDB() (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%v user=%v password=%v dbname=%v port=%v", os.Getenv("DB_HOST"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"), os.Getenv("DB_PORT"))

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		NowFunc: func() time.Time {
			t, _ := time.LoadLocation("Asia/Jakarta")
			return time.Now().In(t)
		},
	})
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	if err = db.AutoMigrate(migrate...); err == nil && (db.Migrator().HasTable(&repository.User{}) || db.Migrator().HasTable(&repository.Role{})) {
		if err := db.First(&repository.Role{}).Error; errors.Is(err, gorm.ErrRecordNotFound) {
			seeds.SeedRole(db)
		}
		if err := db.First(&repository.User{}).Error; errors.Is(err, gorm.ErrRecordNotFound) {
			seeds.SeedUser(db)
		}
	}

	return db, nil
}

func main() {
	appName := os.Getenv("APP_NAME")
	env := os.Getenv("APP_ENV")

	slog.Info("Starting the application", "app", appName, "env", env)

	db, err := initDB()
	if err != nil {
		log.Fatal(err)
	}

	router := gin.Default()
	api1 := router.Group("/api/v1")

	routers.RouterUser(db, api1)
	routers.RouterRole(db, api1)

	router.Run(":8080")
}
