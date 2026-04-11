package db

import (
	"fmt"
	"log"

	"clientcompany/config"
	"clientcompany/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		config.App.DBHost, config.App.DBPort,
		config.App.DBUser, config.App.DBPassword, config.App.DBName,
	)
	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}
	log.Println("database connected")
	migrate()
}

func migrate() {
	if err := DB.AutoMigrate(&models.User{}, &models.Session{}); err != nil {
		log.Fatalf("migration failed: %v", err)
	}
}
