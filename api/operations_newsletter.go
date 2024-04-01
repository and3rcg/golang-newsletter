package api

import (
	"newsletter-go/internal"
	"newsletter-go/models"

	"gorm.io/gorm"
)

func CreateNewsletterOperation(repo *internal.Repository, obj *models.Newsletter) error {
	result := repo.DB.Create(obj)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func GetNewsletterByIDOperation(repo *internal.Repository, id int) (*models.Newsletter, error) {
	var obj models.Newsletter
	result := repo.DB.Preload("Users").Where("id = ?", id).First(&obj)
	if result.Error != nil {
		return nil, result.Error
	}
	return &obj, nil
}

func GetAllNewslettersOperation(repo *internal.Repository) ([]models.Newsletter, error) {
	var objs []models.Newsletter
	result := repo.DB.Find(&objs)
	if result.Error != nil {
		return nil, result.Error
	}
	return objs, nil
}

func UpdateNewsletterOperation(repo *internal.Repository, obj *models.Newsletter, id int) error {
	result := repo.DB.Where("id = ?", id).Updates(obj)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func DeleteNewsletterOperation(repo *internal.Repository, id int) error {
	result := repo.DB.Where("id = ?", id).Delete(&models.Newsletter{})
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func SubscribeToNewsletterOperation(repo *internal.Repository, user *models.NewsletterUser) error {
	result := repo.DB.Create(user)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func UnsubscribeFromNewsletterOperation(repo *internal.Repository, email string, newsletterID uint) error {
	var user models.NewsletterUser

	result := repo.DB.Where("newsletter_id = ? AND email = ?", newsletterID, email).First(&user)
	if result.Error == gorm.ErrRecordNotFound {
		return nil // ignore if the email is not subscribed to the newsletter
	} else if result.Error != nil {
		return result.Error
	}

	result = repo.DB.Delete(user)
	if result.Error != nil {
		return result.Error
	}

	return nil
}
