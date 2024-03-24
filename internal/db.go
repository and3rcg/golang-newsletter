package internal

import (
	"log"
	"newsletter-go/models"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Repository struct {
	DB *gorm.DB
	// Redis
}

func StartDB() (*Repository, error) {

	log.Println("Connecting to database...")
	db, err := gorm.Open(sqlite.Open("database.db"), &gorm.Config{})
	if err != nil {
		log.Println("Failed to connect to database")
		return nil, err
	}

	log.Println("Running migrations...")
	db.AutoMigrate(&models.Newsletter{})

	repoObj := &Repository{DB: db}

	log.Println("Database started successfully!")
	return repoObj, nil
}
