package usecase

import (
	"uni_app/database"
	"uni_app/models"
	"uni_app/pkg/course_instance/repository"
	"uni_app/services/env"

	"github.com/google/uuid"
)

type CourseInstanceUsecase interface {
	Create(courseInstance *models.CourseInstance) error
	Update(courseInstance *models.CourseInstance) error
	Delete(id database.PID) error
	GetByID(id database.PID) (*models.CourseInstance, error)
	List(filters *models.FetchCourseInstanceRequest) ([]*models.CourseInstance, error)
}

type courseInstanceUsecase struct {
	repo   repository.CourseInstanceRepository
	config *env.Config
}

func NewCourseInstanceUsecase(repo repository.CourseInstanceRepository, config *env.Config) CourseInstanceUsecase {
	return &courseInstanceUsecase{
		repo:   repo,
		config: config,
	}
}

func (u *courseInstanceUsecase) Create(courseInstance *models.CourseInstance) error {
	id, err := database.ParsePID(uuid.New().String())
	if err != nil {
		return err
	}
	courseInstance.ID = id
	return u.repo.Create(courseInstance)
}

func (u *courseInstanceUsecase) Update(courseInstance *models.CourseInstance) error {
	return u.repo.Update(courseInstance)
}

func (u *courseInstanceUsecase) Delete(id database.PID) error {
	return u.repo.Delete(id)
}

func (u *courseInstanceUsecase) GetByID(id database.PID) (*models.CourseInstance, error) {
	return u.repo.GetByID(id)
}

func (u *courseInstanceUsecase) List(filters *models.FetchCourseInstanceRequest) ([]*models.CourseInstance, error) {
	return u.repo.List(filters)
}
