package repository

import (
	"uni_app/database"
	"uni_app/models"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type AddressRepository interface {
	Create(address *models.Address) error
	GetByID(ctx echo.Context, ID database.PID, useCache bool) (*models.Address, error)
	Update(address *models.Address) error
	Delete(ID database.PID) error
	GetAll() ([]models.Address, error)
}

type addressRepository struct {
	db *gorm.DB
}

func NewAddressRepository(db *gorm.DB) AddressRepository {
	return &addressRepository{db}
}

func (r *addressRepository) Create(address *models.Address) error {
	return r.db.Create(address).Error
}

func (r *addressRepository) GetByID(ctx echo.Context, ID database.PID, useCache bool) (*models.Address, error) {
	var address models.Address
	if err := r.db.First(&address, ID).Error; err != nil {
		return nil, err
	}
	return &address, nil
}

func (r *addressRepository) Update(address *models.Address) error {
	return r.db.Save(address).Error
}

func (r *addressRepository) Delete(ID database.PID) error {
	return r.db.Delete(&models.Address{}, ID).Error
}

func (r *addressRepository) GetAll() ([]models.Address, error) {
	var addresses []models.Address
	if err := r.db.Find(&addresses).Error; err != nil {
		return nil, err
	}
	return addresses, nil
} 