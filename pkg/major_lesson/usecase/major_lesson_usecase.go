package usecase

import (
	"uni_app/database"
	"uni_app/models"
	repository "uni_app/pkg/major_lesson/repository"

	"github.com/labstack/echo/v4"
)

type MajorLessonUsecase interface {
	CreateMajorLesson(majorLesson *models.MajorLesson) error
	GetMajorLessonByID(ctx echo.Context, ID database.PID, useCache bool) (*models.MajorLesson, error)
	UpdateMajorLesson(majorLesson *models.MajorLesson) error
	DeleteMajorLesson(ID database.PID) error
	GetAllMajorLessons() ([]models.MajorLesson, error)
}

type majorLessonUsecase struct {
	repo repository.MajorLessonRepository
}

func NewMajorLessonUsecase(repo repository.MajorLessonRepository) MajorLessonUsecase {
	return &majorLessonUsecase{repo}
}

func (u *majorLessonUsecase) CreateMajorLesson(majorLesson *models.MajorLesson) error {
	return u.repo.Create(majorLesson)
}

func (u *majorLessonUsecase) GetMajorLessonByID(ctx echo.Context, ID database.PID, useCache bool) (*models.MajorLesson, error) {
	return u.repo.GetByID(ctx, ID, useCache)
}

func (u *majorLessonUsecase) UpdateMajorLesson(majorLesson *models.MajorLesson) error {
	return u.repo.Update(majorLesson)
}

func (u *majorLessonUsecase) DeleteMajorLesson(ID database.PID) error {
	return u.repo.Delete(ID)
}

func (u *majorLessonUsecase) GetAllMajorLessons() ([]models.MajorLesson, error) {
	return u.repo.GetAll()
} 