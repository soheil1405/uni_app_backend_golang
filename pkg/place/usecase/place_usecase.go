package usecases

import (
	"uni_app/database"
	"uni_app/models"
	repositories "uni_app/pkg/place/repository"
)

type PlaceUsecase interface {
	CreatePlace(place *models.Place) error
	GetPlaceByID(ID database.PID) (*models.Place, error)
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

func (u *placeUsecase) GetPlaceByID(ID database.PID) (*models.Place, error) {
	return u.repo.GetByID(ID)
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
