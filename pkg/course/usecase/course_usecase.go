package usecase

import (
	"uni_app/database"
	"uni_app/models"
	"uni_app/pkg/course/repository"
	"uni_app/services/env"

	"github.com/google/uuid"
)

type CourseUsecase interface {
	Create(course *models.Course) error
	Update(course *models.Course) error
	Delete(id database.PID) error
	GetByID(id database.PID) (*models.Course, error)
	List(filters *models.FetchCourseRequest) ([]*models.Course, error)
}

type courseUsecase struct {
	repo   repository.CourseRepository
	config *env.Config
}

func NewCourseUsecase(repo repository.CourseRepository, config *env.Config) CourseUsecase {
	return &courseUsecase{
		repo:   repo,
		config: config,
	}
}

func (u *courseUsecase) Create(course *models.Course) error {
	id, err := database.ParsePID(uuid.New().String())
	if err != nil {
		return err
	}
	course.ID = id
	return u.repo.Create(course)
}

func (u *courseUsecase) Update(course *models.Course) error {
	return u.repo.Update(course)
}

func (u *courseUsecase) Delete(id database.PID) error {
	return u.repo.Delete(id)
}

func (u *courseUsecase) GetByID(id database.PID) (*models.Course, error) {
	return u.repo.GetByID(id)
}

func (u *courseUsecase) List(filters *models.FetchCourseRequest) ([]*models.Course, error) {
	return u.repo.List(filters)
}
