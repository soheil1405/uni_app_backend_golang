package usecase

import (
	"errors"
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
	GetLessonsByMajorID(majorID database.PID) ([]models.MajorLesson, error)
	GetLessonsByLessonID(lessonID database.PID) ([]models.MajorLesson, error)
}

type majorLessonUsecase struct {
	repo repository.MajorLessonRepository
}

func NewMajorLessonUsecase(repo repository.MajorLessonRepository) MajorLessonUsecase {
	return &majorLessonUsecase{repo}
}

func (u *majorLessonUsecase) CreateMajorLesson(majorLesson *models.MajorLesson) error {
	// Validate required fields
	if majorLesson.MajorID == 0 {
		return errors.New("major ID is required")
	}
	if majorLesson.LessonID == 0 {
		return errors.New("lesson ID is required")
	}
	// if majorLesson.UnitCount <= 0 {
	// 	return errors.New("unit count must be greater than 0")
	// }

	// Check if the combination already exists
	existingLessons, err := u.repo.GetAll()
	if err != nil {
		return err
	}

	for _, existing := range existingLessons {
		if existing.MajorID == majorLesson.MajorID && existing.LessonID == majorLesson.LessonID {
			return errors.New("this lesson already exists for this major")
		}
	}

	return u.repo.Create(majorLesson)
}

func (u *majorLessonUsecase) GetMajorLessonByID(ctx echo.Context, ID database.PID, useCache bool) (*models.MajorLesson, error) {
	return u.repo.GetByID(ctx, ID, useCache)
}

func (u *majorLessonUsecase) UpdateMajorLesson(majorLesson *models.MajorLesson) error {
	// Validate required fields
	if majorLesson.MajorID == 0 {
		return errors.New("major ID is required")
	}
	if majorLesson.LessonID == 0 {
		return errors.New("lesson ID is required")
	}
	// if majorLesson.UnitCount <= 0 {
	// 	return errors.New("unit count must be greater than 0")
	// }

	// Check if the combination already exists (excluding current record)
	existingLessons, err := u.repo.GetAll()
	if err != nil {
		return err
	}

	for _, existing := range existingLessons {
		if existing.ID != majorLesson.ID &&
			existing.MajorID == majorLesson.MajorID &&
			existing.LessonID == majorLesson.LessonID {
			return errors.New("this lesson already exists for this major")
		}
	}

	return u.repo.Update(majorLesson)
}

func (u *majorLessonUsecase) DeleteMajorLesson(ID database.PID) error {
	// Check if there are any related records before deletion
	// This could be expanded based on your requirements
	return u.repo.Delete(ID)
}

func (u *majorLessonUsecase) GetAllMajorLessons() ([]models.MajorLesson, error) {
	return u.repo.GetAll()
}

func (u *majorLessonUsecase) GetLessonsByMajorID(majorID database.PID) ([]models.MajorLesson, error) {
	allLessons, err := u.repo.GetAll()
	if err != nil {
		return nil, err
	}

	var filteredLessons []models.MajorLesson
	for _, lesson := range allLessons {
		if lesson.MajorID == majorID {
			filteredLessons = append(filteredLessons, lesson)
		}
	}

	return filteredLessons, nil
}

func (u *majorLessonUsecase) GetLessonsByLessonID(lessonID database.PID) ([]models.MajorLesson, error) {
	allLessons, err := u.repo.GetAll()
	if err != nil {
		return nil, err
	}

	var filteredLessons []models.MajorLesson
	for _, lesson := range allLessons {
		if lesson.LessonID == lessonID {
			filteredLessons = append(filteredLessons, lesson)
		}
	}

	return filteredLessons, nil
}
