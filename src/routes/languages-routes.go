package routes

import (
	"github.com/gin-gonic/gin"
	"go-api/src/controllers"
)

func SetupLanguageRoutes(r *gin.RouterGroup) {
	r.POST("", controllers.AddLanguage)
	r.GET("", controllers.GetAllLanguages)
	r.GET("/:id", controllers.GetLanguageById)
}
