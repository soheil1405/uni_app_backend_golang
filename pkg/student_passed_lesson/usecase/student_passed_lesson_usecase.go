package usecases

import (
	"uni_app/database"
	"uni_app/models"
	repositories "uni_app/pkg/student_passed_lesson/repository"
	"uni_app/services/env"

	"github.com/labstack/echo/v4"
)

type StudentPassedLessonUsecase interface {
	AddPassedLesson(passedLesson *models.StudentPassedLesson) error
	GetPassedLessonByID(ctx echo.Context, ID database.PID, useCache bool) (*models.StudentPassedLesson, error)
	UpdatePassedLesson(passedLesson *models.StudentPassedLesson) error
	DeletePassedLesson(ID database.PID) error
	GetAllPassedLessons(ctx echo.Context, request models.FetchStudentPassedLessonRequest) ([]models.StudentPassedLesson, error)
	GetStudentPassedLessons(studentID database.PID) ([]models.StudentPassedLesson, error)
}

type studentPassedLessonUsecase struct {
	repo   repositories.StudentPassedLessonRepository
	config *env.Config
}

func NewStudentPassedLessonUsecase(repo repositories.StudentPassedLessonRepository, config *env.Config) StudentPassedLessonUsecase {
	return &studentPassedLessonUsecase{repo, config}
}

func (u *studentPassedLessonUsecase) AddPassedLesson(passedLesson *models.StudentPassedLesson) error {
	return u.repo.Create(passedLesson)
}

func (u *studentPassedLessonUsecase) GetPassedLessonByID(ctx echo.Context, ID database.PID, useCache bool) (*models.StudentPassedLesson, error) {
	return u.repo.GetByID(ctx, ID, useCache)
}

func (u *studentPassedLessonUsecase) UpdatePassedLesson(passedLesson *models.StudentPassedLesson) error {
	return u.repo.Update(passedLesson)
}

func (u *studentPassedLessonUsecase) DeletePassedLesson(ID database.PID) error {
	return u.repo.Delete(ID)
}

func (u *studentPassedLessonUsecase) GetAllPassedLessons(ctx echo.Context, request models.FetchStudentPassedLessonRequest) ([]models.StudentPassedLesson, error) {
	return u.repo.GetAll(ctx, request)
}

func (u *studentPassedLessonUsecase) GetStudentPassedLessons(studentID database.PID) ([]models.StudentPassedLesson, error) {
	return u.repo.GetByStudentID(studentID)
}
