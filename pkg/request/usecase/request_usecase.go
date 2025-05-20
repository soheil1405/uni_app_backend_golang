package usecase

import (
	"uni_app/database"
	"uni_app/models"
	repository "uni_app/pkg/request/repository"

	"github.com/labstack/echo/v4"
)

type RequestUsecase interface {
	CreateRequest(request *models.FetchRequest) error
	GetRequestByID(ctx echo.Context, ID database.PID, useCache bool) (*models.FetchRequest, error)
	UpdateRequest(request *models.FetchRequest) error
	DeleteRequest(ID database.PID) error
	GetAllRequests() ([]models.FetchRequest, error)
}

type requestUsecase struct {
	repo repository.RequestRepository
}

func NewRequestUsecase(repo repository.RequestRepository) RequestUsecase {
	return &requestUsecase{repo}
}

func (u *requestUsecase) CreateRequest(request *models.FetchRequest) error {
	return u.repo.Create(request)
}

func (u *requestUsecase) GetRequestByID(ctx echo.Context, ID database.PID, useCache bool) (*models.FetchRequest, error) {
	return u.repo.GetByID(ctx, ID, useCache)
}

func (u *requestUsecase) UpdateRequest(request *models.FetchRequest) error {
	return u.repo.Update(request)
}

func (u *requestUsecase) DeleteRequest(ID database.PID) error {
	return u.repo.Delete(ID)
}

func (u *requestUsecase) GetAllRequests() ([]models.FetchRequest, error) {
	return u.repo.GetAll()
} 