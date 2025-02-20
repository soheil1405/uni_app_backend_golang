package usecases

import (
	"uni_app/database"
	"uni_app/models"
	repositories "uni_app/pkg/daneshkadeh/repository"
)

type FacultyUsecase interface {
	CreateDaneshKadeh(faculty *models.DaneshKadeh) error
	GetDaneshKadehByID(ID database.PID) (*models.DaneshKadeh, error)
	UpdateDaneshKadeh(faculty *models.DaneshKadeh) error
	DeleteDaneshKadeh(ID database.PID) error
	GetAllDaneshKadeha() (*models.DaneshKadeha, error)
}

type facultyUsecase struct {
	repo repositories.FacultyRepository
}

func NewDaneshKadehUsecase(repo repositories.FacultyRepository) FacultyUsecase {
	return &facultyUsecase{repo}
}

func (u *facultyUsecase) CreateDaneshKadeh(faculty *models.DaneshKadeh) error {
	return u.repo.Create(faculty)
}

func (u *facultyUsecase) GetDaneshKadehByID(ID database.PID) (*models.DaneshKadeh, error) {
	return u.repo.GetByID(ID)
}

func (u *facultyUsecase) UpdateDaneshKadeh(faculty *models.DaneshKadeh) error {
	return u.repo.Update(faculty)
}

func (u *facultyUsecase) DeleteDaneshKadeh(ID database.PID) error {
	return u.repo.Delete(ID)
}

func (u *facultyUsecase) GetAllDaneshKadeha() (*models.DaneshKadeha, error) {
	return u.repo.GetAll()
}
