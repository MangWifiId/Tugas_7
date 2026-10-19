package database

import (
	"context"
	"errors"
	"os"
	"time"

	_ "github.com/jackc/pgx/v5/stdlib"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"github.com/joho/godotenv"
)

func Connect() (*gorm.DB, error) {
	// Membaca konfigurasi dari file .env
	if err := godotenv.Load(); err != nil {
		return nil, errors.New("file .env tidak ditemukan atau gagal dibaca")
	}
	
	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		return nil, errors.New("environment variable DATABASE_URL belum diatur")
	}

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger:               logger.Default.LogMode(logger.Info),
		DisableAutomaticPing: true,
	})
	if err != nil {
		return nil, err
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	if err := sqlDB.PingContext(ctx); err != nil {
		sqlDB.Close()
		return nil, err
	}

	return db, nil
}

// func Connect() (*gorm.DB, error) {
// 	dsn := os.Getenv("DATABASE_URL")
// 	// if err := godotenv.Load(); err != nil {
// 	// 	return nil, fmt.Errorf("gagal membaca file .env: %w", err)
// 	// }

// 	dsn := os.Getenv("DATABASE_URL")
// 	if dsn == "" {
// 		return nil, errors.New("environment variable DATABASE_URL belum diatur")
// 	}

// 	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
// 	if err != nil {
// 		return nil, err
// 	}

// 	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
// 	defer cancel()

// 	if err := db.PingContext(ctx); err != nil {
// 		db.Close()
// 		return nil, fmt.Errorf("gagal terhubung ke PostgreSQL: %w", err)
// 	}

// 	return db, nil
// }
