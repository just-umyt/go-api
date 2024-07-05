package usecase

import "github.com/just-umyt/go-api/internal/models"

type UserRepo interface {
	GetAll() (*[]models.User, error)
	GetById(id int) (*models.User, error)
	Create(user *models.User) error
	UpdateUser(id int, newUser *models.User) error
	DeleteUser(id int) error
}

type UserUseCaseImpl struct {
	URepo UserRepo
}

func NewUserUsecase(ur UserRepo) *UserUseCaseImpl {
	return &UserUseCaseImpl{URepo: ur}
}

func (ur *UserUseCaseImpl) GetUsers() (*[]models.User, error) {
	return ur.URepo.GetAll()
}

func (ur *UserUseCaseImpl) GetUserById(id int) (*models.User, error) {
	return ur.URepo.GetById(id)
}

func (ur *UserUseCaseImpl) CreateUser(user *models.User) error {
	return ur.URepo.Create(user)
}

func (ur *UserUseCaseImpl) UpdateUser(id int, user *models.User) error {
	return ur.URepo.UpdateUser(id, user)
}

func (ur *UserUseCaseImpl) DeleteUser(id int) error {
	return ur.URepo.DeleteUser(id)
}
