package repositories

import (
	"uni_app/models"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type RouteRepository interface {
	FetchRoute(ctx echo.Context, route *models.Route) (*models.Route, error)
}

type authRepository struct {
	db *gorm.DB
}

func NewRouteRepository(db *gorm.DB) RouteRepository {
	return &authRepository{db}
}

func (a *authRepository) FetchRoute(ctx echo.Context, route *models.Route) (*models.Route, error) {
	if err := a.db.Find(&route).Error; err != nil {
		return nil, err
	}
	return route, nil
}
