package repository

import (
	"uni_app/database"
	"uni_app/models"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type RequestRepository interface {
	Create(request *models.FetchRequest) error
	GetByID(ctx echo.Context, ID database.PID, useCache bool) (*models.FetchRequest, error)
	Update(request *models.FetchRequest) error
	Delete(ID database.PID) error
	GetAll() ([]models.FetchRequest, error)
}

type requestRepository struct {
	db *gorm.DB
}

func NewRequestRepository(db *gorm.DB) RequestRepository {
	return &requestRepository{db}
}

func (r *requestRepository) Create(request *models.FetchRequest) error {
	return r.db.Create(request).Error
}

func (r *requestRepository) GetByID(ctx echo.Context, ID database.PID, useCache bool) (*models.FetchRequest, error) {
	var request models.FetchRequest
	if err := r.db.First(&request, ID).Error; err != nil {
		return nil, err
	}
	return &request, nil
}

func (r *requestRepository) Update(request *models.FetchRequest) error {
	return r.db.Save(request).Error
}

func (r *requestRepository) Delete(ID database.PID) error {
	return r.db.Delete(&models.FetchRequest{}, ID).Error
}

func (r *requestRepository) GetAll() ([]models.FetchRequest, error) {
	var requests []models.FetchRequest
	if err := r.db.Find(&requests).Error; err != nil {
		return nil, err
	}
	return requests, nil
} 