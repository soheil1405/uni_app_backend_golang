package usecases

import (
	"errors"
	"time"
	"uni_app/database"
	"uni_app/models"
	tokenRepository "uni_app/pkg/token/repository"
	repositories "uni_app/pkg/user/repository"
	"uni_app/services/env"
	"uni_app/utils/helpers"
	"uni_app/utils/jwt"

	"github.com/labstack/echo/v4"
)

type UserUsecase interface {
	Login(ctx echo.Context, req *models.UserLoginRequst) (*models.Token, error)
	CreateUser(user *models.User) error
	GetUserByID(ctx echo.Context, ID database.PID, useCache bool) (*models.User, error)
	UpdateUser(user *models.User) error
	DeleteUser(ID database.PID) error
	GetAllUsers() ([]models.User, error)
}

type userUsecase struct {
	repo       repositories.UserRepository
	tokenRepo  tokenRepository.TokenRepository
	Config     *env.Config
	authConfig map[string]string
}

func NewUserUsecase(repo repositories.UserRepository, config *env.Config) UserUsecase {
	return &userUsecase{
		repo:       repo,
		Config:     config,
		authConfig: config.GetStringMapString("service.auth"),
	}
}

func (u *userUsecase) Login(ctx echo.Context, request *models.UserLoginRequst) (token *models.Token, err error) {
	var (
		user     *models.User
		tokenKey string
		expTime  time.Time
	)

	if request.Password == "" || request.Username == "" {
		return nil, models.ErrorInvalidUserPass
	}

	if user, err = u.repo.GetByUserName(request.Username); err != nil {
		return
	}
	if isValid := helpers.ComparePassword(user.Password, request.Password); !isValid {
		err = errors.New("wrong password")
		return
	}

	if tokenKey, expTime, err = jwt.GenerateToken(u.authConfig, user); err != nil {
		return
	}

	token = &models.Token{
		TokenKey:   tokenKey,
		Revoked:    false,
		ExpireTime: expTime,
		OwnerType:  "user",
		OwnerID:    user.ID,
	}

	if token, err = u.tokenRepo.Create(token); err != nil {
		return nil, err
	}

	token.User = user
	return
}

func (u *userUsecase) CreateUser(user *models.User) error {
	return u.repo.Create(user)
}

func (u *userUsecase) GetUserByID(ctx echo.Context, ID database.PID, useCache bool) (*models.User, error) {
	return u.repo.GetByID(ctx, ID, useCache)
}

func (u *userUsecase) UpdateUser(user *models.User) error {
	return u.repo.Update(user)
}

func (u *userUsecase) DeleteUser(ID database.PID) error {
	return u.repo.Delete(ID)
}

func (u *userUsecase) GetAllUsers() ([]models.User, error) {
	return u.repo.GetAll()
}
