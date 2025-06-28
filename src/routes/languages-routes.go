package routes

import (
	"github.com/gin-gonic/gin"
	"go-api/src/controllers"
)

// @BasePath /api/v1

// Add language
// @Summary Add a new programming language
// @Description Add a new programming language to the list
// @Tags languages
// @Accept json
// @Produce json
// @Param language body models.Language true "Language object"
// @Success 201 {object} models.Language
// @Failure 400
// @Failure 500
// @Router /languages [post]
func AddLanguageRoute(r *gin.RouterGroup) {
	r.POST("", controllers.AddLanguage)
}

// Get all languages
// @Summary Get all programming languages
// @Description Retrieve a list of all programming languages
// @Tags languages
// @Accept json
// @Produce json
// @Success 200 {array} models.Language
// @Failure 500
// @Router /languages [get]
func GetAllLanguagesRoute(r *gin.RouterGroup) {
	r.GET("", controllers.GetAllLanguages)
}

// Get language by ID
// @Summary Get a programming language by ID
// @Description Retrieve a programming language by its ID
// @Tags languages
// @Accept json
// @Produce json
// @Param id path int true "Language ID"
// @Success 200 {object} models.Language
// @Failure 404
// @Failure 500
// @Router /languages/{id} [get]
func GetLanguageByIdRoute(r *gin.RouterGroup) {
	r.GET("/:id", controllers.GetLanguageById)
}
func SetupLanguageRoutes(r *gin.RouterGroup) {
	AddLanguageRoute(r)
	GetAllLanguagesRoute(r)
	GetLanguageByIdRoute(r)
}
