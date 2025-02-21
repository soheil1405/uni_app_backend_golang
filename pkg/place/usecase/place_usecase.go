package usecases

import (
	"uni_app/database"
	"uni_app/models"
	repositories "uni_app/pkg/place/repository"

	"github.com/labstack/echo/v4"
)

type PlaceUsecase interface {
	CreatePlace(place *models.Place) error
	GetPlaceByID(ctx echo.Context, ID database.PID, useCache bool) (*models.Place, error)
	UpdatePlace(place *models.Place) error
	DeletePlace(ID database.PID) error
	GetAllPlaces() ([]models.Place, error)
}

type placeUsecase struct {
	repo repositories.PlaceRepository
}

func NewPlaceUsecase(repo repositories.PlaceRepository) PlaceUsecase {
	return &placeUsecase{repo}
}

func (u *placeUsecase) CreatePlace(place *models.Place) error {
	return u.repo.Create(place)
}

func (u *placeUsecase) GetPlaceByID(ctx echo.Context, ID database.PID, useCache bool) (*models.Place, error) {
	return u.repo.GetByID(ctx, ID, false)
}

func (u *placeUsecase) UpdatePlace(place *models.Place) error {
	return u.repo.Update(place)
}

func (u *placeUsecase) DeletePlace(ID database.PID) error {
	return u.repo.Delete(ID)
}

func (u *placeUsecase) GetAllPlaces() ([]models.Place, error) {
	return u.repo.GetAll()
}
