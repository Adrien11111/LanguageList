package controllers

import (
	"strconv"
	"gorm.io/gorm"
	"github.com/gin-gonic/gin"
	"go-api/src/models"
	"go-api/src/repositories"
)

const commentNotFoundMsg = "Comment not found"

func GetAllComments(c *gin.Context) {
	db := c.MustGet("DB").(*gorm.DB)

	commentRepository := repositories.NewCommentRepository(db)
	comments, err := commentRepository.GetAllComments()

	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to fetch comments"})
		return
	}

	c.JSON(200, comments)
}

func GetCommentById(c *gin.Context) {
	id := c.Param("id")
	commentID, err := strconv.ParseUint(id, 10, 32)

	if err != nil || commentID <= 0 {
		c.JSON(400, gin.H{"error": "Invalid ID"})
		return
	}

	db := c.MustGet("DB").(*gorm.DB)
	commentRepository := repositories.NewCommentRepository(db)
	if err != nil {
		c.JSON(404, gin.H{"error": commentNotFoundMsg})
		return
	}
		return
	}

	c.JSON(200, comment)
}

func AddComment(c *gin.Context) {
	var comment models.Comment

	if err := c.ShouldBindJSON(&comment); err != nil {
		c.JSON(400, gin.H{"error": "Invalid input"})
		return
	}

	db := c.MustGet("DB").(*gorm.DB)
	commentRepository := repositories.NewCommentRepository(db)

	user := c.MustGet("User").(models.User)
	comment.UserID = user.ID
	comment.User = user

	languageID, err := strconv.ParseUint(c.Query("language_id"), 10, 32)
	if err != nil || languageID <= 0 {
		c.JSON(400, gin.H{"error": "Invalid language ID"})
		return
	}
	comment.LanguageID = uint(languageID)
	languageRepository := repositories.NewLanguageRepository(db)
	language, err := languageRepository.GetLanguageByID(comment.LanguageID)
	if err != nil {
		c.JSON(404, gin.H{"error": "Language not found"})
		return
	}
	comment.Language = language

	if err := commentRepository.CreateComment(&comment); err != nil {
		c.JSON(500, gin.H{"error": "Failed to create comment"})
		return
	}

	c.JSON(201, comment)
}

func DeleteComment(c *gin.Context) {
	id := c.Param("id")
	commentID, err := strconv.ParseUint(id, 10, 32)

	if err != nil || commentID <= 0 {
		c.JSON(400, gin.H{"error": "Invalid ID"})
		return
	}

	db := c.MustGet("DB").(*gorm.DB)
	commentRepository := repositories.NewCommentRepository(db)

	user := c.MustGet("User").(models.User)
	comment, err := commentRepository.GetCommentByID(uint(commentID))
	if err != nil {
		c.JSON(404, gin.H{"error": commentNotFoundMsg})
		return
	}
	if comment.UserID != user.ID {
		c.JSON(403, gin.H{"error": "You are not authorized to delete this comment"})
		return
	if err := commentRepository.DeleteComment(uint(commentID)); err != nil {
		c.JSON(404, gin.H{"error": commentNotFoundMsg})
		return
	}
		return
	}

	c.Status(204)
}