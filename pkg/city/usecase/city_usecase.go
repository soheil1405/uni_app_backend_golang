package usecases

import (
	"uni_app/database"
	"uni_app/models"
	repositories "uni_app/pkg/city/repository"

	"github.com/labstack/echo/v4"
)

type CityUsecase interface {
	CreateCity(city *models.City) error
	GetCityByID(ctx echo.Context, ID database.PID, useCache bool) (*models.City, error)
	UpdateCity(city *models.City) error
	DeleteCity(ID database.PID) error
	GetAllCities() ([]models.City, error)
}

type cityUsecase struct {
	repo repositories.CityRepository
}

func NewCityUsecase(repo repositories.CityRepository) CityUsecase {
	return &cityUsecase{repo}
}

func (u *cityUsecase) CreateCity(city *models.City) error {
	return u.repo.Create(city)
}

func (u *cityUsecase) GetCityByID(ctx echo.Context, ID database.PID, useCache bool) (*models.City, error) {
	return u.repo.GetByID(ctx, ID, useCache)
}

func (u *cityUsecase) UpdateCity(city *models.City) error {
	return u.repo.Update(city)
}

func (u *cityUsecase) DeleteCity(ID database.PID) error {
	return u.repo.Delete(ID)
}

func (u *cityUsecase) GetAllCities() ([]models.City, error) {
	return u.repo.GetAll()
}
