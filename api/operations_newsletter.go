package api

import (
	"newsletter-go/internal"
	"newsletter-go/models"
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
