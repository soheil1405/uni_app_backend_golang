package usecase

import (
	"uni_app/database"
	"uni_app/models"
	repository "uni_app/pkg/user_role/repository"

	"github.com/labstack/echo/v4"
)

type UserRoleUsecase interface {
	CreateUserRole(userRole *models.UserRole) error
	GetUserRoleByID(ctx echo.Context, ID database.PID, useCache bool) (*models.UserRole, error)
	UpdateUserRole(userRole *models.UserRole) error
	DeleteUserRole(ID database.PID) error
	GetAllUserRoles() ([]models.UserRole, error)
}

type userRoleUsecase struct {
	repo repository.UserRoleRepository
}

func NewUserRoleUsecase(repo repository.UserRoleRepository) UserRoleUsecase {
	return &userRoleUsecase{repo}
}

func (u *userRoleUsecase) CreateUserRole(userRole *models.UserRole) error {
	return u.repo.Create(userRole)
}

func (u *userRoleUsecase) GetUserRoleByID(ctx echo.Context, ID database.PID, useCache bool) (*models.UserRole, error) {
	return u.repo.GetByID(ctx, ID, useCache)
}

func (u *userRoleUsecase) UpdateUserRole(userRole *models.UserRole) error {
	return u.repo.Update(userRole)
}

func (u *userRoleUsecase) DeleteUserRole(ID database.PID) error {
	return u.repo.Delete(ID)
}

func (u *userRoleUsecase) GetAllUserRoles() ([]models.UserRole, error) {
	return u.repo.GetAll()
} 