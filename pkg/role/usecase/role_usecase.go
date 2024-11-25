package usecases

import (
	"uni_app/models"
	repositories "uni_app/pkg/role/repository"
)

type RoleUsecase interface {
	CreateRole(role *models.Role) error
	GetRoleByID(id uint) (*models.Role, error)
	UpdateRole(role *models.Role) error
	DeleteRole(id uint) error
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

func (u *roleUsecase) GetRoleByID(id uint) (*models.Role, error) {
	return u.repo.GetByID(id)
}

func (u *roleUsecase) UpdateRole(role *models.Role) error {
	return u.repo.Update(role)
}

func (u *roleUsecase) DeleteRole(id uint) error {
	return u.repo.Delete(id)
}

func (u *roleUsecase) GetAllRoles() ([]models.Role, error) {
	return u.repo.GetAll()
}
