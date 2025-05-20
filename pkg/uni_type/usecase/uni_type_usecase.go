package usecase

import (
	"uni_app/database"
	"uni_app/models"
	"uni_app/pkg/uni_type/repository"

	"github.com/labstack/echo/v4"
)

type UniTypeUsecase interface {
	CreateUniType(uniType *models.UniType) error
	GetUniTypeByID(ctx echo.Context, ID database.PID, useCache bool) (*models.UniType, error)
	UpdateUniType(uniType *models.UniType) error
	DeleteUniType(ID database.PID) error
	GetAllUniTypes() ([]models.UniType, error)
}

type uniTypeUsecase struct {
	repo repository.UniTypeRepository
}

func NewUniTypeUsecase(repo repository.UniTypeRepository) UniTypeUsecase {
	return &uniTypeUsecase{repo}
}

func (u *uniTypeUsecase) CreateUniType(uniType *models.UniType) error {
	return u.repo.Create(uniType)
}

func (u *uniTypeUsecase) GetUniTypeByID(ctx echo.Context, ID database.PID, useCache bool) (*models.UniType, error) {
	return u.repo.GetByID(ctx, ID, useCache)
}

func (u *uniTypeUsecase) UpdateUniType(uniType *models.UniType) error {
	return u.repo.Update(uniType)
}

func (u *uniTypeUsecase) DeleteUniType(ID database.PID) error {
	return u.repo.Delete(ID)
}

func (u *uniTypeUsecase) GetAllUniTypes() ([]models.UniType, error) {
	return u.repo.GetAll()
} 