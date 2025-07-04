package usecase

import (
	"errors"
	"uni_app/database"
	"uni_app/models"
	repository "uni_app/pkg/student/repository"
	"uni_app/services/env"
	"uni_app/utils/helpers"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

type StudentUsecase interface {
	Create(student *models.Student) error
	Update(student *models.Student) error
	Delete(id database.PID) error
	GetByID(ctx echo.Context, id database.PID, usecache bool) (*models.Student, error)
	List(filters *models.FetchStudentRequest) ([]*models.Student, error)
	GetAllStudents(ctx echo.Context, request models.FetchStudentRequest) ([]models.Student, *helpers.PaginateTemplate, error)
	RegisterStudent(student *models.Student) error
	LoginStudent(studentCode database.PID, password string) (*models.Student, error)
	GetStudentByStudentCode(studentCode database.PID) (*models.Student, error)
	GetStudentByNationalCode(nationalCode database.PID) (*models.Student, error)
}

type studentUsecase struct {
	repo   repository.StudentRepository
	config *env.Config
}

func NewStudentUsecase(repo repository.StudentRepository, config *env.Config) StudentUsecase {
	return &studentUsecase{
		repo:   repo,
		config: config,
	}
}

func (u *studentUsecase) Create(student *models.Student) error {
	id, err := database.ParsePID(uuid.New().String())
	if err != nil {
		return err
	}
	student.ID = id
	return u.repo.Create(student)
}

func (u *studentUsecase) Update(student *models.Student) error {
	return u.repo.Update(student)
}

func (u *studentUsecase) Delete(id database.PID) error {
	return u.repo.Delete(id)
}

func (u *studentUsecase) GetByID(ctx echo.Context, id database.PID, usecache bool) (*models.Student, error) {
	return u.repo.GetByID(ctx, id, usecache)
}

func (u *studentUsecase) List(filters *models.FetchStudentRequest) ([]*models.Student, error) {
	return u.repo.List(filters)
}

func (u *studentUsecase) GetAllStudents(ctx echo.Context, request models.FetchStudentRequest) ([]models.Student, *helpers.PaginateTemplate, error) {
	return u.repo.GetAll(ctx, request)
}

func (u *studentUsecase) RegisterStudent(student *models.Student) error {
	// Validate required fields
	if student.StudentCode == 0 {
		return errors.New("student code is required")
	}
	if student.NationalCode == 0 {
		return errors.New("national code is required")
	}
	if student.Password == "" {
		return errors.New("password is required")
	}

	// Check if student code already exists
	existingStudent, err := u.repo.GetByStudentCode(student.StudentCode)
	if err != nil {
		return err
	}
	if existingStudent != nil {
		return errors.New("student code already exists")
	}

	// Check if national code already exists
	existingStudent, err = u.repo.GetByNationalCode(student.NationalCode)
	if err != nil {
		return err
	}
	if existingStudent != nil {
		return errors.New("national code already exists")
	}

	// Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(student.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	student.Password = string(hashedPassword)

	// Set default status
	student.Status = models.StudentStatusPending

	return u.repo.Create(student)
}

func (u *studentUsecase) LoginStudent(studentCode database.PID, password string) (*models.Student, error) {
	student, err := u.repo.GetByStudentCode(studentCode)
	if err != nil {
		return nil, err
	}
	if student == nil {
		return nil, errors.New("invalid student code or password")
	}

	// Verify password
	if err := bcrypt.CompareHashAndPassword([]byte(student.Password), []byte(password)); err != nil {
		return nil, errors.New("invalid student code or password")
	}

	return student, nil
}

func (u *studentUsecase) GetStudentByStudentCode(studentCode database.PID) (*models.Student, error) {
	return u.repo.GetByStudentCode(studentCode)
}

func (u *studentUsecase) GetStudentByNationalCode(nationalCode database.PID) (*models.Student, error) {
	return u.repo.GetByNationalCode(nationalCode)
}
