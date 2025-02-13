package repositories

import (
	"uni_app/database"
	"uni_app/models"

	"gorm.io/gorm"
)

type ChartRepository interface {
	Create(chart *models.MajorsChart) error
	GetByID(ID database.PID) (*models.MajorsChart, error)
	Update(chart *models.MajorsChart) error
	Delete(ID database.PID) error
	GetAll() ([]models.MajorsChart, error)
}

type chartRepository struct {
	db *gorm.DB
}

func NewChartRepository(db *gorm.DB) ChartRepository {
	return &chartRepository{db}
}

func (r *chartRepository) Create(chart *models.MajorsChart) error {
	return r.db.Create(chart).Error
}

func (r *chartRepository) GetByID(ID database.PID) (*models.MajorsChart, error) {
	var chart models.MajorsChart
	if err := r.db.First(&chart, ID).Error; err != nil {
		return nil, err
	}
	return &chart, nil
}

func (r *chartRepository) Update(chart *models.MajorsChart) error {
	return r.db.Save(chart).Error
}

func (r *chartRepository) Delete(ID database.PID) error {
	return r.db.Delete(&models.MajorsChart{}, ID).Error
}

func (r *chartRepository) GetAll() ([]models.MajorsChart, error) {
	var charts []models.MajorsChart
	if err := r.db.Find(&charts).Error; err != nil {
		return nil, err
	}
	return charts, nil
}
