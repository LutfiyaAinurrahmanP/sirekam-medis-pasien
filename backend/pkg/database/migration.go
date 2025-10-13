package database

import (
	"fmt"
	"log"

	"gorm.io/gorm"
)

type Migrator struct {
	db *gorm.DB
}

func NewMigrator(db *gorm.DB) *Migrator {
	return &Migrator{
		db: db,
	}
}

// Safe migration (membuat tabel baru jika belum ada dan tidak menghapus kolom yang sudah ada)
func (m *Migrator) AutoMigrate(models ...interface{}) error {
	log.Println("🔄 Starting database migration...")
	for _, model := range models {
		modelName := fmt.Sprintf("%T", model)
		log.Printf("   - Migrating: %s", modelName)
	}

	// Gorm AutoMigrate handle semua models sekaligus
	if err := m.db.AutoMigrate(models...); err != nil {
		return fmt.Errorf("failed to migrate database: %w", err)
	}
	log.Println("✅ Database migration completed successfully")
	return nil
}

func (m *Migrator) RunMigrations(models ...interface{}) error {
	// Auto migrate semua models
	if err := m.AutoMigrate(models...); err != nil {
		return err
	}
	return nil
}

// Drop all tables
func (m *Migrator) DropAllTables(models ...interface{}) error {
	log.Println("⚠️  WARNING: Dropping all tables...")

	// DropTable akan menghapus table jika ada
	migrator := m.db.Migrator()

	for _, model := range models {
		// Cek apakah table ada
		if migrator.HasTable(model) {
			modelName := fmt.Sprintf("%T", model)
			log.Printf("   - Dropping table: %s", modelName)

			// Drop table
			if err := migrator.DropTable(model); err != nil {
				return fmt.Errorf("failed to drop table %s: %w", modelName, err)
			}
		}
	}
	log.Println("✅ All tables dropped successfully")
	return nil
}

// Menampilkan status migrations
func (m *Migrator) GetMigrationStatus(models ...interface{}) {
	log.Println("📊 Migration Status:")

	migrator := m.db.Migrator()

	for _, model := range models {
		modelName := fmt.Sprintf("%T", model)

		// Cek apakah tabel sudah ada
		if migrator.HasTable(model) {
			log.Printf("   ✅ %s - Table exists", modelName)
		} else {
			log.Printf("   ❌ %s - Table not found", modelName)
		}
	}
}
