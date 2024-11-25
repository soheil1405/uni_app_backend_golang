package repositories

import (
	"uni_app/models"

	"gorm.io/gorm"
)

type PlaceRepository interface {
	Create(place *models.Place) error
	GetByID(id uint) (*models.Place, error)
	Update(place *models.Place) error
	Delete(id uint) error
	GetAll() ([]models.Place, error)
}

type placeRepository struct {
	db *gorm.DB
}

func NewPlaceRepository(db *gorm.DB) PlaceRepository {
	return &placeRepository{db}
}

func (r *placeRepository) Create(place *models.Place) error {
	return r.db.Create(place).Error
}

func (r *placeRepository) GetByID(id uint) (*models.Place, error) {
	var place models.Place
	if err := r.db.First(&place, id).Error; err != nil {
		return nil, err
	}
	return &place, nil
}

func (r *placeRepository) Update(place *models.Place) error {
	return r.db.Save(place).Error
}

func (r *placeRepository) Delete(id uint) error {
	return r.db.Delete(&models.Place{}, id).Error
}

func (r *placeRepository) GetAll() ([]models.Place, error) {
	var places []models.Place
	if err := r.db.Find(&places).Error; err != nil {
		return nil, err
	}
	return places, nil
}
