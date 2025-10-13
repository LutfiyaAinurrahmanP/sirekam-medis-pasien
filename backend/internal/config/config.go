package config

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
)

// Struct untuk seluruh akses konfigurasi aplikasi
type Config struct {
	AppName    string
	AppEnv     string
	AppPort    string
	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBName     string
	JWTSecret  string
	JWTExpire  string
}

// LoadConfig membaci env variables dan mengembalikan ke Config struct
// Fungsi yang dipanggil pertama kali saat aplikasi dijalankan
func LoadConfig() (*Config, error) {
	// Load file .env kedalam environtment variables
	// godotenv.Load() akan membaca file .env dan set sebagai env variables
	if err := loadEnvFile(); err != nil {
		// Jika .env tidak ditemukan, coba gunakan env variables yang sudah ada
		// Ini berguna untuk production environment (Docker, K8s, dll)
		fmt.Println("⚠️  Warning: .env file not found, using system environment variables")
	}

	// Membuat instance Config
	// os.Getenv() untuk membaca nilai environment variable
	config := &Config{
		AppName:    os.Getenv("APP_NAME"),
		AppEnv:     os.Getenv("APP_ENV"),
		AppPort:    os.Getenv("APP_PORT"),
		DBHost:     os.Getenv("DB_HOST"),
		DBPort:     os.Getenv("DB_PORT"),
		DBUser:     os.Getenv("DB_USER"),
		DBPassword: os.Getenv("DB_PASSWORD"),
		DBName:     os.Getenv("DB_NAME"),
		JWTSecret:  os.Getenv("JWT_SECRET"),
		JWTExpire:  os.Getenv("JWT_EXPIRE"),
	}

	return config, nil
}

func loadEnvFile() error {
	// List path yang akan dicoba untuk mencari .env file
	// Urutan: dari yang paling specific ke yang paling general
	envPaths := []string{
		".env",                    // Current directory (saat run dari backend/)
		"../.env",                 // Parent directory (root folder)
		"../../.env",              // 2 level up (jika run dari backend/cmd/server/)
		filepath.Join(getRootPath(), ".env"), // Root project path
	}

	// Coba load dari setiap path
	for _, path := range envPaths {
		if _, err := os.Stat(path); err == nil {
			// File ditemukan, load .env
			if err := godotenv.Load(path); err == nil {
				fmt.Printf("✅ Loaded .env from: %s\n", path)
				return nil
			}
		}
	}

	// Jika tidak ditemukan, return error
	return fmt.Errorf(".env file not found in any of the expected locations")
}

func getRootPath() string {
	// Mulai dari current directory
	dir, err := os.Getwd()
	if err != nil {
		return "."
	}

	// Traverse up untuk mencari go.mod
	for {
		// Check jika go.mod ada di directory ini
		if _, err := os.Stat(filepath.Join(dir, "go.mod")); err == nil {
			// go.mod ditemukan, naik 1 level ke parent (root folder)
			return filepath.Dir(dir)
		}

		// Naik 1 level
		parent := filepath.Dir(dir)
		
		// Jika sudah sampai root filesystem, stop
		if parent == dir {
			break
		}
		
		dir = parent
	}

	// Fallback ke current directory
	return "."
}

// GetDSN menghasilkan Data Source Name untuk koneksi MySQL
// DSN adalah string koneksi yang berisi info host, port, user, password, dan database
func (c *Config) GetDSN() string {
	// Format DSN untuk MySQL: user:password@tcp(host:port)/dbname?charset=utf8mb4&parseTime=True&loc=Local
	// parseTime=True: mengkonversi DATE/DATETIME dari MySQL ke time.Time di Go
	// charset=utf8mb4: support untuk emoji dan karakter unicode lengkap
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		c.DBUser,
		c.DBPassword,
		c.DBHost,
		c.DBPort,
		c.DBName,
	)
}
