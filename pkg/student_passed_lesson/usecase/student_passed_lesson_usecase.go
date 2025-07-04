package usecases

import (
	"uni_app/database"
	"uni_app/models"
	repositories "uni_app/pkg/student_passed_lesson/repository"
	"uni_app/services/env"

	"github.com/labstack/echo/v4"
)

type StudentPassedCourseUsecase interface {
	AddPassedCourse(passedCourse *models.StudentPassedCourse) error
	GetPassedCourseByID(ctx echo.Context, ID database.PID, useCache bool) (*models.StudentPassedCourse, error)
	UpdatePassedCourse(passedCourse *models.StudentPassedCourse) error
	DeletePassedCourse(ID database.PID) error
	GetAllPassedCourses(ctx echo.Context, request models.FetchStudentPassedCourseRequest) ([]models.StudentPassedCourse, error)
	GetStudentPassedCourses(studentID database.PID) ([]models.StudentPassedCourse, error)
}

type studentPassedCourseUsecase struct {
	repo   repositories.StudentPassedCourseRepository
	config *env.Config
}

func NewStudentPassedCourseUsecase(repo repositories.StudentPassedCourseRepository, config *env.Config) StudentPassedCourseUsecase {
	return &studentPassedCourseUsecase{repo, config}
}

func (u *studentPassedCourseUsecase) AddPassedCourse(passedCourse *models.StudentPassedCourse) error {
	return u.repo.Create(passedCourse)
}

func (u *studentPassedCourseUsecase) GetPassedCourseByID(ctx echo.Context, ID database.PID, useCache bool) (*models.StudentPassedCourse, error) {
	return u.repo.GetByID(ctx, ID, useCache)
}

func (u *studentPassedCourseUsecase) UpdatePassedCourse(passedCourse *models.StudentPassedCourse) error {
	return u.repo.Update(passedCourse)
}

func (u *studentPassedCourseUsecase) DeletePassedCourse(ID database.PID) error {
	return u.repo.Delete(ID)
}

func (u *studentPassedCourseUsecase) GetAllPassedCourses(ctx echo.Context, request models.FetchStudentPassedCourseRequest) ([]models.StudentPassedCourse, error) {
	return u.repo.GetAll(ctx, request)
}

func (u *studentPassedCourseUsecase) GetStudentPassedCourses(studentID database.PID) ([]models.StudentPassedCourse, error) {
	return u.repo.GetByStudentID(studentID)
}
