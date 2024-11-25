package usecases

import (
	"uni_app/models"
	repositories "uni_app/pkg/uni/repository"
)

type UniUsecase interface {
	CreateUni(uni *models.Uni) error
	GetUniByID(id uint) (*models.Uni, error)
	UpdateUni(uni *models.Uni) error
	DeleteUni(id uint) error
	GetAllUnis() ([]models.Uni, error)
}

type uniUsecase struct {
	repo repositories.UniRepository
}

func NewUniUsecase(repo repositories.UniRepository) UniUsecase {
	return &uniUsecase{repo}
}

func (u *uniUsecase) CreateUni(uni *models.Uni) error {
	return u.repo.Create(uni)
}

func (u *uniUsecase) GetUniByID(id uint) (*models.Uni, error) {
	return u.repo.GetByID(id)
}

func (u *uniUsecase) UpdateUni(uni *models.Uni) error {
	return u.repo.Update(uni)
}

func (u *uniUsecase) DeleteUni(id uint) error {
	return u.repo.Delete(id)
}

func (u *uniUsecase) GetAllUnis() ([]models.Uni, error) {
	return u.repo.GetAll()
}
