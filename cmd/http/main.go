package main

import (
	"errors"
	"log/slog"
	"os"

	"github.com/joho/godotenv"
	handler "github.com/lMikadal/POS_Backend_golang.git/internal/adapter/handler/http"
	repository "github.com/lMikadal/POS_Backend_golang.git/internal/adapter/repository/postgres"
	"github.com/lMikadal/POS_Backend_golang.git/internal/core/domain"
	"github.com/lMikadal/POS_Backend_golang.git/internal/core/seeds"
	"github.com/lMikadal/POS_Backend_golang.git/internal/core/service"
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

func main() {
	appName := os.Getenv("APP_NAME")
	env := os.Getenv("APP_ENV")
	dbConn := os.Getenv("DB_CONNECTION")
	httpUrl := os.Getenv("HTTP_URL")
	httpPort := os.Getenv("HTTP_PORT")
	listenAddr := httpUrl + ":" + httpPort

	slog.Info("Starting the application", "app", appName, "env", env)

	// Init database
	db, err := repository.NewDB()
	if err != nil {
		slog.Error("Error initializing database connection", "error", err)
		os.Exit(1)
	}
	slog.Info("Successfully connected to the database", "db", dbConn)

	var migrate = []interface{}{
		&domain.Role{},
		&domain.User{},
		&domain.Tag{},
		&domain.Goods{},
		&domain.Price{},
	}

	if err = db.AutoMigrate(migrate...); err == nil && db.Migrator().HasTable(&domain.Role{}) {
		if err := db.First(&domain.Role{}).Error; errors.Is(err, gorm.ErrRecordNotFound) {
			seeds.SeedRole(db.DB)
		}
		if err := db.First(&domain.User{}).Error; errors.Is(err, gorm.ErrRecordNotFound) {
			seeds.SeedUser(db.DB)
		}
	}

	// Role
	roleRepo := repository.NewRoleRepository(db)
	roleService := service.NewRoleService(roleRepo)
	roleHander := handler.NewRoleHandler(roleService)

	// User
	userRepo := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepo)
	userHandler := handler.NewUserHandler(userService)

	// Tag
	tagRepo := repository.NewTagRepository(db)
	tagService := service.NewTagService(tagRepo)
	tagHandler := handler.NewTagHandler(tagService)

	// Goods
	goodsRepo := repository.NewGoodsRepository(db)
	goodsService := service.NewGoodsService(goodsRepo)
	goodsHandler := handler.NewGoodsHandler(goodsService)

	// Price
	priceRepo := repository.NewPriceRepository(db)
	priceService := service.NewPriceService(priceRepo)
	priceHandler := handler.NewPriceHandler(priceService)

	// Init router
	router, err := handler.NewRouter(
		*roleHander,
		*userHandler,
		*tagHandler,
		*goodsHandler,
		*priceHandler,
	)
	if err != nil {
		slog.Error("Error initializing router", "error", err)
		os.Exit(1)
	}

	// Start server
	slog.Info("Starting the HTTP server", "listen_address", listenAddr)
	err = router.Serve(listenAddr)
	if err != nil {
		slog.Error("Error starting the HTTP server", "error", err)
		os.Exit(1)
	}
}
