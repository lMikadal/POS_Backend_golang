package main

import (
	"fmt"
	"log/slog"
	"os"

	"github.com/joho/godotenv"
	repository "github.com/lMikadal/POS_Backend_golang.git/internal/adapter/repository/postgres"
	"github.com/lMikadal/POS_Backend_golang.git/internal/core/service"
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

func main() {
	appName := os.Getenv("APP_NAME")
	env := os.Getenv("APP_ENV")
	dbConn := os.Getenv("DB_CONNECTION")

	slog.Info("Starting the application", "app", appName, "env", env)

	// Init database
	db, err := repository.NewDB()
	if err != nil {
		slog.Error("Error initializing database connection", "error", err)
		os.Exit(1)
	}
	slog.Info("Successfully connected to the database", "db", dbConn)

	var migrate = []interface{}{
		// &repository.Role{},
		// &repository.User{},
	}

	db.AutoMigrate(migrate...)

	// User
	userRepo := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepo)

	fmt.Println(userService.GetUserById(1))

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
