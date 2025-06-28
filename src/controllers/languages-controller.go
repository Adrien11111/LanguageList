package controllers

import (
	"strconv"
	"gorm.io/gorm"
	"github.com/gin-gonic/gin"
	"go-api/src/models"
	"go-api/src/repositories"
)

func GetAllLanguages(c *gin.Context) {
	db := c.MustGet("DB").(*gorm.DB)

	languageRepository := repositories.NewLanguageRepository(db)
	languages, err := languageRepository.GetAllLanguages()

	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to fetch languages"})
		return
	}

	c.JSON(200, languages)
}

func GetLanguageById(c *gin.Context) {
	id := c.Param("id")
	languageID, err := strconv.ParseUint(id, 10, 32)

	if err != nil || languageID <= 0 {
		c.JSON(400, gin.H{"error": "Invalid ID"})
		return
	}

	db := c.MustGet("DB").(*gorm.DB)
	languageRepository := repositories.NewLanguageRepository(db)
	language, err := languageRepository.GetLanguageByID(uint(languageID))

	if err != nil {
		c.JSON(404, gin.H{"error": "Language not found"})
		return
	}

	c.JSON(200, language)
}

func AddLanguage(c *gin.Context) {
	var language models.Language

	if err := c.ShouldBindJSON(&language); err != nil {
		c.JSON(400, gin.H{"error": "Invalid input"})
		return
	}

	db := c.MustGet("DB").(*gorm.DB)
	languageRepository := repositories.NewLanguageRepository(db)

	if err := languageRepository.CreateLanguage(&language); err != nil {
		c.JSON(500, gin.H{"error": "Failed to create language"})
		return
	}

	c.JSON(201, language)
}

func DeleteLanguage(c *gin.Context) {
	id := c.Param("id")
	languageID, err := strconv.ParseUint(id, 10, 32)

	if err != nil || languageID <= 0 {
		c.JSON(400, gin.H{"error": "Invalid ID"})
		return
	}

	db := c.MustGet("DB").(*gorm.DB)
	languageRepository := repositories.NewLanguageRepository(db)

	if err := languageRepository.DeleteLanguage(uint(languageID)); err != nil {
		c.JSON(404, gin.H{"error": "Language not found"})
		return
	}

	c.Status(204) // No Content
}