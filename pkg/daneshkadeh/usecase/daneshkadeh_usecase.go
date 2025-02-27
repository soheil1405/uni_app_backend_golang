package usecases

import (
	"net/http"
	"uni_app/database"
	"uni_app/models"
	repositories "uni_app/pkg/daneshkadeh/repository"
	"uni_app/utils/helpers"
	"uni_app/utils/templates"

	"github.com/labstack/echo/v4"
)

type FacultyUsecase interface {
	CreateDaneshKadeh(faculty *models.DaneshKadeh) error
	GetDaneshKadehByID(ctx echo.Context, ID database.PID, useCache bool) (map[string]interface{}, helpers.MyError)
	UpdateDaneshKadeh(faculty *models.DaneshKadeh) error
	DeleteDaneshKadeh(ID database.PID) error
	GetAllDaneshKadeha(ctx echo.Context, request models.FetchRequest) (map[string]interface{}, *templates.PaginateTemplate, helpers.MyError)
}

type facultyUsecase struct {
	repo repositories.FacultyRepository
}

func NewDaneshKadehUsecase(repo repositories.FacultyRepository) FacultyUsecase {
	return &facultyUsecase{repo}
}

func (u *facultyUsecase) CreateDaneshKadeh(faculty *models.DaneshKadeh) error {
	return u.repo.Create(faculty)
}

func (u *facultyUsecase) GetDaneshKadehByID(ctx echo.Context, ID database.PID, useCache bool) (map[string]interface{}, helpers.MyError) {
	var (
		response = make(map[string]interface{})
		MyErr    helpers.MyError
	)
	MyErr.Default()
	daneshKadeh, err := u.repo.GetByID(ctx, ID, useCache)
	if err != nil {
		MyErr.SetError(err, http.StatusBadRequest)
	}
	response["daneshkadeh"] = daneshKadeh
	return response, MyErr
}

func (u *facultyUsecase) UpdateDaneshKadeh(faculty *models.DaneshKadeh) error {
	return u.repo.Update(faculty)
}

func (u *facultyUsecase) DeleteDaneshKadeh(ID database.PID) error {
	return u.repo.Delete(ID)
}

func (u *facultyUsecase) GetAllDaneshKadeha(ctx echo.Context, request models.FetchRequest) (map[string]interface{}, *templates.PaginateTemplate, helpers.MyError) {
	var (
		myErr    helpers.MyError
		response = make(map[string]interface{})
	)
	daneshkadeha, meta, err := u.repo.GetAll(ctx, request)
	if err != nil {
		myErr.SetError(err, http.StatusBadRequest)
	}
	response["daneshkadeha"] = daneshkadeha
	myErr.Default()
	return response, meta, myErr
}
