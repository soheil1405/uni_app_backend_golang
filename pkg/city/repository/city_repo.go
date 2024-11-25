package repositories

import (
	"uni_app/models"

	"gorm.io/gorm"
)

type CityRepository interface {
	Create(city *models.City) error
	GetByID(id uint) (*models.City, error)
	Update(city *models.City) error
	Delete(id uint) error
	GetAll() ([]models.City, error)
}

type cityRepository struct {
	db *gorm.DB
}

func NewCityRepository(db *gorm.DB) CityRepository {
	return &cityRepository{db}
}

func (r *cityRepository) Create(city *models.City) error {
	return r.db.Create(city).Error
}

func (r *cityRepository) GetByID(id uint) (*models.City, error) {
	var city models.City
	if err := r.db.First(&city, id).Error; err != nil {
		return nil, err
	}
	return &city, nil
}

func (r *cityRepository) Update(city *models.City) error {
	return r.db.Save(city).Error
}

func (r *cityRepository) Delete(id uint) error {
	return r.db.Delete(&models.City{}, id).Error
}

func (r *cityRepository) GetAll() ([]models.City, error) {
	var cities []models.City
	if err := r.db.Find(&cities).Error; err != nil {
		return nil, err
	}
	return cities, nil
}
