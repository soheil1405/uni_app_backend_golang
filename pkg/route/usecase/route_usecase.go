package usecase

import (
	"uni_app/database"
	"uni_app/models"
	repository "uni_app/pkg/route/repository"
	"uni_app/utils/helpers"

	"github.com/labstack/echo/v4"
)

type RouteUsecase interface {
	CreateRoute(route *models.Route) error
	GetRouteByID(ctx echo.Context, ID database.PID, useCache bool) (*models.Route, error)
	UpdateRoute(route *models.Route) error
	DeleteRoute(ID database.PID) error
	GetAllRoutes(ctx echo.Context, request models.FetchRouteRequest) ([]models.Route, *helpers.PaginateTemplate, error)
}

type routeUsecase struct {
	repo repository.RouteRepository
}

func NewRouteUsecase(repo repository.RouteRepository) RouteUsecase {
	return &routeUsecase{repo}
}

func (u *routeUsecase) CreateRoute(route *models.Route) error {
	return u.repo.Create(route)
}

func (u *routeUsecase) GetRouteByID(ctx echo.Context, ID database.PID, useCache bool) (*models.Route, error) {
	return u.repo.GetByID(ctx, ID, useCache)
}

func (u *routeUsecase) UpdateRoute(route *models.Route) error {
	return u.repo.Update(route)
}

func (u *routeUsecase) DeleteRoute(ID database.PID) error {
	return u.repo.Delete(ID)
}

func (u *routeUsecase) GetAllRoutes(ctx echo.Context, request models.FetchRouteRequest) ([]models.Route, *helpers.PaginateTemplate, error) {
	return u.repo.GetAll(ctx, request)
}
