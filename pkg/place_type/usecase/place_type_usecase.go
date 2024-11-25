package usecases

import (
	"uni_app/models"
	repositories "uni_app/pkg/place_type/repository"
)

type PlaceTypeUsecase interface {
	CreatePlaceType(placeType *models.PlaceType) error
	GetPlaceTypeByID(id uint) (*models.PlaceType, error)
	UpdatePlaceType(placeType *models.PlaceType) error
	DeletePlaceType(id uint) error
	GetAllPlaceTypes() ([]models.PlaceType, error)
}

type placeTypeUsecase struct {
	repo repositories.PlaceTypeRepository
}

func NewPlaceTypeUsecase(repo repositories.PlaceTypeRepository) PlaceTypeUsecase {
	return &placeTypeUsecase{repo}
}

func (u *placeTypeUsecase) CreatePlaceType(placeType *models.PlaceType) error {
	return u.repo.Create(placeType)
}

func (u *placeTypeUsecase) GetPlaceTypeByID(id uint) (*models.PlaceType, error) {
	return u.repo.GetByID(id)
}

func (u *placeTypeUsecase) UpdatePlaceType(placeType *models.PlaceType) error {
	return u.repo.Update(placeType)
}

func (u *placeTypeUsecase) DeletePlaceType(id uint) error {
	return u.repo.Delete(id)
}

func (u *placeTypeUsecase) GetAllPlaceTypes() ([]models.PlaceType, error) {
	return u.repo.GetAll()
}
