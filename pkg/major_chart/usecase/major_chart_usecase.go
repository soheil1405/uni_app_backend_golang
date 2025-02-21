package usecases

import (
	"uni_app/database"
	"uni_app/models"
	repositories "uni_app/pkg/major_chart/repository"

	"github.com/labstack/echo/v4"
)

type ChartUsecase interface {
	CreateChart(chart *models.MajorsChart) error
	GetChartByID(ctx echo.Context, ID database.PID, useCache bool) (*models.MajorsChart, error)
	UpdateChart(chart *models.MajorsChart) error
	DeleteChart(ID database.PID) error
	GetAllCharts() ([]models.MajorsChart, error)
}

type chartUsecase struct {
	repo repositories.ChartRepository
}

func NewChartUsecase(repo repositories.ChartRepository) ChartUsecase {
	return &chartUsecase{repo}
}

func (u *chartUsecase) CreateChart(chart *models.MajorsChart) error {
	return u.repo.Create(chart)
}

func (u *chartUsecase) GetChartByID(ctx echo.Context, ID database.PID, useCache bool) (*models.MajorsChart, error) {
	return u.repo.GetByID(ctx, ID, useCache)
}

func (u *chartUsecase) UpdateChart(chart *models.MajorsChart) error {
	return u.repo.Update(chart)
}

func (u *chartUsecase) DeleteChart(ID database.PID) error {
	return u.repo.Delete(ID)
}

func (u *chartUsecase) GetAllCharts() ([]models.MajorsChart, error) {
	return u.repo.GetAll()
}
