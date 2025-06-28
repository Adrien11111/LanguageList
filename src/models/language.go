package models

import "gorm.io/gorm"

// Language represents a programming language.
// @Description Language model for storing programming languages in the database.
// @ID Language
// @Model
// @TableName languages
// @Example {"id":1, "name":"Go", "description":"A statically typed, compiled programming language designed for simplicity and efficiency."}
type Language struct {
	// @swaggerignore
	gorm.Model `swaggerignore:"true"`
	Name        string `json:"name" gorm:"unique;not null"`
	Description string `json:"description" gorm:"not null"`
}

