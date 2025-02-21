package usecases

import (
	"uni_app/database"
	"uni_app/models"
	repositories "uni_app/pkg/major/repository"

	"github.com/labstack/echo/v4"
)

type MajorUsecase interface {
	CreateMajor(major *models.Major) error
	GetMajorByID(ctx echo.Context, ID database.PID, useCache bool) (*models.Major, error)
	UpdateMajor(major *models.Major) error
	DeleteMajor(ID database.PID) error
	GetAllMajors() ([]models.Major, error)
}

type majorUsecase struct {
	repo repositories.MajorRepository
}

func NewMajorUsecase(repo repositories.MajorRepository) MajorUsecase {
	return &majorUsecase{repo}
}

func (u *majorUsecase) CreateMajor(major *models.Major) error {
	return u.repo.Create(major)
}

func (u *majorUsecase) GetMajorByID(ctx echo.Context, ID database.PID, useCache bool) (*models.Major, error) {
	return u.repo.GetByID(ctx, ID, useCache)
}

func (u *majorUsecase) UpdateMajor(major *models.Major) error {
	return u.repo.Update(major)
}

func (u *majorUsecase) DeleteMajor(ID database.PID) error {
	return u.repo.Delete(ID)
}

func (u *majorUsecase) GetAllMajors() ([]models.Major, error) {
	return u.repo.GetAll()
}
