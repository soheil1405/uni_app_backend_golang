package usecases

import (
	"net/http"
	"uni_app/database"
	"uni_app/models"
	repositories "uni_app/pkg/uni/repository"
	"uni_app/utils/helpers"
	"uni_app/utils/templates"

	"github.com/labstack/echo/v4"
)

type UniUsecase interface {
	CreateUni(uni *models.Uni) error
	GetUniByID(ctx echo.Context, ID database.PID, useCache bool) (*models.Uni, error)
	UpdateUni(uni *models.Uni) error
	DeleteUni(ID database.PID) error
	GetAllUnis(ctx echo.Context, request models.FetchRequest) (map[string]interface{}, *templates.PaginateTemplate, helpers.MyError)
}

type uniUsecase struct {
	repo repositories.UniRepository
}

func NewUniUsecase(repo repositories.UniRepository) UniUsecase {
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

func (u *uniUsecase) GetAllUnis(ctx echo.Context, request models.FetchRequest) (map[string]interface{}, *templates.PaginateTemplate, helpers.MyError) {
	var (
		myErr    helpers.MyError
		response = make(map[string]interface{})
	)
	unis, meta, err := u.repo.GetAll(ctx, request)
	if err != nil {
		myErr.SetError(err, http.StatusBadRequest)
	}
	response["unis"] = unis
	myErr.Default()
	return response, meta, myErr
}
