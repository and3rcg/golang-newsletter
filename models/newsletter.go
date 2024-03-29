package models

type Newsletter struct {
	ID          uint             `gorm:"primarykey" json:"id"`
	Name        string           `json:"name" validate:"required"`
	Description string           `json:"description"`
	Users       []NewsletterUser `gorm:"foreignKey:NewsletterID"`
}

type NewsletterUser struct {
	ID           uint   `gorm:"primarykey" json:"id"`
	Name         string `json:"name" validate:"required"`
	Email        string `gorm:"index:email_newsletteruser_index,unique" json:"email" validate:"required,valid_email"`
	NewsletterID uint   `gorm:"index:email_newsletteruser_index,unique" json:"newsletter_id" validate:"required"`
}
