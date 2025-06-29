package models

import "gorm.io/gorm"

// Comment posted by users about programming languages.
// @Description Comment model for storing user comments about programming languages.
type Comment struct {
	gorm.Model `swaggerignore:"true"`
	Content    string `json:"content" gorm:"not null" example:"I love Go for its simplicity and performance."`
	User       User   `json:"user" gorm:"foreignKey:UserID;references:ID"`
	UserID     uint   `json:"user_id" gorm:"not null"`
	Language   Language `json:"language" gorm:"foreignKey:LanguageID;references:ID"`
	LanguageID uint   `json:"language_id" gorm:"not null"`
}

