package usecase

import (
	"uni_app/database"
	"uni_app/models"
	repository "uni_app/pkg/uni/repository"
	"uni_app/utils/helpers"

	"github.com/labstack/echo/v4"
)

type UniUsecase interface {
	CreateUni(uni *models.Uni) error
	GetUniByID(ctx echo.Context, ID database.PID, useCache bool) (*models.Uni, error)
	UpdateUni(uni *models.Uni) error
	DeleteUni(ID database.PID) error
	GetAllUnis(ctx echo.Context, request models.FetchUniRequest) ([]models.Uni, *helpers.PaginateTemplate, error)
}

type uniUsecase struct {
	repo repository.UniRepository
}

func NewUniUsecase(repo repository.UniRepository) UniUsecase {
	return &uniUsecase{repo}
}

func (u *uniUsecase) CreateUni(uni *models.Uni) error {
	return u.repo.Create(uni)
}

func (u *uniUsecase) GetUniByID(ctx echo.Context, ID database.PID, useCache bool) (*models.Uni, error) {
	return u.repo.GetByID(ctx, ID, useCache)
}

func (u *uniUsecase) UpdateUni(uni *models.Uni) error {
	return u.repo.Update(uni)
}

func (u *uniUsecase) DeleteUni(ID database.PID) error {
	return u.repo.Delete(ID)
}

func (u *uniUsecase) GetAllUnis(ctx echo.Context, request models.FetchUniRequest) ([]models.Uni, *helpers.PaginateTemplate, error) {
	return u.repo.GetAll(ctx, request)
}
