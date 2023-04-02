package model

import (
	"time"

	validation "github.com/go-ozzo/ozzo-validation"
	"gorm.io/gorm"
)

type Book struct {
	ID        int64          `json:"id" gorm:"primaryKey"`
	Title     string         `json:"name_book" `
	Author    string         `json:"author" `
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"-"`
}

func (b Book) Validation() error { // custom validation
	return validation.ValidateStruct(&b,
		validation.Field(&b.Title, validation.Required),
		validation.Field(&b.Author, validation.Required),
	)

}
