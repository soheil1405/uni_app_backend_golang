package repositories

import (
	"uni_app/database"
	"uni_app/models"
	"uni_app/utils/helpers"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type CityRepository interface {
	Create(city *models.City) error
	GetByID(ctx echo.Context, ID database.PID, useCache bool) (*models.City, error)
	Update(city *models.City) error
	Delete(ID database.PID) error
	GetAll(ctx echo.Context, request models.FetchCityRequest) ([]models.City, *helpers.PaginateTemplate, error)
}

type cityRepository struct {
	db *gorm.DB
}

func NewCityRepository(db *gorm.DB) CityRepository {
	return &cityRepository{db}
}

func (r *cityRepository) Create(city *models.City) error {
	return r.db.Create(city).Error
}

func (r *cityRepository) GetByID(ctx echo.Context, ID database.PID, useCache bool) (*models.City, error) {
	var city models.City
	if err := r.db.First(&city, ID).Error; err != nil {
		return nil, err
	}
	return &city, nil
}

func (r *cityRepository) Update(city *models.City) error {
	return r.db.Save(city).Error
}

func (r *cityRepository) Delete(ID database.PID) error {
	return r.db.Delete(&models.City{}, ID).Error
}

func (r *cityRepository) GetAll(ctx echo.Context, request models.FetchCityRequest) ([]models.City, *helpers.PaginateTemplate, error) {
	var cities []models.City
	query := r.db.Model(&models.City{})

	// Apply pagination
	paginate := helpers.NewPaginateTemplate(request.Page, request.Limit)
	query = paginate.Paginate(query)

	// Apply includes
	if len(request.Includes) > 0 {
		for _, include := range request.Includes {
			query = query.Preload(include)
		}
	}

	if err := query.Find(&cities).Error; err != nil {
		return nil, nil, err
	}

	return cities, paginate, nil
}
