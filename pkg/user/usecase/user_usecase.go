package usecases

import (
	"uni_app/database"
	"uni_app/models"
	repositories "uni_app/pkg/user/repository"
)

type UserUsecase interface {
	CreateUser(user *models.User) error
	GetUserByID(ID database.PID) (*models.User, error)
	UpdateUser(user *models.User) error
	DeleteUser(ID database.PID) error
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

func (u *userUsecase) GetUserByID(ID database.PID) (*models.User, error) {
	return u.repo.GetByID(ID)
}

func (u *userUsecase) UpdateUser(user *models.User) error {
	return u.repo.Update(user)
}

func (u *userUsecase) DeleteUser(ID database.PID) error {
	return u.repo.Delete(ID)
}

func (u *userUsecase) GetAllUsers() ([]models.User, error) {
	return u.repo.GetAll()
}
