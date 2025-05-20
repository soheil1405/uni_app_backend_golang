package usecase

import (
	"errors"
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
	GetAddressesByUserID(userID database.PID) ([]models.Address, error)
	GetAddressesByCityID(cityID database.PID) ([]models.Address, error)
	GetAddressesByProvinceID(provinceID database.PID) ([]models.Address, error)
}

type addressUsecase struct {
	repo repository.AddressRepository
}

func NewAddressUsecase(repo repository.AddressRepository) AddressUsecase {
	return &addressUsecase{repo}
}

func (u *addressUsecase) CreateAddress(address *models.Address) error {
	// Validate required fields
	if address.UserID == 0 {
		return errors.New("user ID is required")
	}
	if address.CityID == 0 {
		return errors.New("city ID is required")
	}
	if address.ProvinceID == 0 {
		return errors.New("province ID is required")
	}
	if address.Address == "" {
		return errors.New("address is required")
	}
	if address.PostalCode == "" {
		return errors.New("postal code is required")
	}

	// Validate postal code format (assuming Iranian postal code format)
	if len(address.PostalCode) != 10 {
		return errors.New("postal code must be 10 digits")
	}

	// Check if this is the user's first address
	existingAddresses, err := u.repo.GetAll()
	if err != nil {
		return err
	}

	isFirstAddress := true
	for _, existing := range existingAddresses {
		if existing.UserID == address.UserID {
			isFirstAddress = false
			break
		}
	}

	// If it's the first address, set it as default
	if isFirstAddress {
		address.IsDefault = true
	}

	return u.repo.Create(address)
}

func (u *addressUsecase) GetAddressByID(ctx echo.Context, ID database.PID, useCache bool) (*models.Address, error) {
	return u.repo.GetByID(ctx, ID, useCache)
}

func (u *addressUsecase) UpdateAddress(address *models.Address) error {
	// Validate required fields
	if address.UserID == 0 {
		return errors.New("user ID is required")
	}
	if address.CityID == 0 {
		return errors.New("city ID is required")
	}
	if address.ProvinceID == 0 {
		return errors.New("province ID is required")
	}
	if address.Address == "" {
		return errors.New("address is required")
	}
	if address.PostalCode == "" {
		return errors.New("postal code is required")
	}

	// Validate postal code format
	if len(address.PostalCode) != 10 {
		return errors.New("postal code must be 10 digits")
	}

	// If setting as default, update other addresses
	if address.IsDefault {
		existingAddresses, err := u.repo.GetAll()
		if err != nil {
			return err
		}

		for _, existing := range existingAddresses {
			if existing.UserID == address.UserID && existing.ID != address.ID {
				existing.IsDefault = false
				if err := u.repo.Update(&existing); err != nil {
					return err
				}
			}
		}
	}

	return u.repo.Update(address)
}

func (u *addressUsecase) DeleteAddress(ID database.PID) error {
	// Check if this is the user's only address
	address, err := u.repo.GetByID(nil, ID, false)
	if err != nil {
		return err
	}

	existingAddresses, err := u.repo.GetAll()
	if err != nil {
		return err
	}

	addressCount := 0
	for _, existing := range existingAddresses {
		if existing.UserID == address.UserID {
			addressCount++
		}
	}

	if addressCount <= 1 {
		return errors.New("cannot delete the only address")
	}

	return u.repo.Delete(ID)
}

func (u *addressUsecase) GetAllAddresses() ([]models.Address, error) {
	return u.repo.GetAll()
}

func (u *addressUsecase) GetAddressesByUserID(userID database.PID) ([]models.Address, error) {
	allAddresses, err := u.repo.GetAll()
	if err != nil {
		return nil, err
	}

	var filteredAddresses []models.Address
	for _, address := range allAddresses {
		if address.UserID == userID {
			filteredAddresses = append(filteredAddresses, address)
		}
	}

	return filteredAddresses, nil
}

func (u *addressUsecase) GetAddressesByCityID(cityID database.PID) ([]models.Address, error) {
	allAddresses, err := u.repo.GetAll()
	if err != nil {
		return nil, err
	}

	var filteredAddresses []models.Address
	for _, address := range allAddresses {
		if address.CityID == cityID {
			filteredAddresses = append(filteredAddresses, address)
		}
	}

	return filteredAddresses, nil
}

func (u *addressUsecase) GetAddressesByProvinceID(provinceID database.PID) ([]models.Address, error) {
	allAddresses, err := u.repo.GetAll()
	if err != nil {
		return nil, err
	}

	var filteredAddresses []models.Address
	for _, address := range allAddresses {
		if address.ProvinceID == provinceID {
			filteredAddresses = append(filteredAddresses, address)
		}
	}

	return filteredAddresses, nil
} 