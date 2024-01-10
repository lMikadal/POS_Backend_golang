package repository

import (
	"context"
	"fmt"
	"os"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DB struct {
	*gorm.DB
}

// NewDB creates a new PostgreSQL database instance
func NewDB(ctx context.Context) (*DB, error) {
	dsn := fmt.Sprintf("host=%v user=%v password=%v dbname=%v port=%v sslmode=disable",
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
		return nil, err
	}

	return &DB{
		db,
	}, nil
}
