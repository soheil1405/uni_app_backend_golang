package usecase

import (
	"errors"
	"uni_app/models"
	repositories "uni_app/pkg/auth/repository"
	"uni_app/services/env"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type AuthUsecase interface {
	Register(user *models.User) error
	Login(username, password string) (*models.User, error)
	Update(user *models.User) error
	Delete(id string) error
	GetByID(id string) (*models.User, error)
	List(filters *models.FetchUserRequest) ([]*models.User, error)
}

type authUsecase struct {
	repo   repositories.AuthRepository
	config *env.AppConfig
}

func NewAuthUsecase(repo repositories.AuthRepository, config *env.AppConfig) AuthUsecase {
	return &authUsecase{
		repo:   repo,
		config: config,
	}
}

func (u *authUsecase) Register(user *models.User) error {
	// Check if username already exists
	existingUser, err := u.repo.GetByUsername(user.Username)
	if err == nil && existingUser != nil {
		return errors.New("username already exists")
	}

	// Check if email already exists
	existingUser, err = u.repo.GetByEmail(user.Email)
	if err == nil && existingUser != nil {
		return errors.New("email already exists")
	}

	// Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(hashedPassword)

	// Generate ID
	id := uuid.New().String()
	user.ID = id

	return u.repo.Create(user)
}

func (u *authUsecase) Login(username, password string) (*models.User, error) {
	user, err := u.repo.GetByUsername(username)
	if err != nil {
		return nil, errors.New("invalid username or password")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return nil, errors.New("invalid username or password")
	}

	return user, nil
}

func (u *authUsecase) Update(user *models.User) error {
	return u.repo.Update(user)
}

func (u *authUsecase) Delete(id string) error {
	return u.repo.Delete(id)
}

func (u *authUsecase) GetByID(id string) (*models.User, error) {
	return u.repo.GetByID(id)
}

func (u *authUsecase) List(filters *models.FetchUserRequest) ([]*models.User, error) {
	return u.repo.List(filters)
}
