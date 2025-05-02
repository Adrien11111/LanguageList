package models

import "gorm.io/gorm"

// Language represents a programming language.
type Language struct {
	gorm.Model
	Name        string `json:"name" gorm:"unique;not null"`
	Description string `json:"description" gorm:"not null"`
}

