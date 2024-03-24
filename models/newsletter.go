package models

import "gorm.io/gorm"

type Newsletter struct {
	gorm.Model

	Name        string   `json:"name"`
	Description string   `json:"description"`
	EmailList   []string `json:"email_list" gorm:"type:text[];default:[]"`
}
