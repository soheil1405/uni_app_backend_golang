package usecases

import (
	"errors"
	"time"
	"uni_app/database"
	"uni_app/models"
	repositories "uni_app/pkg/user/repository"
	"uni_app/services/env"
	"uni_app/utils/helpers"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

type UserUsecase interface {
	CreateUser(user *models.User) error
	GetUserByID(ctx echo.Context, ID database.PID, useCache bool) (*models.User, error)
	UpdateUser(user *models.User) error
	DeleteUser(ID database.PID) error
	GetAllUsers(ctx echo.Context, request models.FetchUserRequest) ([]models.User, *helpers.PaginateTemplate, error)
	RegisterUser(ctx echo.Context, request *models.UserRegisterRequest) error
	LoginUser(username, password string) (*models.User, error)
	ValidateToken(token string) (*models.User, error)
}

type userUsecase struct {
	repo   repositories.UserRepository
	config *env.Config
}

func NewUserUsecase(repo repositories.UserRepository, config *env.Config) UserUsecase {
	return &userUsecase{repo, config}
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

func (u *userUsecase) GetAllUsers(ctx echo.Context, request models.FetchUserRequest) ([]models.User, *helpers.PaginateTemplate, error) {
	return u.repo.GetAll(ctx, request)
}

func (u *userUsecase) RegisterUser(ctx echo.Context, request *models.UserRegisterRequest) error {
	// Validate required fields
	if request.UserName == "" {
		return errors.New("username is required")
	}
	if request.Email == "" {
		return errors.New("email is required")
	}
	if request.Password == "" {
		return errors.New("password is required")
	}
	fetchRequest := models.FetchUserRequest{
		DegreeLevel:   request.DegreeLevel,
		DegreeMajorID: request.DegreeMajorID,
		DegreeUniID:   request.DegreeUniID,
		NationalCode:  request.NationalCode,
		Email:         request.Email,
		Number:        request.Number,
		PersonalCode:  request.PersonalCode,
		FetchRequest: models.FetchRequest{
			Limit:  1,
			Offset: 0,
		},
	}
	existingUser, _, err := u.repo.GetAll(ctx, fetchRequest)
	if err != nil || len(existingUser) > 0 {
		return errors.New("username already exists")
	}

	user := &models.User{
		UserName:      request.UserName,
		Email:         request.Email,
		Status:        models.USER_STATUS_ACTIVE,
		FirstName:     request.FirstName,
		LastName:      request.LastName,
		Number:        request.Number,
		PersonalCode:  request.PersonalCode,
		DegreeLevel:   request.DegreeLevel,
		DegreeMajorID: request.DegreeMajorID,
	}

	// Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(hashedPassword)
	user.Status = models.USER_STATUS_ACTIVE
	if err := u.repo.Create(user); err != nil {
		return err
	}

	// Generate JWT token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": user.ID,
		// "role":    user.,
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	})

	// Sign the token with the secret key
	tokenString, err := token.SignedString([]byte(u.config.GetString("JWT_SECRET")))
	if err != nil {
		return err
	}

	// Set the token in the user object
	user.Token.TokenKey = tokenString

	return nil
}

func (u *userUsecase) LoginUser(username, password string) (*models.User, error) {
	user, err := u.repo.GetByUsername(username)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, errors.New("invalid username or password")
	}

	// Verify password
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return nil, errors.New("invalid username or password")
	}

	// Generate JWT token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": user.ID,
		// "role":    user.,
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	})

	// Sign the token with the secret key
	tokenString, err := token.SignedString([]byte(u.config.GetString("JWT_SECRET")))
	if err != nil {
		return nil, err
	}

	// Set the token in the user object
	user.Token.TokenKey = tokenString

	return user, nil
}

func (u *userUsecase) ValidateToken(tokenString string) (*models.User, error) {
	// Parse the token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Validate the signing method
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return []byte(u.config.GetString("JWT_SECRET")), nil
	})

	if err != nil {
		return nil, err
	}

	// Check if the token is valid
	if !token.Valid {
		return nil, errors.New("invalid token")
	}

	// Get the claims
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, errors.New("invalid token claims")
	}

	// Get the user ID from the claims
	userID, ok := claims["user_id"].(float64)
	if !ok {
		return nil, errors.New("invalid user ID in token")
	}

	// Get the user from the database
	user, err := u.repo.GetByID(nil, database.PID(userID), false)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, errors.New("user not found")
	}

	return user, nil
}
