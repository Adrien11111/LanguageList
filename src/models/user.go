package models

import "gorm.io/gorm"

// User model for storing user information in the database.
// @Description User model for storing user information in the database.
type User struct {
	gorm.Model `swaggerignore:"true"`
	Email string `json:"email" gorm:"not null" example:"adrienpani@gmail.com"`
	Password string `swaggerignore:"true" gorm:"not null" example:"password123"`
}

