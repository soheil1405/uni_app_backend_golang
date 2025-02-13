package repositories

import (
	"uni_app/database"
	"uni_app/models"

	"gorm.io/gorm"
)

type RoleRepository interface {
	Create(role *models.Role) error
	GetByID(ID database.PID) (*models.Role, error)
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

func (r *roleRepository) GetByID(ID database.PID) (*models.Role, error) {
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
