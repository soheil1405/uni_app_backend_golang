package repository

import (
	"uni_app/database"
	"uni_app/models"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type DegreeRepository interface {
	Create(degree *models.DegreeLevel) error
	GetByID(ctx echo.Context, ID database.PID, useCache bool) (*models.DegreeLevel, error)
	Update(degree *models.DegreeLevel) error
	Delete(ID database.PID) error
	GetAll() ([]models.DegreeLevel, error)
}

type degreeRepository struct {
	db *gorm.DB
}

func NewDegreeRepository(db *gorm.DB) DegreeRepository {
	return &degreeRepository{db}
}

func (r *degreeRepository) Create(degree *models.DegreeLevel) error {
	return r.db.Create(degree).Error
}

func (r *degreeRepository) GetByID(ctx echo.Context, ID database.PID, useCache bool) (*models.DegreeLevel, error) {
	var degree models.DegreeLevel
	if err := r.db.First(&degree, ID).Error; err != nil {
		return nil, err
	}
	return &degree, nil
}

func (r *degreeRepository) Update(degree *models.DegreeLevel) error {
	return r.db.Save(degree).Error
}

func (r *degreeRepository) Delete(ID database.PID) error {
	return r.db.Delete(&models.DegreeLevel{}, ID).Error
}

func (r *degreeRepository) GetAll() ([]models.DegreeLevel, error) {
	var degrees []models.DegreeLevel
	if err := r.db.Find(&degrees).Error; err != nil {
		return nil, err
	}
	return degrees, nil
} 