package repositories

import (
	"uni_app/models"

	"gorm.io/gorm"
)

type PlaceTypeRepository interface {
	Create(placeType *models.PlaceType) error
	GetByID(id uint) (*models.PlaceType, error)
	Update(placeType *models.PlaceType) error
	Delete(id uint) error
	GetAll() ([]models.PlaceType, error)
}

type placeTypeRepository struct {
	db *gorm.DB
}

func NewPlaceTypeRepository(db *gorm.DB) PlaceTypeRepository {
	return &placeTypeRepository{db}
}

func (r *placeTypeRepository) Create(placeType *models.PlaceType) error {
	return r.db.Create(placeType).Error
}

func (r *placeTypeRepository) GetByID(id uint) (*models.PlaceType, error) {
	var placeType models.PlaceType
	if err := r.db.First(&placeType, id).Error; err != nil {
		return nil, err
	}
	return &placeType, nil
}

func (r *placeTypeRepository) Update(placeType *models.PlaceType) error {
	return r.db.Save(placeType).Error
}

func (r *placeTypeRepository) Delete(id uint) error {
	return r.db.Delete(&models.PlaceType{}, id).Error
}

func (r *placeTypeRepository) GetAll() ([]models.PlaceType, error) {
	var placeTypes []models.PlaceType
	if err := r.db.Find(&placeTypes).Error; err != nil {
		return nil, err
	}
	return placeTypes, nil
}
