package usecases

import (
	"uni_app/database"
	"uni_app/models"
	repositories "uni_app/pkg/student/repository"
)

type StudentUsecase interface {
	CreateStudent(student *models.Student) error
	GetStudentByID(ID database.PID) (*models.Student, error)
	UpdateStudent(student *models.Student) error
	DeleteStudent(ID database.PID) error
	GetAllStudents() ([]models.Student, error)
}

type userUsecase struct {
	repo repositories.StudentRepository
}

func NewUserUsecase(repo repositories.StudentRepository) StudentUsecase {
	return &userUsecase{repo}
}

func (u *userUsecase) CreateStudent(student *models.Student) error {
	return u.repo.Create(student)
}

func (u *userUsecase) GetStudentByID(ID database.PID) (*models.Student, error) {
	return u.repo.GetByID(ID)
}

func (u *userUsecase) UpdateStudent(student *models.Student) error {
	return u.repo.Update(student)
}

func (u *userUsecase) DeleteStudent(ID database.PID) error {
	return u.repo.Delete(ID)
}

func (u *userUsecase) GetAllStudents() ([]models.Student, error) {
	return u.repo.GetAll()
}
