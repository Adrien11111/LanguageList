package repositories

import (
	"go-api/src/models"
	"gorm.io/gorm"
)

type LanguageRepository struct {
	DB *gorm.DB
}

func NewLanguageRepository(db *gorm.DB) *LanguageRepository {
	return &LanguageRepository{DB: db}
}

func (repo *LanguageRepository) GetAllLanguages() ([]models.Language, error) {
	var languages []models.Language
	if err := repo.DB.Find(&languages).Error; err != nil {
		return nil, err
	}
	return languages, nil
}

func (repo *LanguageRepository) GetLanguageByID(id uint) (*models.Language, error) {
	var language models.Language
	if err := repo.DB.First(&language, id).Error; err != nil {
		return nil, err
	}
	return &language, nil
}

func (repo *LanguageRepository) CreateLanguage(language *models.Language) error {
	if err := repo.DB.Create(language).Error; err != nil {
		return err
	}
	return nil
}