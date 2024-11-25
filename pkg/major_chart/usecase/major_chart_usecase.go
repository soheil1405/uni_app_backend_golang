package usecases

import (
	"uni_app/models"
	repositories "uni_app/pkg/major_chart/repository"
)

type ChartUsecase interface {
	CreateChart(chart *models.MajorsChart) error
	GetChartByID(id uint) (*models.MajorsChart, error)
	UpdateChart(chart *models.MajorsChart) error
	DeleteChart(id uint) error
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

func (u *chartUsecase) GetChartByID(id uint) (*models.MajorsChart, error) {
	return u.repo.GetByID(id)
}

func (u *chartUsecase) UpdateChart(chart *models.MajorsChart) error {
	return u.repo.Update(chart)
}

func (u *chartUsecase) DeleteChart(id uint) error {
	return u.repo.Delete(id)
}

func (u *chartUsecase) GetAllCharts() ([]models.MajorsChart, error) {
	return u.repo.GetAll()
}
