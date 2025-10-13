package database

import (
	"fmt"
	"log"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// NewMySQLConnection membuat koneksi ke database MySQL menggunakan GORM
// Parameter dsn adalah Data Source Name (connection string)
// Return *gorm.DB yang akan digunakan untuk query database
func NewMySQLConnection(dsn string) (*gorm.DB, error) {
	// Konfigurasi GORM
	// logger.Info: menampilkan semua SQL query yang dijalankan (berguna untuk debugging)
	// SlowThreshold: query yang lebih lama dari 200ms akan ditandai sebagai slow query
	gormConfig := &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
		NowFunc: func() time.Time {
			// Waktu lokal untuk timestamp
			return time.Now().Local()
		},
	}

	// Membuka koneksi ke MySQL menggunakan driver mysql dan konfigurasi GORM
	db, err := gorm.Open(mysql.Open(dsn), gormConfig)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	// Mendapatkan generic database object untuk konfigurasi connection pool
	sqlDB, err := db.DB()
	if err != nil {
		return nil, fmt.Errorf("failed to get database instance: %w", err)
	}

	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)

	if err := sqlDB.Ping(); err != nil {
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}

	log.Println("Successfully connected to MySQL database")

	return db, nil
}
