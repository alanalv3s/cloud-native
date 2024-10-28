package database

import (
	"fmt"
	"log"
	"os"

	"github.com/alanalv3s/cloud-native/go-api/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

// InitDB initializes the Postgres database connection using GORM.
func InitDB() {
	dsn := getDSN()
	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// Enable uuid-ossp extension for UUID generation
	if err := enableUUIDExtension(); err != nil {
		log.Fatalf("Failed to enable uuid-ossp extension: %v", err)
	}

	// Auto-migrate models
	if err := DB.AutoMigrate(&models.User{}); err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}
}

// enableUUIDExtension enables the 'uuid-ossp' extension in PostgreSQL.
func enableUUIDExtension() error {
	// Executes the raw SQL query to enable the uuid-ossp extension
	return DB.Exec("CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\"").Error
}

func getDSN() string {
	host := os.Getenv("DB_HOST")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")
	port := os.Getenv("DB_PORT")
	sslmode := os.Getenv("DB_SSLMODE")
	return fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		host, user, password, dbname, port, sslmode)
}
