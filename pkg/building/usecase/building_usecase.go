package usecase

import (
	"uni_app/database"
	"uni_app/models"
	"uni_app/pkg/building/repository"
	"uni_app/services/env"

	"github.com/google/uuid"
)

type BuildingUsecase interface {
	Create(building *models.Building) error
	Update(building *models.Building) error
	Delete(id database.PID) error
	GetByID(id database.PID) (*models.Building, error)
	List(filters *models.FetchBuildingRequest) ([]*models.Building, error)
}

type buildingUsecase struct {
	repo   repository.BuildingRepository
	config *env.Config
}

func NewBuildingUsecase(repo repository.BuildingRepository, config *env.Config) BuildingUsecase {
	return &buildingUsecase{
		repo:   repo,
		config: config,
	}
}

func (u *buildingUsecase) Create(building *models.Building) error {
	id, err := database.ParsePID(uuid.New().String())
	if err != nil {
		return err
	}
	building.ID = id
	return u.repo.Create(building)
}

func (u *buildingUsecase) Update(building *models.Building) error {
	return u.repo.Update(building)
}

func (u *buildingUsecase) Delete(id database.PID) error {
	return u.repo.Delete(id)
}

func (u *buildingUsecase) GetByID(id database.PID) (*models.Building, error) {
	return u.repo.GetByID(id)
}

func (u *buildingUsecase) List(filters *models.FetchBuildingRequest) ([]*models.Building, error) {
	return u.repo.List(filters)
}
