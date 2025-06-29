package controllers

import (
	"strconv"
	"gorm.io/gorm"
	"github.com/gin-gonic/gin"
	"go-api/src/models"
	"go-api/src/repositories"
)

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
	comment, err := commentRepository.GetCommentByID(uint(commentID))

	if err != nil {
		c.JSON(404, gin.H{"error": "Comment not found"})
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

	if err := commentRepository.DeleteComment(uint(commentID)); err != nil {
		c.JSON(404, gin.H{"error": "Comment not found"})
		return
	}

	c.Status(204)
}