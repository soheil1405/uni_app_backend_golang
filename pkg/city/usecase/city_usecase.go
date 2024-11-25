package usecases

import (
	"uni_app/models"
	repositories "uni_app/pkg/city/repository"
)

type CityUsecase interface {
	CreateCity(city *models.City) error
	GetCityByID(id uint) (*models.City, error)
	UpdateCity(city *models.City) error
	DeleteCity(id uint) error
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

func (u *cityUsecase) GetCityByID(id uint) (*models.City, error) {
	return u.repo.GetByID(id)
}

func (u *cityUsecase) UpdateCity(city *models.City) error {
	return u.repo.Update(city)
}

func (u *cityUsecase) DeleteCity(id uint) error {
	return u.repo.Delete(id)
}

func (u *cityUsecase) GetAllCities() ([]models.City, error) {
	return u.repo.GetAll()
}
