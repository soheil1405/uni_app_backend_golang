package usecases

import (
	"uni_app/database"
	"uni_app/models"
	repositories "uni_app/pkg/student/repository"

	"github.com/labstack/echo/v4"
)

type StudentUsecase interface {
	CreateStudent(student *models.Student) error
	GetStudentByID(ctx echo.Context, ID database.PID, useCache bool) (*models.Student, error)
	UpdateStudent(student *models.Student) error
	DeleteStudent(ID database.PID) error
	GetAllStudents() ([]models.Student, error)
}

type userUsecase struct {
	repo repositories.StudentRepository
}

func NewStudentUsecase(repo repositories.StudentRepository) StudentUsecase {
	return &userUsecase{repo}
}

func (u *userUsecase) CreateStudent(student *models.Student) error {
	return u.repo.Create(student)
}

func (u *userUsecase) GetStudentByID(ctx echo.Context, ID database.PID, useCache bool) (*models.Student, error) {
	return u.repo.GetByID(ctx, ID, useCache)
}

func (u *userUsecase) UpdateStudent(student *models.Student) error {
	return u.repo.Update(student)
}

func (u *userUsecase) DeleteStudent(ID database.PID) error {
	return u.repo.Delete(ID)
}

func (u *userUsecase) GetAllStudents() ([]models.Student, error) {
	return u.repo.GetAll()
}
