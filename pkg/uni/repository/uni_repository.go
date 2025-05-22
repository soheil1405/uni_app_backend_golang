package repository

import (
	"uni_app/database"
	"uni_app/models"
	"uni_app/utils/helpers"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type UniRepository interface {
	Create(uni *models.Uni) error
	GetByID(ctx echo.Context, ID database.PID, useCache bool) (*models.Uni, error)
	Update(uni *models.Uni) error
	Delete(ID database.PID) error
	GetAll(ctx echo.Context, request models.FetchUniRequest) ([]models.Uni, *helpers.PaginateTemplate, error)
}

type uniRepository struct {
	db *gorm.DB
}

func NewUniRepository(db *gorm.DB) UniRepository {
	return &uniRepository{db}
}

func (r *uniRepository) Create(uni *models.Uni) error {
	return r.db.Create(uni).Error
}

func (r *uniRepository) GetByID(ctx echo.Context, ID database.PID, useCache bool) (*models.Uni, error) {
	var uni models.Uni
	if err := r.db.First(&uni, ID).Error; err != nil {
		return nil, err
	}
	return &uni, nil
}

func (r *uniRepository) Update(uni *models.Uni) error {
	return r.db.Save(uni).Error
}

func (r *uniRepository) Delete(ID database.PID) error {
	return r.db.Delete(&models.Uni{}, ID).Error
}

func (r *uniRepository) GetAll(ctx echo.Context, request models.FetchUniRequest) ([]models.Uni, *helpers.PaginateTemplate, error) {
	var unis []models.Uni
	query := r.db.Model(&models.Uni{})

	// Apply pagination
	paginate := helpers.NewPaginateTemplate(request.Page, request.Limit)
	query = paginate.Paginate(query)

	// Apply includes
	if len(request.Includes) > 0 {
		for _, include := range request.Includes {
			query = query.Preload(include)
		}
	}

	if err := query.Find(&unis).Error; err != nil {
		return nil, nil, err
	}

	return unis, paginate, nil
}
