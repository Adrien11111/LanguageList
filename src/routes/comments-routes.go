package routes

import (
	"github.com/gin-gonic/gin"
	"go-api/src/controllers"
	"go-api/src/middlewares"
)

// @BasePath /api/v1

// Add comment
// @Summary Add a new comment
// @Description Add a new comment to the list
// @Tags comments
// @Security ApiKeyAuth
// @Accept json
// @Produce json
// @Param comment body models.Comment true "Comment object"
// @Success 201 {object} models.Comment
// @Failure 400
// @Failure 500
// @Router /comments [post]
func AddCommentRoute(r *gin.RouterGroup) {
	r.POST("", middlewares.CheckAuth, controllers.AddComment)
}

// Get all comments
// @Summary Get all comments
// @Description Retrieve a list of all comments
// @Tags comments
// @Accept json
// @Produce json
// @Success 200 {array} models.Comment
// @Failure 500
// @Router /comments [get]
func GetAllCommentsRoute(r *gin.RouterGroup) {
	r.GET("", controllers.GetAllComments)
}

// Get comment by ID
// @Summary Get a comment by ID
// @Description Retrieve a comment by its ID
// @Tags comments
// @Accept json
// @Produce json
// @Param id path int true "Comment ID"
// @Success 200 {object} models.Comment
// @Failure 404
// @Failure 500
// @Router /comments/{id} [get]
func GetCommentByIdRoute(r *gin.RouterGroup) {
	r.GET("/:id", controllers.GetCommentById)
}

// Delete comment by ID
// @Summary Delete a comment by ID
// @Description Delete a comment by its ID
// @Tags comments
// @Security ApiKeyAuth
// @Accept json
// @Produce json
// @Param id path int true "Comment ID"
// @Success 204
// @Failure 404
// @Failure 500
// @Router /comments/{id} [delete]
func DeleteCommentRoute(r *gin.RouterGroup) {
	r.DELETE("/:id", middlewares.CheckAuth, controllers.DeleteComment)
}

func SetupCommentRoutes(r *gin.RouterGroup) {
	AddCommentRoute(r)
	GetAllCommentsRoute(r)
	GetCommentByIdRoute(r)
	DeleteCommentRoute(r)
}
