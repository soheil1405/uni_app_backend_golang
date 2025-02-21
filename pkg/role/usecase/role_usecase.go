package usecases

import (
	"uni_app/database"
	"uni_app/models"
	repositories "uni_app/pkg/role/repository"

	"github.com/labstack/echo/v4"
)

type RoleUsecase interface {
	CreateRole(role *models.Role) error
	GetRoleByID(ctx echo.Context, ID database.PID, useCache bool) (*models.Role, error)
	UpdateRole(role *models.Role) error
	DeleteRole(ID database.PID) error
	GetAllRoles() ([]models.Role, error)
}

type roleUsecase struct {
	repo repositories.RoleRepository
}

func NewRoleUsecase(repo repositories.RoleRepository) RoleUsecase {
	return &roleUsecase{repo}
}

func (u *roleUsecase) CreateRole(role *models.Role) error {
	return u.repo.Create(role)
}

func (u *roleUsecase) GetRoleByID(ctx echo.Context, ID database.PID, useCache bool) (*models.Role, error) {
	return u.repo.GetByID(ctx, ID, useCache)
}

func (u *roleUsecase) UpdateRole(role *models.Role) error {
	return u.repo.Update(role)
}

func (u *roleUsecase) DeleteRole(ID database.PID) error {
	return u.repo.Delete(ID)
}

func (u *roleUsecase) GetAllRoles() ([]models.Role, error) {
	return u.repo.GetAll()
}
