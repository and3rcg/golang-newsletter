package models

import (
	"database/sql/driver"
	"errors"
	"strings"

	"gorm.io/gorm"
)

type EmailList []string

func (e EmailList) Value() (driver.Value, error) {
	if len(e) == 0 {
		return nil, nil
	}

	return strings.Join(e, ","), nil
}

// this function controls the way the system handles an array of strings for manipulation, turning it into a slice of strings
// SQLite saves an array of strings (gorm:"type:text[]" tag) as values separated by a comma, i.e.: val1,val2,val3,etc
func (e *EmailList) Scan(src any) error {
	if src == nil {
		return nil
	}

	strVal, ok := src.(string)
	if !ok {
		return errors.New("failed to type assert to string")
	}

	*e = strings.Split(strVal, ",")
	return nil
}

type Newsletter struct {
	gorm.Model

	Name        string    `json:"name" validate:"required"`
	Description string    `json:"description"`
	EmailList   EmailList `json:"email_list" gorm:"type:text[]"`
}
