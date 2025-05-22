package repositories

import (
	"errors"
	"uni_app/database"
	"uni_app/models"
	"uni_app/utils/helpers"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type UserRepository interface {
	Create(user *models.User) error
	GetByID(ctx echo.Context, ID database.PID, useCache bool) (*models.User, error)
	Update(user *models.User) error
	Delete(ID database.PID) error
	GetAll(ctx echo.Context, request models.FetchRequest) ([]models.User, *helpers.PaginateTemplate, error)
	GetByUsername(username string) (*models.User, error)
	GetByEmail(email string) (*models.User, error)
	GetByTeacherCode(teacherCode string) (*models.User, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db}
}

func (r *userRepository) Create(user *models.User) error {
	return r.db.Create(user).Error
}

func (r *userRepository) GetByID(ctx echo.Context, ID database.PID, useCache bool) (*models.User, error) {
	var user models.User
	if err := r.db.First(&user, ID).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userRepository) Update(user *models.User) error {
	return r.db.Save(user).Error
}

func (r *userRepository) Delete(ID database.PID) error {
	return r.db.Delete(&models.User{}, ID).Error
}

func (r *userRepository) GetAll(ctx echo.Context, request models.FetchRequest) ([]models.User, *helpers.PaginateTemplate, error) {
	var users []models.User
	query := r.db.Model(&models.User{})

	// Apply pagination
	paginate := helpers.NewPaginateTemplate(request.Page, request.Limit)
	query = paginate.Paginate(query)

	// Apply includes
	if len(request.Includes) > 0 {
		for _, include := range request.Includes {
			query = query.Preload(include)
		}
	}

	if err := query.Find(&users).Error; err != nil {
		return nil, nil, err
	}

	return users, paginate, nil
}

func (r *userRepository) GetByUsername(username string) (*models.User, error) {
	var user models.User
	if err := r.db.Where("user_name = ?", username).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}

func (r *userRepository) GetByEmail(email string) (*models.User, error) {
	var user models.User
	if err := r.db.Where("email = ?", email).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}

func (r *userRepository) GetByTeacherCode(teacherCode string) (*models.User, error) {
	var user models.User
	if err := r.db.Where("teacher_code = ?", teacherCode).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}
