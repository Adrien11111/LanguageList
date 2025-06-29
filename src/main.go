package main

import (
    "log"
	"go-api/src/database"
	"go-api/src/middlewares"
	"go-api/src/routes"
    "go-api/src/models"
	"github.com/gin-gonic/gin"
    docs "go-api/src/docs"
    swaggerfiles "github.com/swaggo/files"
    ginSwagger "github.com/swaggo/gin-swagger"

)

// @title Language List API
// @version 1.0
// @description API for managing a list of programming languages.

// @contact.name Adrien Panis
// @contact.email adrienpani@gmail.com

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

// @externalDocs.description  OpenAPI
// @externalDocs.url          https://swagger.io/resources/open-api/
func main() {
    // Connect to database
    if err := database.Connect(); err != nil {
        log.Fatal("Failed to connect to database:", err)
    }

    // Auto-migrate the database schema
    if err := database.DB.AutoMigrate(&models.Language{}, &models.User{}, &models.Comment{}); err != nil {
        log.Fatal("Failed to migrate database:", err)
    }

    r := gin.Default()
    docs.SwaggerInfo.BasePath = "/api/v1"

	// Use middlewares
    r.Use(middlewares.DatabaseMiddleware(database.DB))

	v1 := r.Group("/api/v1")
	routes.SetupLanguageRoutes(v1.Group("/languages"))
	routes.SetupCommentRoutes(v1.Group("/comments"))
	routes.SetupAuthRoutes(v1.Group("/auth"))

    r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	r.Run(":8080")
}
