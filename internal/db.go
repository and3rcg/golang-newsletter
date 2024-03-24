package internal

import (
	"log"
	"net/mail"
	"newsletter-go/models"

	"github.com/go-playground/validator/v10"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Repository struct {
	DB        *gorm.DB
	Validator *validator.Validate
	// Redis
}

func StartRepository() (*Repository, error) {

	log.Println("Connecting to database...")
	db, err := gorm.Open(sqlite.Open("database.db"), &gorm.Config{})
	if err != nil {
		log.Println("Failed to connect to database")
		return nil, err
	}

	log.Println("Running migrations...")
	db.AutoMigrate(&models.Newsletter{})

	repoObj := &Repository{DB: db}

	log.Println("Database connected successfully!")

	log.Println("Starting validator instance...")
	repoObj.Validator = validator.New()

	repoObj.Validator.RegisterValidation("valid_email", func(f1 validator.FieldLevel) bool {
		_, err := mail.ParseAddress(f1.Field().String())
		return err == nil
	})

	return repoObj, nil
}
