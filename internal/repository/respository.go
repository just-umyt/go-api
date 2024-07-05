package repository

import (
	"github.com/just-umyt/go-api/internal/models"
	"gorm.io/gorm"
)

type UserRepository struct {
	Database *gorm.DB
}

func NewUserRepo(db *gorm.DB) *UserRepository {
	return &UserRepository{Database: db}
}

func (repo *UserRepository) GetAll() (*[]models.User, error) {
	var users []models.User

	if err := repo.Database.Find(&users).Error; err != nil {
		return nil, err
	}

	return &users, nil
}

func (repo *UserRepository) GetById(id int) (*models.User, error) {
	var user models.User

	if err := repo.Database.First(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

func (repo *UserRepository) Create(user *models.User) error {
	if err := repo.Database.Create(&user).Error; err != nil {
		return nil
	}

	return nil
}

func (repo *UserRepository) UpdateUser(id int, newUser *models.User) error {

	var oldUser models.User

	if err := repo.Database.First(&oldUser, id).Error; err != nil {
		return err
	}

	oldUser.Name = newUser.Name
	oldUser.Email = newUser.Email

	repo.Database.Save(&oldUser)

	return nil
}

func (repo *UserRepository) DeleteUser(id int) error {

	repo.Database.Delete(&models.User{}, id)

	return nil
}
