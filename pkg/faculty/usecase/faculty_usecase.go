package usecase

import (
	"uni_app/database"
	"uni_app/models"
	"uni_app/pkg/faculty/repository"
	"uni_app/services/env"
)

type FacultyUsecase interface {
	Create(faculty *models.Faculty) error
	Update(faculty *models.Faculty) error
	Delete(id database.PID) error
	GetByID(id database.PID) (*models.Faculty, error)
	List(filters *models.FetchFacultyRequest) ([]*models.Faculty, error)
	AddDepartment(facultyID, departmentID database.PID) error
	RemoveDepartment(facultyID, departmentID database.PID) error
	AddTeacher(facultyID, teacherID database.PID) error
	RemoveTeacher(facultyID, teacherID database.PID) error
	AddStaff(facultyID, userID database.PID) error
	RemoveStaff(facultyID, userID database.PID) error
}

type facultyUsecase struct {
	repo   repository.FacultyRepository
	config *env.Config
}

func NewFacultyUsecase(repo repository.FacultyRepository, config *env.Config) FacultyUsecase {
	return &facultyUsecase{
		repo:   repo,
		config: config,
	}
}

func (u *facultyUsecase) Create(faculty *models.Faculty) error {
	return u.repo.Create(faculty)
}

func (u *facultyUsecase) Update(faculty *models.Faculty) error {
	return u.repo.Update(faculty)
}

func (u *facultyUsecase) Delete(id database.PID) error {
	return u.repo.Delete(id)
}

func (u *facultyUsecase) GetByID(id database.PID) (*models.Faculty, error) {
	return u.repo.GetByID(id)
}

func (u *facultyUsecase) List(filters *models.FetchFacultyRequest) ([]*models.Faculty, error) {
	return u.repo.List(filters)
}

func (u *facultyUsecase) AddDepartment(facultyID, departmentID database.PID) error {
	return u.repo.AddDepartment(facultyID, departmentID)
}

func (u *facultyUsecase) RemoveDepartment(facultyID, departmentID database.PID) error {
	return u.repo.RemoveDepartment(facultyID, departmentID)
}

func (u *facultyUsecase) AddTeacher(facultyID, teacherID database.PID) error {
	return u.repo.AddTeacher(facultyID, teacherID)
}

func (u *facultyUsecase) RemoveTeacher(facultyID, teacherID database.PID) error {
	return u.repo.RemoveTeacher(facultyID, teacherID)
}

func (u *facultyUsecase) AddStaff(facultyID, userID database.PID) error {
	return u.repo.AddStaff(facultyID, userID)
}

func (u *facultyUsecase) RemoveStaff(facultyID, userID database.PID) error {
	return u.repo.RemoveStaff(facultyID, userID)
}
