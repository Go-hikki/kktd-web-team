package database

import (
	"fmt"
	"os"

	"github.com/Rajil1213/go-admin/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	db_username := os.Getenv("DB_USERNAME")
	db_host := os.Getenv("DB_HOST")
	db_password := os.Getenv("DB_PASSWORD")
	db_database := os.Getenv("DB_DATABASE")
	db_port := os.Getenv("DB_PORT") // <-- удобно задавать порт через переменную

	if db_port == "" {
		db_port = "5432" // по умолчанию
	}

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		db_host, db_username, db_password, db_database, db_port)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(fmt.Sprintf("Could not connect to the database: %v", err))
	}

	DB = db

	// Автоматическая миграция таблиц
	err = db.AutoMigrate(&models.User{}, &models.Role{}, &models.Permission{}, &models.Product{}, &models.Order{}, &models.OrderItem{})
	if err != nil {
		panic(fmt.Sprintf("Migration error: %v", err))
	}
}
