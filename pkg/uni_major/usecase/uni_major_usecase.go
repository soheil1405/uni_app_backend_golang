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
	GetUniMajorsByUniversityID(universityID database.PID) ([]models.UniMajor, error)
	GetUniMajorsByMajorID(majorID database.PID) ([]models.UniMajor, error)
}

type uniMajorUsecase struct {
	repo repository.UniMajorRepository
}

func NewUniMajorUsecase(repo repository.UniMajorRepository) UniMajorUsecase {
	return &uniMajorUsecase{repo}
}

func (u *uniMajorUsecase) CreateUniMajor(uniMajor *models.UniMajor) error {
	// Validate required fields
	if uniMajor.UniversityID == 0 {
		return errors.New("university ID is required")
	}
	if uniMajor.MajorID == 0 {
		return errors.New("major ID is required")
	}
	if uniMajor.DegreeLevelID == 0 {
		return errors.New("degree level ID is required")
	}

	// Check if the combination already exists
	existingMajors, err := u.repo.GetAll()
	if err != nil {
		return err
	}

	for _, existing := range existingMajors {
		if existing.UniversityID == uniMajor.UniversityID &&
			existing.MajorID == uniMajor.MajorID &&
			existing.DegreeLevelID == uniMajor.DegreeLevelID {
			return errors.New("this major already exists for this university and degree level")
		}
	}

	return u.repo.Create(uniMajor)
}

func (u *uniMajorUsecase) GetUniMajorByID(ctx echo.Context, ID database.PID, useCache bool) (*models.UniMajor, error) {
	return u.repo.GetByID(ctx, ID, useCache)
}

func (u *uniMajorUsecase) UpdateUniMajor(uniMajor *models.UniMajor) error {
	// Validate required fields
	if uniMajor.UniversityID == 0 {
		return errors.New("university ID is required")
	}
	if uniMajor.MajorID == 0 {
		return errors.New("major ID is required")
	}
	if uniMajor.DegreeLevelID == 0 {
		return errors.New("degree level ID is required")
	}

	// Check if the combination already exists (excluding current record)
	existingMajors, err := u.repo.GetAll()
	if err != nil {
		return err
	}

	for _, existing := range existingMajors {
		if existing.ID != uniMajor.ID &&
			existing.UniversityID == uniMajor.UniversityID &&
			existing.MajorID == uniMajor.MajorID &&
			existing.DegreeLevelID == uniMajor.DegreeLevelID {
			return errors.New("this major already exists for this university and degree level")
		}
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

func (u *uniMajorUsecase) GetUniMajorsByUniversityID(universityID database.PID) ([]models.UniMajor, error) {
	allMajors, err := u.repo.GetAll()
	if err != nil {
		return nil, err
	}

	var filteredMajors []models.UniMajor
	for _, major := range allMajors {
		if major.UniversityID == universityID {
			filteredMajors = append(filteredMajors, major)
		}
	}

	return filteredMajors, nil
}

func (u *uniMajorUsecase) GetUniMajorsByMajorID(majorID database.PID) ([]models.UniMajor, error) {
	allMajors, err := u.repo.GetAll()
	if err != nil {
		return nil, err
	}

	var filteredMajors []models.UniMajor
	for _, major := range allMajors {
		if major.MajorID == majorID {
			filteredMajors = append(filteredMajors, major)
		}
	}

	return filteredMajors, nil
} 