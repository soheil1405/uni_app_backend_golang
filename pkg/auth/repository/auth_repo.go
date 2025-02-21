package repositories

import (
	"uni_app/database"
	"uni_app/models"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type AuthRepository interface {
	Enforce(ctx echo.Context, auth models.AuthRules, useCache bool) (*models.AuthRules, error)
	Create(auth *models.AuthRules) error
	GetByID(ctx echo.Context, ID database.PID, useCache bool) (*models.AuthRules, error)
	Update(auth *models.AuthRules) error
	Delete(ID database.PID) error
	GetAll() ([]models.AuthRules, error)
}

type authRepository struct {
	db *gorm.DB
}

func NewAuthRepository(db *gorm.DB) AuthRepository {
	return &authRepository{db}
}

func (r *authRepository) Create(auth *models.AuthRules) error {
	return r.db.Create(auth).Error
}

func (r *authRepository) Enforce(ctx echo.Context, auth models.AuthRules, useCache bool) (*models.AuthRules, error) {
	if err := r.db.Find(&auth).Error; err != nil {
		return nil, err
	}
	return &auth, nil
}

func (r *authRepository) GetByID(ctx echo.Context, ID database.PID, useCache bool) (*models.AuthRules, error) {
	var auth models.AuthRules
	if err := r.db.First(&auth, ID).Error; err != nil {
		return nil, err
	}
	return &auth, nil
}

func (r *authRepository) Update(auth *models.AuthRules) error {
	return r.db.Save(auth).Error
}

func (r *authRepository) Delete(ID database.PID) error {
	return r.db.Delete(&models.AuthRules{}, ID).Error
}

func (r *authRepository) GetAll() ([]models.AuthRules, error) {
	var auths []models.AuthRules
	if err := r.db.Find(&auths).Error; err != nil {
		return nil, err
	}
	return auths, nil
}
