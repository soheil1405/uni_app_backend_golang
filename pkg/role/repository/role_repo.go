package repositories

import (
	"uni_app/database"
	"uni_app/models"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type RoleRepository interface {
	Create(role *models.Role) error
	GetByID(ctx echo.Context, ID database.PID, useCache bool) (*models.Role, error)
	Update(role *models.Role) error
	Delete(ID database.PID) error
	GetAll() ([]models.Role, error)
}

type roleRepository struct {
	db *gorm.DB
}

func NewRoleRepository(db *gorm.DB) RoleRepository {
	return &roleRepository{db}
}

func (r *roleRepository) Create(role *models.Role) error {
	return r.db.Create(role).Error
}

func (r *roleRepository) GetByID(ctx echo.Context, ID database.PID, useCache bool) (*models.Role, error) {
	var role models.Role
	if err := r.db.First(&role, ID).Error; err != nil {
		return nil, err
	}
	return &role, nil
}

func (r *roleRepository) Update(role *models.Role) error {
	return r.db.Save(role).Error
}

func (r *roleRepository) Delete(ID database.PID) error {
	return r.db.Delete(&models.Role{}, ID).Error
}

func (r *roleRepository) GetAll() ([]models.Role, error) {
	var roles []models.Role
	if err := r.db.Find(&roles).Error; err != nil {
		return nil, err
	}
	return roles, nil
}
