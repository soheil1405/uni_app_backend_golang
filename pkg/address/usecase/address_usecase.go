package usecase

import (
	"uni_app/database"
	"uni_app/models"
	repository "uni_app/pkg/address/repository"

	"github.com/labstack/echo/v4"
)

type AddressUsecase interface {
	CreateAddress(address *models.Address) error
	GetAddressByID(ctx echo.Context, ID database.PID, useCache bool) (*models.Address, error)
	UpdateAddress(address *models.Address) error
	DeleteAddress(ID database.PID) error
	GetAllAddresses() ([]models.Address, error)
}

type addressUsecase struct {
	repo repository.AddressRepository
}

func NewAddressUsecase(repo repository.AddressRepository) AddressUsecase {
	return &addressUsecase{repo}
}

func (u *addressUsecase) CreateAddress(address *models.Address) error {
	return u.repo.Create(address)
}

func (u *addressUsecase) GetAddressByID(ctx echo.Context, ID database.PID, useCache bool) (*models.Address, error) {
	return u.repo.GetByID(ctx, ID, useCache)
}

func (u *addressUsecase) UpdateAddress(address *models.Address) error {
	return u.repo.Update(address)
}

func (u *addressUsecase) DeleteAddress(ID database.PID) error {
	return u.repo.Delete(ID)
}

func (u *addressUsecase) GetAllAddresses() ([]models.Address, error) {
	return u.repo.GetAll()
} 