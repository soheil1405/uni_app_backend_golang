package repositories

import (
	"uni_app/database"
	"uni_app/models"
	"uni_app/utils/templates"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type PlaceRepository interface {
	Create(place *models.Place) error
	GetByID(ctx echo.Context, ID database.PID, useCache bool) (*models.Place, error)
	Update(place *models.Place) error
	Delete(ID database.PID) error
	GetAll(ctx echo.Context, request models.FetchPlaceRequest) ([]models.Place, *templates.PaginateTemplate, error)
}

type placeRepository struct {
	db *gorm.DB
}

func NewPlaceRepository(db *gorm.DB) PlaceRepository {
	return &placeRepository{db}
}

func (r *placeRepository) Create(place *models.Place) error {
	return r.db.Create(place).Error
}

func (r *placeRepository) GetByID(ctx echo.Context, ID database.PID, useCache bool) (*models.Place, error) {
	var place models.Place
	if err := r.db.First(&place, ID).Error; err != nil {
		return nil, err
	}
	return &place, nil
}

func (r *placeRepository) Update(place *models.Place) error {
	return r.db.Save(place).Error
}

func (r *placeRepository) Delete(ID database.PID) error {
	return r.db.Delete(&models.Place{}, ID).Error
}

func (r *placeRepository) GetAll(ctx echo.Context, request models.FetchPlaceRequest) ([]models.Place, *templates.PaginateTemplate, error) {
	var (
		places   []models.Place
		query    = r.db
		limit    = request.Limit
		offset   = request.Offset
		includes = request.Includes
		total    int64
		sorts    = request.Sorts
	)

	if request.CityID != 0 {
		query = query.Where("city_id = ?", request.CityID)
	}

	if request.PlaceTypeID != 0 {
		query = query.Where("place_type_id = ?", request.PlaceTypeID)
	}

	if request.Search != "" {
		query = query.Where("name LIKE ?", "%"+request.Search+"%")
	}

	if err := query.Model(&models.Place{}).Count(&total).Error; err != nil {
		return nil, nil, err
	}
	for _, sort := range sorts {
		query = query.Order(sort)
	}

	for _, include := range includes {
		query = query.Preload(include)
	}

	for {
		query.Limit(limit).Offset(offset).Find(&places)
		if limit > len(places) && int(total) > offset+limit {
			offset += limit
		} else {
			break
		}
	}

	meta := templates.CreatePaginateTemplate(int(total), offset, limit)

	return places, meta, nil
}
