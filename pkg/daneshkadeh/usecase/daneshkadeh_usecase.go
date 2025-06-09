package usecase

import (
	"uni_app/database"
	"uni_app/models"
	repository "uni_app/pkg/daneshkadeh/repository"
	"uni_app/utils/templates"

	"github.com/labstack/echo/v4"
)

type DaneshKadehUsecase interface {
	CreateDaneshKadeh(daneshKadeh *models.DaneshKadeh) error
	GetDaneshKadehByID(ctx echo.Context, ID database.PID, useCache bool) (*models.DaneshKadeh, error)
	UpdateDaneshKadeh(daneshKadeh *models.DaneshKadeh) error
	DeleteDaneshKadeh(ID database.PID) error
	GetAllDaneshKadehs(ctx echo.Context, request models.FetchDaneshKadehRequest) (models.DaneshKadeha, *templates.PaginateTemplate, error)
}

type daneshKadehUsecase struct {
	repo repository.FacultyRepository
}

func NewDaneshKadehUsecase(repo repository.FacultyRepository) DaneshKadehUsecase {
	return &daneshKadehUsecase{repo}
}

func (u *daneshKadehUsecase) CreateDaneshKadeh(daneshKadeh *models.DaneshKadeh) error {
	return u.repo.Create(daneshKadeh)
}

func (u *daneshKadehUsecase) GetDaneshKadehByID(ctx echo.Context, ID database.PID, useCache bool) (*models.DaneshKadeh, error) {
	return u.repo.GetByID(ctx, ID, useCache)
}

func (u *daneshKadehUsecase) UpdateDaneshKadeh(daneshKadeh *models.DaneshKadeh) error {
	return u.repo.Update(daneshKadeh)
}

func (u *daneshKadehUsecase) DeleteDaneshKadeh(ID database.PID) error {
	return u.repo.Delete(ID)
}

func (u *daneshKadehUsecase) GetAllDaneshKadehs(ctx echo.Context, request models.FetchDaneshKadehRequest) (models.DaneshKadeha, *templates.PaginateTemplate, error) {
	if request.Limit == 0 {
		request.Limit = 10
	}

	return u.repo.GetAll(ctx, request)
}
