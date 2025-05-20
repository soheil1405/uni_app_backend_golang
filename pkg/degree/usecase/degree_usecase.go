package usecase

import (
	"errors"
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
	GetDegreesByType(degreeType string) ([]models.DegreeLevel, error)
}

type degreeUsecase struct {
	repo repository.DegreeRepository
}

func NewDegreeUsecase(repo repository.DegreeRepository) DegreeUsecase {
	return &degreeUsecase{repo}
}

func (u *degreeUsecase) CreateDegree(degree *models.DegreeLevel) error {
	// Validate required fields
	if degree.Name == "" {
		return errors.New("degree name is required")
	}
	if degree.Type == "" {
		return errors.New("degree type is required")
	}

	// Check if degree with same name already exists
	existingDegrees, err := u.repo.GetAll()
	if err != nil {
		return err
	}

	for _, existing := range existingDegrees {
		if existing.Name == degree.Name {
			return errors.New("degree with this name already exists")
		}
	}

	return u.repo.Create(degree)
}

func (u *degreeUsecase) GetDegreeByID(ctx echo.Context, ID database.PID, useCache bool) (*models.DegreeLevel, error) {
	return u.repo.GetByID(ctx, ID, useCache)
}

func (u *degreeUsecase) UpdateDegree(degree *models.DegreeLevel) error {
	// Validate required fields
	if degree.Name == "" {
		return errors.New("degree name is required")
	}
	if degree.Type == "" {
		return errors.New("degree type is required")
	}

	// Check if degree with same name already exists (excluding current record)
	existingDegrees, err := u.repo.GetAll()
	if err != nil {
		return err
	}

	for _, existing := range existingDegrees {
		if existing.ID != degree.ID && existing.Name == degree.Name {
			return errors.New("degree with this name already exists")
		}
	}

	return u.repo.Update(degree)
}

func (u *degreeUsecase) DeleteDegree(ID database.PID) error {
	// Check if there are any related records before deletion
	// This could be expanded based on your requirements
	return u.repo.Delete(ID)
}

func (u *degreeUsecase) GetAllDegrees() ([]models.DegreeLevel, error) {
	return u.repo.GetAll()
}

func (u *degreeUsecase) GetDegreesByType(degreeType string) ([]models.DegreeLevel, error) {
	allDegrees, err := u.repo.GetAll()
	if err != nil {
		return nil, err
	}

	var filteredDegrees []models.DegreeLevel
	for _, degree := range allDegrees {
		if degree.Type == degreeType {
			filteredDegrees = append(filteredDegrees, degree)
		}
	}

	return filteredDegrees, nil
} 