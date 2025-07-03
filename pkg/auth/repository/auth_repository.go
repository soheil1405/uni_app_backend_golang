package repositories

import (
	"uni_app/models"

	"gorm.io/gorm"
)

type AuthRepository interface {
	Create(user *models.User) error
	Update(user *models.User) error
	Delete(id string) error
	GetByID(id string) (*models.User, error)
	GetByUsername(username string) (*models.User, error)
	GetByEmail(email string) (*models.User, error)
	List(filters *models.FetchUserRequest) ([]*models.User, error)
}

type authRepository struct {
	db *gorm.DB
}

func NewAuthRepository(db *gorm.DB) AuthRepository {
	return &authRepository{db}
}

func (r *authRepository) Create(user *models.User) error {
	return r.db.Create(user).Error
}

func (r *authRepository) Update(user *models.User) error {
	return r.db.Save(user).Error
}

func (r *authRepository) Delete(id string) error {
	return r.db.Delete(&models.User{}, id).Error
}

func (r *authRepository) GetByID(id string) (*models.User, error) {
	var user models.User
	err := r.db.First(&user, id).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *authRepository) GetByUsername(username string) (*models.User, error) {
	var user models.User
	err := r.db.Where("username = ?", username).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *authRepository) GetByEmail(email string) (*models.User, error) {
	var user models.User
	err := r.db.Where("email = ?", email).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *authRepository) List(filters *models.FetchUserRequest) ([]*models.User, error) {
	var users []*models.User
	query := r.db.Model(&models.User{})

	if filters.Username != "" {
		query = query.Where("username LIKE ?", "%"+filters.Username+"%")
	}
	if filters.Email != "" {
		query = query.Where("email LIKE ?", "%"+filters.Email+"%")
	}
	if filters.Role != "" {
		query = query.Where("role = ?", filters.Role)
	}

	err := query.Find(&users).Error
	if err != nil {
		return nil, err
	}
	return users, nil
}
