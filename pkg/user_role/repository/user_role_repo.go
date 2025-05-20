package repository

import (
	"uni_app/database"
	"uni_app/models"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type UserRoleRepository interface {
	Create(userRole *models.UserRole) error
	GetByID(ctx echo.Context, ID database.PID, useCache bool) (*models.UserRole, error)
	Update(userRole *models.UserRole) error
	Delete(ID database.PID) error
	GetAll() ([]models.UserRole, error)
}

type userRoleRepository struct {
	db *gorm.DB
}

func NewUserRoleRepository(db *gorm.DB) UserRoleRepository {
	return &userRoleRepository{db}
}

func (r *userRoleRepository) Create(userRole *models.UserRole) error {
	return r.db.Create(userRole).Error
}

func (r *userRoleRepository) GetByID(ctx echo.Context, ID database.PID, useCache bool) (*models.UserRole, error) {
	var userRole models.UserRole
	if err := r.db.First(&userRole, ID).Error; err != nil {
		return nil, err
	}
	return &userRole, nil
}

func (r *userRoleRepository) Update(userRole *models.UserRole) error {
	return r.db.Save(userRole).Error
}

func (r *userRoleRepository) Delete(ID database.PID) error {
	return r.db.Delete(&models.UserRole{}, ID).Error
}

func (r *userRoleRepository) GetAll() ([]models.UserRole, error) {
	var userRoles []models.UserRole
	if err := r.db.Find(&userRoles).Error; err != nil {
		return nil, err
	}
	return userRoles, nil
} 