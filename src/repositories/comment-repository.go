package repositories

import (
	"go-api/src/models"
	"gorm.io/gorm"
)

type CommentRepository struct {
	DB *gorm.DB
}

func NewCommentRepository(db *gorm.DB) *CommentRepository {
	return &CommentRepository{DB: db}
}

func (repo *CommentRepository) GetAllComments() ([]models.Comment, error) {
	var comments []models.Comment
	if err := repo.DB.Find(&comments).Error; err != nil {
		return nil, err
	}
	return comments, nil
}

func (repo *CommentRepository) GetCommentByID(id uint) (*models.Comment, error) {
	var comment models.Comment
	if err := repo.DB.First(&comment, id).Error; err != nil {
		return nil, err
	}
	return &comment, nil
}

func (repo *CommentRepository) CreateComment(comment *models.Comment) error {
	if err := repo.DB.Create(comment).Error; err != nil {
		return err
	}
	return nil
}

func (repo *CommentRepository) DeleteComment(id uint) error {
	if err := repo.DB.Delete(&models.Comment{}, id).Error; err != nil {
		return err
	}
	return nil
}