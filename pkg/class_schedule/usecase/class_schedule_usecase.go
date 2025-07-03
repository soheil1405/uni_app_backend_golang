package usecase

import (
	"uni_app/database"
	"uni_app/models"
	"uni_app/pkg/class_schedule/repository"
	"uni_app/services/env"

	"github.com/google/uuid"
)

type ClassScheduleUsecase interface {
	Create(classSchedule *models.ClassSchedule) error
	Update(classSchedule *models.ClassSchedule) error
	Delete(id database.PID) error
	GetByID(id database.PID) (*models.ClassSchedule, error)
	List(filters *models.FetchClassScheduleRequest) ([]*models.ClassSchedule, error)
}

type classScheduleUsecase struct {
	repo   repository.ClassScheduleRepository
	config *env.Config
}

func NewClassScheduleUsecase(repo repository.ClassScheduleRepository, config *env.Config) ClassScheduleUsecase {
	return &classScheduleUsecase{
		repo:   repo,
		config: config,
	}
}

func (u *classScheduleUsecase) Create(classSchedule *models.ClassSchedule) error {
	id, err := database.ParsePID(uuid.New().String())
	if err != nil {
		return err
	}
	classSchedule.ID = id
	return u.repo.Create(classSchedule)
}

func (u *classScheduleUsecase) Update(classSchedule *models.ClassSchedule) error {
	return u.repo.Update(classSchedule)
}

func (u *classScheduleUsecase) Delete(id database.PID) error {
	return u.repo.Delete(id)
}

func (u *classScheduleUsecase) GetByID(id database.PID) (*models.ClassSchedule, error) {
	return u.repo.GetByID(id)
}

func (u *classScheduleUsecase) List(filters *models.FetchClassScheduleRequest) ([]*models.ClassSchedule, error) {
	return u.repo.List(filters)
}
