package usecase

import (
	"uni_app/database"
	"uni_app/models"
	"uni_app/pkg/department/repository"
	"uni_app/services/env"
)

type DepartmentUsecase interface {
	Create(department *models.Department) error
	Update(department *models.Department) error
	Delete(id database.PID) error
	GetByID(id database.PID) (*models.Department, error)
	List(filters *models.FetchDepartmentRequest) ([]*models.Department, error)
	SetHead(departmentID, teacherID database.PID) error
	AddTeacher(departmentID, teacherID database.PID) error
	RemoveTeacher(departmentID, teacherID database.PID) error
	AddMajor(departmentID, majorID database.PID) error
	RemoveMajor(departmentID, majorID database.PID) error
}

type departmentUsecase struct {
	repo   repository.DepartmentRepository
	config *env.Config
}

func NewDepartmentUsecase(repo repository.DepartmentRepository, config *env.Config) DepartmentUsecase {
	return &departmentUsecase{
		repo:   repo,
		config: config,
	}
}

func (u *departmentUsecase) Create(department *models.Department) error {
	return u.repo.Create(department)
}

func (u *departmentUsecase) Update(department *models.Department) error {
	return u.repo.Update(department)
}

func (u *departmentUsecase) Delete(id database.PID) error {
	return u.repo.Delete(id)
}

func (u *departmentUsecase) GetByID(id database.PID) (*models.Department, error) {
	return u.repo.GetByID(id)
}

func (u *departmentUsecase) List(filters *models.FetchDepartmentRequest) ([]*models.Department, error) {
	return u.repo.List(filters)
}

func (u *departmentUsecase) SetHead(departmentID, teacherID database.PID) error {
	return u.repo.SetHead(departmentID, teacherID)
}

func (u *departmentUsecase) AddTeacher(departmentID, teacherID database.PID) error {
	return u.repo.AddTeacher(departmentID, teacherID)
}

func (u *departmentUsecase) RemoveTeacher(departmentID, teacherID database.PID) error {
	return u.repo.RemoveTeacher(departmentID, teacherID)
}

func (u *departmentUsecase) AddMajor(departmentID, majorID database.PID) error {
	return u.repo.AddMajor(departmentID, majorID)
}

func (u *departmentUsecase) RemoveMajor(departmentID, majorID database.PID) error {
	return u.repo.RemoveMajor(departmentID, majorID)
}
