package usecase

import (
	"errors"
	"uni_app/database"
	"uni_app/models"
	repository "uni_app/pkg/uni_major/repository"

	"github.com/labstack/echo/v4"
)

type UniMajorUsecase interface {
	CreateUniMajor(uniMajor *models.UniMajor) error
	GetUniMajorByID(ctx echo.Context, ID database.PID, useCache bool) (*models.UniMajor, error)
	UpdateUniMajor(uniMajor *models.UniMajor) error
	DeleteUniMajor(ID database.PID) error
	GetAllUniMajors() ([]models.UniMajor, error)
}

type uniMajorUsecase struct {
	repo repository.UniMajorRepository
}

func NewUniMajorUsecase(repo repository.UniMajorRepository) UniMajorUsecase {
	return &uniMajorUsecase{repo}
}

func (u *uniMajorUsecase) CreateUniMajor(uniMajor *models.UniMajor) error {
	if uniMajor.MajorID == 0 {
		return errors.New("major ID is required")
	}
	return u.repo.Create(uniMajor)
}

func (u *uniMajorUsecase) GetUniMajorByID(ctx echo.Context, ID database.PID, useCache bool) (*models.UniMajor, error) {
	return u.repo.GetByID(ctx, ID, useCache)
}

func (u *uniMajorUsecase) UpdateUniMajor(uniMajor *models.UniMajor) error {
	if uniMajor.MajorID == 0 {
		return errors.New("major ID is required")
	}
	return u.repo.Update(uniMajor)
}

func (u *uniMajorUsecase) DeleteUniMajor(ID database.PID) error {
	// Check if there are any related records before deletion
	// This could be expanded based on your requirements
	return u.repo.Delete(ID)
}

func (u *uniMajorUsecase) GetAllUniMajors() ([]models.UniMajor, error) {
	return u.repo.GetAll()
}
