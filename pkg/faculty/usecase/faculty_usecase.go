package usecases

import (
	"uni_app/database"
	"uni_app/models"
	repositories "uni_app/pkg/faculty/repository"
)

type FacultyUsecase interface {
	CreateFaculty(faculty *models.Faculty) error
	GetFacultyByID(ID database.PID) (*models.Faculty, error)
	UpdateFaculty(faculty *models.Faculty) error
	DeleteFaculty(ID database.PID) error
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

func (u *facultyUsecase) GetFacultyByID(ID database.PID) (*models.Faculty, error) {
	return u.repo.GetByID(ID)
}

func (u *facultyUsecase) UpdateFaculty(faculty *models.Faculty) error {
	return u.repo.Update(faculty)
}

func (u *facultyUsecase) DeleteFaculty(ID database.PID) error {
	return u.repo.Delete(ID)
}

func (u *facultyUsecase) GetAllFaculties() ([]models.Faculty, error) {
	return u.repo.GetAll()
}
