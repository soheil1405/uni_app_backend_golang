package usecase

import (
	"errors"
	"uni_app/database"
	"uni_app/models"
	repository "uni_app/pkg/address/repository"
	"uni_app/utils/helpers"

	"github.com/labstack/echo/v4"
)

type AddressUsecase interface {
	CreateAddress(address *models.Address) error
	GetAddressByID(ctx echo.Context, ID database.PID, useCache bool) (*models.Address, error)
	UpdateAddress(address *models.Address) error
	DeleteAddress(ID database.PID) error
	GetAllAddresses(ctx echo.Context, request models.FetchAddressRequest) ([]models.Address, *helpers.PaginateTemplate, error)
}

type addressUsecase struct {
	repo repository.AddressRepository
}

func NewAddressUsecase(repo repository.AddressRepository) AddressUsecase {
	return &addressUsecase{repo}
}

func (u *addressUsecase) CreateAddress(address *models.Address) error {
	if address.CityID == 0 {
		return errors.New("city ID is required")
	}
	return u.repo.Create(address)
}

func (u *addressUsecase) GetAddressByID(ctx echo.Context, ID database.PID, useCache bool) (*models.Address, error) {
	return u.repo.GetByID(ctx, ID, useCache)
}

func (u *addressUsecase) UpdateAddress(address *models.Address) error {
	if address.CityID == 0 {
		return errors.New("city ID is required")
	}

	return u.repo.Update(address)
}

func (u *addressUsecase) DeleteAddress(ID database.PID) error {
	// Check if this is the user's only address
	address, err := u.repo.GetByID(nil, ID, false)
	if err != nil || address == nil {
		return errors.New("address not found")
	}

	return u.repo.Delete(ID)
}

func (u *addressUsecase) GetAllAddresses(ctx echo.Context, request models.FetchAddressRequest) ([]models.Address, *helpers.PaginateTemplate, error) {
	return u.repo.GetAll(ctx, request)
}
