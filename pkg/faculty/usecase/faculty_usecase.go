package usecases

import (
	"uni_app/models"
	repositories "uni_app/pkg/faculty/repository"
)

type FacultyUsecase interface {
	CreateFaculty(faculty *models.Faculty) error
	GetFacultyByID(id uint) (*models.Faculty, error)
	UpdateFaculty(faculty *models.Faculty) error
	DeleteFaculty(id uint) error
	GetAllFaculties() ([]models.Faculty, error)
}

type facultyUsecase struct {
	repo repositories.FacultyRepository
}

func NewFacultyUsecase(repo repositories.FacultyRepository) FacultyUsecase {
	return &facultyUsecase{repo}
}

func (u *facultyUsecase) CreateFaculty(faculty *models.Faculty) error {
	return u.repo.Create(faculty)
}

func (u *facultyUsecase) GetFacultyByID(id uint) (*models.Faculty, error) {
	return u.repo.GetByID(id)
}

func (u *facultyUsecase) UpdateFaculty(faculty *models.Faculty) error {
	return u.repo.Update(faculty)
}

func (u *facultyUsecase) DeleteFaculty(id uint) error {
	return u.repo.Delete(id)
}

func (u *facultyUsecase) GetAllFaculties() ([]models.Faculty, error) {
	return u.repo.GetAll()
}
