package routes

import (
	"github.com/gin-gonic/gin"
	"go-api/src/controllers"
	"go-api/src/middlewares"
)

// @BasePath /api/v1
// LoginRoute handles user login
// @Summary User login
// @Description Authenticate user and return token
// @Tags auth
// @Accept json
// @Produce json
// @Param user body models.RegisterRequest true "User object with email and password"
// @Success 200
// @Failure 400 
// @Failure 401 
// @Router /auth/login [post]
func LoginRoute(r *gin.RouterGroup) {
	r.POST("/login", controllers.Login)
}

// RegisterRoute handles user registration
// @Summary User registration
// @Description Register a new user and return success message
// @Tags auth
// @Accept json
// @Produce json
// @Param user body models.RegisterRequest true "User object with email and password"
// @Success 201 {object} models.User
// @Failure 400 {object} string "Invalid input"
// @Failure 409 {object} string "User already exists"
// @Router /auth/register [post]
func RegisterRoute(r *gin.RouterGroup) {
	r.POST("/register", controllers.Register)
}

// DeleteAccountRoute handles account deletion
// @Summary Delete user account
// @Description Delete the authenticated user's account
// @Tags auth
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Success 204
// @Failure 401
// @Failure 404
// @Router /auth/account [delete]
func DeleteAccountRoute(r *gin.RouterGroup) {
	r.DELETE("/account", middlewares.CheckAuth, controllers.DeleteAccount)
}

func SetupAuthRoutes(r *gin.RouterGroup) {
	// Define the authentication routes
	LoginRoute(r)
	RegisterRoute(r)
	DeleteAccountRoute(r)
}
