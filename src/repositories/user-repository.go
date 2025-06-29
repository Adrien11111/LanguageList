package repositories

import (
	"go-api/src/models"
	"gorm.io/gorm"
)

type UserRepository struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{DB: db}
}

func (repo *UserRepository) GetAllUsers() ([]models.User, error) {
	var users []models.User
	if err := repo.DB.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func (repo *UserRepository) GetUserByID(id uint) (*models.User, error) {
	var user models.User
	if err := repo.DB.First(&user, id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (repo *UserRepository) GetUserByEmail(email string) (*models.User, error) {
	var user models.User
    if err := repo.DB.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (repo *UserRepository) CreateUser(user *models.User) error {
	if err := repo.DB.Create(user).Error; err != nil {
		return err
	}
	return nil
}

func (repo *UserRepository) DeleteUser(id uint) error {
	if err := repo.DB.Delete(&models.User{}, id).Error; err != nil {
		return err
	}
	return nil
}