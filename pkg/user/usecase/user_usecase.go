package usecases

import (
	"uni_app/models"
	repositories "uni_app/pkg/user/repository"
)

type UserUsecase interface {
	CreateUser(user *models.User) error
	GetUserByID(id uint) (*models.User, error)
	UpdateUser(user *models.User) error
	DeleteUser(id uint) error
	GetAllUsers() ([]models.User, error)
}

type userUsecase struct {
	repo repositories.UserRepository
}

func NewUserUsecase(repo repositories.UserRepository) UserUsecase {
	return &userUsecase{repo}
}

func (u *userUsecase) CreateUser(user *models.User) error {
	return u.repo.Create(user)
}

func (u *userUsecase) GetUserByID(id uint) (*models.User, error) {
	return u.repo.GetByID(id)
}

func (u *userUsecase) UpdateUser(user *models.User) error {
	return u.repo.Update(user)
}

func (u *userUsecase) DeleteUser(id uint) error {
	return u.repo.Delete(id)
}

func (u *userUsecase) GetAllUsers() ([]models.User, error) {
	return u.repo.GetAll()
}
