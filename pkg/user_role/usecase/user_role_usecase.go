package usecase

import (
	"errors"
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
	GetUserRolesByUserID(userID database.PID) ([]models.UserRole, error)
	GetUserRolesByRoleID(roleID database.PID) ([]models.UserRole, error)
}

type userRoleUsecase struct {
	repo repository.UserRoleRepository
}

func NewUserRoleUsecase(repo repository.UserRoleRepository) UserRoleUsecase {
	return &userRoleUsecase{repo}
}

func (u *userRoleUsecase) CreateUserRole(userRole *models.UserRole) error {
	// Validate required fields
	if userRole.UserID == 0 {
		return errors.New("user ID is required")
	}
	if userRole.RoleID == 0 {
		return errors.New("role ID is required")
	}

	// Check if the combination already exists
	existingRoles, err := u.repo.GetAll()
	if err != nil {
		return err
	}

	for _, existing := range existingRoles {
		if existing.UserID == userRole.UserID && existing.RoleID == userRole.RoleID {
			return errors.New("this role is already assigned to this user")
		}
	}

	return u.repo.Create(userRole)
}

func (u *userRoleUsecase) GetUserRoleByID(ctx echo.Context, ID database.PID, useCache bool) (*models.UserRole, error) {
	return u.repo.GetByID(ctx, ID, useCache)
}

func (u *userRoleUsecase) UpdateUserRole(userRole *models.UserRole) error {
	// Validate required fields
	if userRole.UserID == 0 {
		return errors.New("user ID is required")
	}
	if userRole.RoleID == 0 {
		return errors.New("role ID is required")
	}

	// Check if the combination already exists (excluding current record)
	existingRoles, err := u.repo.GetAll()
	if err != nil {
		return err
	}

	for _, existing := range existingRoles {
		if existing.ID != userRole.ID &&
			existing.UserID == userRole.UserID &&
			existing.RoleID == userRole.RoleID {
			return errors.New("this role is already assigned to this user")
		}
	}

	return u.repo.Update(userRole)
}

func (u *userRoleUsecase) DeleteUserRole(ID database.PID) error {
	// Check if there are any related records before deletion
	// This could be expanded based on your requirements
	return u.repo.Delete(ID)
}

func (u *userRoleUsecase) GetAllUserRoles() ([]models.UserRole, error) {
	return u.repo.GetAll()
}

func (u *userRoleUsecase) GetUserRolesByUserID(userID database.PID) ([]models.UserRole, error) {
	allRoles, err := u.repo.GetAll()
	if err != nil {
		return nil, err
	}

	var filteredRoles []models.UserRole
	for _, role := range allRoles {
		if role.UserID == userID {
			filteredRoles = append(filteredRoles, role)
		}
	}

	return filteredRoles, nil
}

func (u *userRoleUsecase) GetUserRolesByRoleID(roleID database.PID) ([]models.UserRole, error) {
	allRoles, err := u.repo.GetAll()
	if err != nil {
		return nil, err
	}

	var filteredRoles []models.UserRole
	for _, role := range allRoles {
		if role.RoleID == roleID {
			filteredRoles = append(filteredRoles, role)
		}
	}

	return filteredRoles, nil
} 