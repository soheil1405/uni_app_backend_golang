package usecase

import (
	"uni_app/database"
	"uni_app/models"
	repository "uni_app/pkg/degree/repository"

	"github.com/labstack/echo/v4"
)

type DegreeUsecase interface {
	CreateDegree(degree *models.DegreeLevel) error
	GetDegreeByID(ctx echo.Context, ID database.PID, useCache bool) (*models.DegreeLevel, error)
	UpdateDegree(degree *models.DegreeLevel) error
	DeleteDegree(ID database.PID) error
	GetAllDegrees() ([]models.DegreeLevel, error)
}

type degreeUsecase struct {
	repo repository.DegreeRepository
}

func NewDegreeUsecase(repo repository.DegreeRepository) DegreeUsecase {
	return &degreeUsecase{repo}
}

func (u *degreeUsecase) CreateDegree(degree *models.DegreeLevel) error {
	return u.repo.Create(degree)
}

func (u *degreeUsecase) GetDegreeByID(ctx echo.Context, ID database.PID, useCache bool) (*models.DegreeLevel, error) {
	return u.repo.GetByID(ctx, ID, useCache)
}

func (u *degreeUsecase) UpdateDegree(degree *models.DegreeLevel) error {
	return u.repo.Update(degree)
}

func (u *degreeUsecase) DeleteDegree(ID database.PID) error {
	return u.repo.Delete(ID)
}

func (u *degreeUsecase) GetAllDegrees() ([]models.DegreeLevel, error) {
	return u.repo.GetAll()
} 