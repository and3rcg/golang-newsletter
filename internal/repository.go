package internal

import (
	"errors"
	"fmt"
	"log"
	"net/mail"
	"newsletter-go/models"
	"os"

	"github.com/go-playground/validator/v10"
	"github.com/mailersend/mailersend-go"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Repository struct {
	DB        *gorm.DB
	Validator *validator.Validate
	MS        *mailersend.Mailersend
}

var repository *Repository

// I had to separate the initialization methods because I was getting null pointer exception when I passed
// the repository directly to the worker. so, in order to avoid starting up the database all over again,
// I separated the methods to start up only the MailerSend instance
func StartRepository() (*Repository, error) {
	db, err := StartDB()
	if err != nil {
		log.Println("Failed to start database")
		return nil, err
	}
	ms, err := StartMailerSendInstance()
	if err != nil {
		log.Println("Failed to start Mailersend instance")
		return nil, err
	}
	v, err := StartValidator()
	if err != nil {
		log.Println("Failed to start validator")
		return nil, err
	}

	// starting the repository
	repository = &Repository{
		DB:        db,
		MS:        ms,
		Validator: v,
	}

	return repository, nil
}

func GetRepository() *Repository {
	return repository
}

func StartDB() (*gorm.DB, error) {
	log.Println("Connecting to database...")

	// refer https://github.com/go-sql-driver/mysql#dsn-data-source-name for details on DSN
	dsn := fmt.Sprintf("%s:%s@tcp(db:3306)/%s", os.Getenv("MYSQL_USER"), os.Getenv("MYSQL_PASSWORD"), os.Getenv("MYSQL_DATABASE"))

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Println("Failed to connect to database")
		return nil, err
	}

	log.Println("Running migrations...")
	db.AutoMigrate(&models.Newsletter{}, &models.NewsletterUser{})
	log.Println("Database started successfully!")

	return db, nil
}

func StartValidator() (*validator.Validate, error) {
	log.Println("Starting validator instance...")
	v := validator.New()

	err := v.RegisterValidation("valid_email", func(f1 validator.FieldLevel) bool {
		_, validationErr := mail.ParseAddress(f1.Field().String())
		return validationErr == nil
	})
	if err != nil {
		return nil, err
	}

	return v, nil
}

func StartMailerSendInstance() (*mailersend.Mailersend, error) {
	log.Println("Starting MailerSend library...")

	ms := mailersend.NewMailersend(os.Getenv("MAILERSEND_API_KEY"))
	if ms == nil {
		return nil, errors.New("failed to create MailerSend instance")
	}

	return ms, nil
}
