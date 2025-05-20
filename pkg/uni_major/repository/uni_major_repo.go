package repository

import (
	"uni_app/database"
	"uni_app/models"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type UniMajorRepository interface {
	Create(uniMajor *models.UniMajor) error
	GetByID(ctx echo.Context, ID database.PID, useCache bool) (*models.UniMajor, error)
	Update(uniMajor *models.UniMajor) error
	Delete(ID database.PID) error
	GetAll() ([]models.UniMajor, error)
}

type uniMajorRepository struct {
	db *gorm.DB
}

func NewUniMajorRepository(db *gorm.DB) UniMajorRepository {
	return &uniMajorRepository{db}
}

func (r *uniMajorRepository) Create(uniMajor *models.UniMajor) error {
	return r.db.Create(uniMajor).Error
}

func (r *uniMajorRepository) GetByID(ctx echo.Context, ID database.PID, useCache bool) (*models.UniMajor, error) {
	var uniMajor models.UniMajor
	if err := r.db.First(&uniMajor, ID).Error; err != nil {
		return nil, err
	}
	return &uniMajor, nil
}

func (r *uniMajorRepository) Update(uniMajor *models.UniMajor) error {
	return r.db.Save(uniMajor).Error
}

func (r *uniMajorRepository) Delete(ID database.PID) error {
	return r.db.Delete(&models.UniMajor{}, ID).Error
}

func (r *uniMajorRepository) GetAll() ([]models.UniMajor, error) {
	var uniMajors []models.UniMajor
	if err := r.db.Find(&uniMajors).Error; err != nil {
		return nil, err
	}
	return uniMajors, nil
} 