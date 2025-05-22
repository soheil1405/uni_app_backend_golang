package repositories

import (
	"uni_app/database"
	"uni_app/models"
	"uni_app/utils/helpers"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type MajorRepository interface {
	Create(major *models.Major) error
	GetByID(ctx echo.Context, ID database.PID, useCache bool) (*models.Major, error)
	Update(major *models.Major) error
	Delete(ID database.PID) error
	GetAll(ctx echo.Context, request models.FetchMajorRequest) ([]models.Major, *helpers.PaginateTemplate, error)
}

type majorRepository struct {
	db *gorm.DB
}

func NewMajorRepository(db *gorm.DB) MajorRepository {
	return &majorRepository{db}
}

func (r *majorRepository) Create(major *models.Major) error {
	return r.db.Create(major).Error
}

func (r *majorRepository) GetByID(ctx echo.Context, ID database.PID, useCache bool) (*models.Major, error) {
	var major models.Major
	if err := r.db.First(&major, ID).Error; err != nil {
		return nil, err
	}
	return &major, nil
}

func (r *majorRepository) Update(major *models.Major) error {
	return r.db.Save(major).Error
}

func (r *majorRepository) Delete(ID database.PID) error {
	return r.db.Delete(&models.Major{}, ID).Error
}

func (r *majorRepository) GetAll(ctx echo.Context, request models.FetchMajorRequest) ([]models.Major, *helpers.PaginateTemplate, error) {
	var majors []models.Major
	query := r.db.Model(&models.Major{})

	// Apply pagination
	paginate := helpers.NewPaginateTemplate(request.Page, request.Limit)
	query = paginate.Paginate(query)

	// Apply includes
	if len(request.Includes) > 0 {
		for _, include := range request.Includes {
			query = query.Preload(include)
		}
	}

	if err := query.Find(&majors).Error; err != nil {
		return nil, nil, err
	}

	return majors, paginate, nil
}
