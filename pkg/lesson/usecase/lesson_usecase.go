package usecase

import (
	"uni_app/database"
	"uni_app/models"
	repository "uni_app/pkg/lesson/repository"

	"github.com/labstack/echo/v4"
)

type LessonUsecase interface {
	CreateLesson(lesson *models.Lesson) error
	GetLessonByID(ctx echo.Context, ID database.PID, useCache bool) (*models.Lesson, error)
	UpdateLesson(lesson *models.Lesson) error
	DeleteLesson(ID database.PID) error
	GetAllLessons() ([]models.Lesson, error)
}

type lessonUsecase struct {
	repo repository.LessonRepository
}

func NewLessonUsecase(repo repository.LessonRepository) LessonUsecase {
	return &lessonUsecase{repo}
}

func (u *lessonUsecase) CreateLesson(lesson *models.Lesson) error {
	return u.repo.Create(lesson)
}

func (u *lessonUsecase) GetLessonByID(ctx echo.Context, ID database.PID, useCache bool) (*models.Lesson, error) {
	return u.repo.GetByID(ctx, ID, useCache)
}

func (u *lessonUsecase) UpdateLesson(lesson *models.Lesson) error {
	return u.repo.Update(lesson)
}

func (u *lessonUsecase) DeleteLesson(ID database.PID) error {
	return u.repo.Delete(ID)
}

func (u *lessonUsecase) GetAllLessons() ([]models.Lesson, error) {
	return u.repo.GetAll()
} 