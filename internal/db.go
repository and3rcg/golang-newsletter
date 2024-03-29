package internal

import (
	"log"
	"net/mail"
	"newsletter-go/models"
	"os"

	"github.com/go-playground/validator/v10"
	"github.com/mailersend/mailersend-go"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Repository struct {
	DB        *gorm.DB
	Validator *validator.Validate
	MS        *mailersend.Mailersend
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
	db.AutoMigrate(&models.Newsletter{}, &models.NewsletterUser{})

	repoObj := &Repository{DB: db}

	log.Println("Database connected successfully!")

	log.Println("Starting validator instance...")
	repoObj.Validator = validator.New()

	repoObj.Validator.RegisterValidation("valid_email", func(f1 validator.FieldLevel) bool {
		_, err := mail.ParseAddress(f1.Field().String())
		return err == nil
	})

	log.Println("Starting MailerSend library...")
	ms := mailersend.NewMailersend(os.Getenv("MAILERSEND_API_KEY"))

	repoObj.MS = ms

	return repoObj, nil
}
