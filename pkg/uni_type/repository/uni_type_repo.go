package repository

import (
	"uni_app/database"
	"uni_app/models"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type UniTypeRepository interface {
	Create(uniType *models.UniType) error
	GetByID(ctx echo.Context, ID database.PID, useCache bool) (*models.UniType, error)
	Update(uniType *models.UniType) error
	Delete(ID database.PID) error
	GetAll() ([]models.UniType, error)
}

type uniTypeRepository struct {
	db *gorm.DB
}

func NewUniTypeRepository(db *gorm.DB) UniTypeRepository {
	return &uniTypeRepository{db}
}

func (r *uniTypeRepository) Create(uniType *models.UniType) error {
	return r.db.Create(uniType).Error
}

func (r *uniTypeRepository) GetByID(ctx echo.Context, ID database.PID, useCache bool) (*models.UniType, error) {
	var uniType models.UniType
	if err := r.db.First(&uniType, ID).Error; err != nil {
		return nil, err
	}
	return &uniType, nil
}

func (r *uniTypeRepository) Update(uniType *models.UniType) error {
	return r.db.Save(uniType).Error
}

func (r *uniTypeRepository) Delete(ID database.PID) error {
	return r.db.Delete(&models.UniType{}, ID).Error
}

func (r *uniTypeRepository) GetAll() ([]models.UniType, error) {
	var uniTypes []models.UniType
	if err := r.db.Find(&uniTypes).Error; err != nil {
		return nil, err
	}
	return uniTypes, nil
} 