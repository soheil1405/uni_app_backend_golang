package usecase

import (
	"uni_app/database"
	"uni_app/models"
	repository "uni_app/pkg/place/repository"
	"uni_app/utils/templates"

	"github.com/labstack/echo/v4"
)

type PlaceUsecase interface {
	CreatePlace(place *models.Place) error
	GetPlaceByID(ctx echo.Context, ID database.PID, useCache bool) (*models.Place, error)
	UpdatePlace(place *models.Place) error
	DeletePlace(ID database.PID) error
	GetAllPlaces(ctx echo.Context, request models.FetchPlaceRequest) ([]models.Place, *templates.PaginateTemplate, error)
}

type placeUsecase struct {
	repo repository.PlaceRepository
}

func NewPlaceUsecase(repo repository.PlaceRepository) PlaceUsecase {
	return &placeUsecase{repo}
}

func (u *placeUsecase) CreatePlace(place *models.Place) error {
	return u.repo.Create(place)
}

func (u *placeUsecase) GetPlaceByID(ctx echo.Context, ID database.PID, useCache bool) (*models.Place, error) {
	return u.repo.GetByID(ctx, ID, useCache)
}

func (u *placeUsecase) UpdatePlace(place *models.Place) error {
	return u.repo.Update(place)
}

func (u *placeUsecase) DeletePlace(ID database.PID) error {
	return u.repo.Delete(ID)
}

func (u *placeUsecase) GetAllPlaces(ctx echo.Context, request models.FetchPlaceRequest) ([]models.Place, *templates.PaginateTemplate, error) {
	return u.repo.GetAll(ctx, request)
}
