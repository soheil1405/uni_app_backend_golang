package usecase

import (
	"errors"
	"uni_app/database"
	"uni_app/models"
	repository "uni_app/pkg/major_lesson/repository"
	"uni_app/utils/helpers"

	"github.com/labstack/echo/v4"
)

type MajorLessonUsecase interface {
	CreateMajorLesson(majorLesson *models.MajorLesson) error
	GetMajorLessonByID(ctx echo.Context, ID database.PID, useCache bool) (*models.MajorLesson, error)
	UpdateMajorLesson(majorLesson *models.MajorLesson) error
	DeleteMajorLesson(ID database.PID) error
	GetAllMajorLessons(ctx echo.Context, request models.FetchMajorLessonRequest) ([]models.MajorLesson, *helpers.PaginateTemplate, error)
	GetLessonsByMajorID(majorID database.PID) ([]models.MajorLesson, error)
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
	request := models.FetchMajorLessonRequest{
		MajorID:  majorLesson.MajorID,
		LessonID: majorLesson.LessonID,
	}
	existingLessons, _, err := u.repo.GetAll(nil, request)
	if err != nil {
		return err
	}

	if len(existingLessons) > 0 {
		return errors.New("this lesson already exists for this major")
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
	request := models.FetchMajorLessonRequest{
		MajorID:  majorLesson.MajorID,
		LessonID: majorLesson.LessonID,
	}
	existingLessons, _, err := u.repo.GetAll(nil, request)
	if err != nil {
		return err
	}

	for _, existing := range existingLessons {
		if existing.ID != majorLesson.ID {
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

func (u *majorLessonUsecase) GetAllMajorLessons(ctx echo.Context, request models.FetchMajorLessonRequest) ([]models.MajorLesson, *helpers.PaginateTemplate, error) {
	return u.repo.GetAll(ctx, request)
}

func (u *majorLessonUsecase) GetLessonsByMajorID(majorID database.PID) ([]models.MajorLesson, error) {
	request := models.FetchMajorLessonRequest{
		MajorID: majorID,
		FetchRequest: models.FetchRequest{
			Limit:  1,
			Offset: 0,
		},
	}
	allLessons, _, err := u.repo.GetAll(nil, request)
	if err != nil {
		return nil, err
	}

	return allLessons, nil
}
