package models

// RegisterRequest represents a user registration request.
// @Description User registration request model
type RegisterRequest struct {
    Email    string `json:"email" binding:"required,email" example:"adrienpani@gmail.com"`
    Password string `json:"password" binding:"required,min=8" example:"password123"`
}