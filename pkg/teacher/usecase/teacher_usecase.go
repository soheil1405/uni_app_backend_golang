package usecase

import (
	"uni_app/database"
	"uni_app/models"
	"uni_app/pkg/teacher/repository"
	"uni_app/services/env"

	"github.com/google/uuid"
)

type TeacherUsecase interface {
	Create(teacher *models.Teacher) error
	Update(teacher *models.Teacher) error
	Delete(id database.PID) error
	GetByID(id database.PID) (*models.Teacher, error)
	List(filters *models.FetchTeacherRequest) ([]*models.Teacher, error)
}

type teacherUsecase struct {
	repo   repository.TeacherRepository
	config *env.Config
}

func NewTeacherUsecase(repo repository.TeacherRepository, config *env.Config) TeacherUsecase {
	return &teacherUsecase{
		repo:   repo,
		config: config,
	}
}

func (u *teacherUsecase) Create(teacher *models.Teacher) error {
	id, err := database.ParsePID(uuid.New().String())
	if err != nil {
		return err
	}
	teacher.ID = id
	return u.repo.Create(teacher)
}

func (u *teacherUsecase) Update(teacher *models.Teacher) error {
	return u.repo.Update(teacher)
}

func (u *teacherUsecase) Delete(id database.PID) error {
	return u.repo.Delete(id)
}

func (u *teacherUsecase) GetByID(id database.PID) (*models.Teacher, error) {
	return u.repo.GetByID(id)
}

func (u *teacherUsecase) List(filters *models.FetchTeacherRequest) ([]*models.Teacher, error) {
	return u.repo.List(filters)
}
