package models

import "gorm.io/gorm"

// Language represents a programming language.
// @Description Language model for storing programming languages in the database.
type Language struct {
	gorm.Model `swaggerignore:"true"`
	Name        string `json:"name" gorm:"unique;not null" example:"Go"`
	Description string `json:"description" gorm:"not null" example:"A statically typed, compiled programming language designed for simplicity and efficiency."`
}

