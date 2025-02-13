package usecases

import (
	"uni_app/database"
	"uni_app/models"
	repositories "uni_app/pkg/place_type/repository"
)

type PlaceTypeUsecase interface {
	CreatePlaceType(placeType *models.PlaceType) error
	GetPlaceTypeByID(ID database.PID) (*models.PlaceType, error)
	UpdatePlaceType(placeType *models.PlaceType) error
	DeletePlaceType(ID database.PID) error
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

func (u *placeTypeUsecase) GetPlaceTypeByID(ID database.PID) (*models.PlaceType, error) {
	return u.repo.GetByID(ID)
}

func (u *placeTypeUsecase) UpdatePlaceType(placeType *models.PlaceType) error {
	return u.repo.Update(placeType)
}

func (u *placeTypeUsecase) DeletePlaceType(ID database.PID) error {
	return u.repo.Delete(ID)
}

func (u *placeTypeUsecase) GetAllPlaceTypes() ([]models.PlaceType, error) {
	return u.repo.GetAll()
}
