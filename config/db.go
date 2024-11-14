package config

import (
	"database/sql"
	"fmt"

	"inventaris/app/category"
	"inventaris/app/items"

	_ "github.com/lib/pq" // Import for PostgreSQL support
	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) {
	db.AutoMigrate(&category.Category{})
	db.AutoMigrate(&items.Items{})
}

func InitDB() (*gorm.DB, error) {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		viper.GetString("DB_HOST"),
		viper.GetString("DB_USER"),
		viper.GetString("DB_PASSWORD"),
		viper.GetString("DB_NAME"),
		viper.GetString("DB_PORT"),
	)

	fmt.Printf("dsn: { %v }\n\n", dsn)

	// Open an SQL database connection using `lib/pq`
	sqlDB, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, fmt.Errorf("failed to open sql connection with lib/pq: %w", err)
	}

	// Bind the `*sql.DB` connection to GORM
	db, err := gorm.Open(postgres.New(postgres.Config{
		Conn: sqlDB,
	}), &gorm.Config{})

	if err != nil {
		return nil, fmt.Errorf("failed to connect to PostgreSQL: %w", err)
	}

	Migrate(db)

	return db, nil
}
