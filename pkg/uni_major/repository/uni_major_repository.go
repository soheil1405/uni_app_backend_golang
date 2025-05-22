package repository

import (
	"uni_app/database"
	"uni_app/models"
	"uni_app/utils/helpers"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type UniMajorRepository interface {
	Create(uniMajor *models.UniMajor) error
	GetByID(ctx echo.Context, ID database.PID, useCache bool) (*models.UniMajor, error)
	Update(uniMajor *models.UniMajor) error
	Delete(ID database.PID) error
	GetAll(ctx echo.Context, request models.FetchUniMajorRequest) ([]models.UniMajor, *helpers.PaginateTemplate, error)
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

func (r *uniMajorRepository) GetAll(ctx echo.Context, request models.FetchUniMajorRequest) ([]models.UniMajor, *helpers.PaginateTemplate, error) {
	var uniMajors []models.UniMajor
	query := r.db.Model(&models.UniMajor{})

	// Apply pagination
	paginate := helpers.NewPaginateTemplate(request.Page, request.Limit)
	query = paginate.Paginate(query)

	// Apply includes
	if len(request.Includes) > 0 {
		for _, include := range request.Includes {
			query = query.Preload(include)
		}
	}

	if err := query.Find(&uniMajors).Error; err != nil {
		return nil, nil, err
	}

	return uniMajors, paginate, nil
}
