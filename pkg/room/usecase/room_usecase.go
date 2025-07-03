package usecase

import (
	"uni_app/database"
	"uni_app/models"
	"uni_app/pkg/room/repository"
	"uni_app/services/env"

	"github.com/google/uuid"
)

type RoomUsecase interface {
	Create(room *models.Room) error
	Update(room *models.Room) error
	Delete(id database.PID) error
	GetByID(id database.PID) (*models.Room, error)
	List(filters *models.FetchRoomRequest) ([]*models.Room, error)
}

type roomUsecase struct {
	repo   repository.RoomRepository
	config *env.Config
}

func NewRoomUsecase(repo repository.RoomRepository, config *env.Config) RoomUsecase {
	return &roomUsecase{
		repo:   repo,
		config: config,
	}
}

func (u *roomUsecase) Create(room *models.Room) error {
	id, err := database.ParsePID(uuid.New().String())
	if err != nil {
		return err
	}
	room.ID = id
	return u.repo.Create(room)
}

func (u *roomUsecase) Update(room *models.Room) error {
	return u.repo.Update(room)
}

func (u *roomUsecase) Delete(id database.PID) error {
	return u.repo.Delete(id)
}

func (u *roomUsecase) GetByID(id database.PID) (*models.Room, error) {
	return u.repo.GetByID(id)
}

func (u *roomUsecase) List(filters *models.FetchRoomRequest) ([]*models.Room, error) {
	return u.repo.List(filters)
}
