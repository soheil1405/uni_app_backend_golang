package usecases

import (
	"uni_app/models"
	repositories "uni_app/pkg/major/repository"
)

type MajorUsecase interface {
	CreateMajor(major *models.Major) error
	GetMajorByID(id uint) (*models.Major, error)
	UpdateMajor(major *models.Major) error
	DeleteMajor(id uint) error
	GetAllMajors() ([]models.Major, error)
}

type majorUsecase struct {
	repo repositories.MajorRepository
}

func NewMajorUsecase(repo repositories.MajorRepository) MajorUsecase {
	return &majorUsecase{repo}
}

func (u *majorUsecase) CreateMajor(major *models.Major) error {
	return u.repo.Create(major)
}

func (u *majorUsecase) GetMajorByID(id uint) (*models.Major, error) {
	return u.repo.GetByID(id)
}

func (u *majorUsecase) UpdateMajor(major *models.Major) error {
	return u.repo.Update(major)
}

func (u *majorUsecase) DeleteMajor(id uint) error {
	return u.repo.Delete(id)
}

func (u *majorUsecase) GetAllMajors() ([]models.Major, error) {
	return u.repo.GetAll()
}
