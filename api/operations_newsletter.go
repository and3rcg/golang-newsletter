package api

import (
	"errors"
	"newsletter-go/internal"
	"newsletter-go/models"
	"slices"
)

func CreateNewsletterOperation(repo *internal.Repository, obj *models.Newsletter) error {
	obj.EmailList = []string{} // initializing empty list
	result := repo.DB.Create(obj)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func GetNewsletterByIDOperation(repo *internal.Repository, id int) (*models.Newsletter, error) {
	var obj models.Newsletter
	result := repo.DB.Where("id = ?", id).First(&obj)
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

func SubscribeToNewsletterOperation(repo *internal.Repository, email string, id int) error {
	var obj models.Newsletter
	result := repo.DB.Where("id = ?", id).First(&obj)

	if result.Error != nil {
		return result.Error
	}

	// checking if the e-mail is already subscribed to the newsletter
	dupeIdx := slices.Index(obj.EmailList, email)
	if dupeIdx != -1 {
		return errors.New("e-mail address already subscribed to newsletter")
	}

	obj.EmailList = append(obj.EmailList, email)

	result = repo.DB.Save(&obj)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func UnsubscribeFromNewsletterOperation(repo *internal.Repository, email string, id int) error {
	var obj models.Newsletter
	result := repo.DB.Where("id = ?", id).First(&obj)

	if result.Error != nil {
		return result.Error
	}

	// checking if the e-mail is actually subscribed to the newsletter
	dupeIdx := slices.Index(obj.EmailList, email)
	if dupeIdx == -1 {
		return nil
	}

	obj.EmailList = append(obj.EmailList[:dupeIdx], obj.EmailList[dupeIdx+1:]...)

	result = repo.DB.Save(&obj)
	if result.Error != nil {
		return result.Error
	}

	return nil
}
