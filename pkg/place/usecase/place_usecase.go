package usecases

import (
	"uni_app/models"
	repositories "uni_app/pkg/place/repository"
)

type PlaceUsecase interface {
	CreatePlace(place *models.Place) error
	GetPlaceByID(id uint) (*models.Place, error)
	UpdatePlace(place *models.Place) error
	DeletePlace(id uint) error
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

func (u *placeUsecase) GetPlaceByID(id uint) (*models.Place, error) {
	return u.repo.GetByID(id)
}

func (u *placeUsecase) UpdatePlace(place *models.Place) error {
	return u.repo.Update(place)
}

func (u *placeUsecase) DeletePlace(id uint) error {
	return u.repo.Delete(id)
}

func (u *placeUsecase) GetAllPlaces() ([]models.Place, error) {
	return u.repo.GetAll()
}
