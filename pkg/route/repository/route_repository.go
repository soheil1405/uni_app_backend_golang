package repository

import (
	"uni_app/database"
	"uni_app/models"
	"uni_app/utils/helpers"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type RouteRepository interface {
	Create(route *models.Route) error
	GetByID(ctx echo.Context, ID database.PID, useCache bool) (*models.Route, error)
	Update(route *models.Route) error
	Delete(ID database.PID) error
	GetAll(ctx echo.Context, request models.FetchRouteRequest) ([]models.Route, *helpers.PaginateTemplate, error)
}

type routeRepository struct {
	db *gorm.DB
}

func NewRouteRepository(db *gorm.DB) RouteRepository {
	return &routeRepository{db}
}

func (r *routeRepository) Create(route *models.Route) error {
	return r.db.Create(route).Error
}

func (r *routeRepository) GetByID(ctx echo.Context, ID database.PID, useCache bool) (*models.Route, error) {
	var route models.Route
	if err := r.db.First(&route, ID).Error; err != nil {
		return nil, err
	}
	return &route, nil
}

func (r *routeRepository) Update(route *models.Route) error {
	return r.db.Save(route).Error
}

func (r *routeRepository) Delete(ID database.PID) error {
	return r.db.Delete(&models.Route{}, ID).Error
}

func (r *routeRepository) GetAll(ctx echo.Context, request models.FetchRouteRequest) ([]models.Route, *helpers.PaginateTemplate, error) {
	var routes []models.Route
	query := r.db.Model(&models.Route{})

	// Apply pagination
	paginate := helpers.NewPaginateTemplate(request.Page, request.Limit)
	query = paginate.Paginate(query)

	// Apply includes
	if len(request.Includes) > 0 {
		for _, include := range request.Includes {
			query = query.Preload(include)
		}
	}

	if err := query.Find(&routes).Error; err != nil {
		return nil, nil, err
	}

	return routes, paginate, nil
}
