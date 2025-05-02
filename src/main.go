package main

import (
	"go-api/src/database"
	"go-api/src/middlewares"
	"go-api/src/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	r:= gin.Default()
    r.Use(middlewares.DatabaseMiddleware(database.DB))

	r.Group("/api/v1")
	routes.SetupLanguageRoutes(r.Group("/languages"))
	// routes.SetupAuthRoutes(r.Group("/auth"))

	r.Run(":8080")
}
