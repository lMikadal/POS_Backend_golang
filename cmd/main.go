package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"log/slog"
	"os"
	"time"

	"github.com/joho/godotenv"
	repository "github.com/lMikadal/POS_Backend_golang.git/internal/adapter/repository/postgres"
	"github.com/lMikadal/POS_Backend_golang.git/seeds"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func init() {
	// Init logger
	var logHandler *slog.JSONHandler

	env := os.Getenv("APP_ENV")
	if env == "production" {
		logHandler = slog.NewJSONHandler(os.Stdout, nil)
	} else {
		logHandler = slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
			Level: slog.LevelDebug,
		})

		// Load	.env file
		err := godotenv.Load()
		if err != nil {
			slog.Error("Error loading .env file", "error", err)
			os.Exit(1)
		}
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
	dsn := fmt.Sprintf("host=%v user=%v password=%v dbname=%v port=%v",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
	)

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
	dbConn := os.Getenv("DB_CONNECTION")

	slog.Info("Starting the application", "app", appName, "env", env)

	// Init database
	ctx := context.Background()
	db, err := repository.NewDB(ctx)
	if err != nil {
		slog.Error("Error initializing database connection", "error", err)
		os.Exit(1)
	}
	slog.Info("Successfully connected to the database", "db", dbConn)

	db.AutoMigrate(migrate...)
	// if err = db.AutoMigrate(migrate...); err == nil && (db.Migrator().HasTable(&repository.User{}) || db.Migrator().HasTable(&repository.Role{})) {
	// 	if err := db.First(&repository.Role{}).Error; errors.Is(err, gorm.ErrRecordNotFound) {
	// 		seeds.SeedRole(db)
	// 	}
	// 	if err := db.First(&repository.User{}).Error; errors.Is(err, gorm.ErrRecordNotFound) {
	// 		seeds.SeedUser(db)
	// 	}
	// }

	// router := gin.Default()
	// api1 := router.Group("/api/v1")

	// routers.RouterUser(db, api1)
	// routers.RouterRole(db, api1)

	// router.Run(":8080")
}
