package repositories

import (
	"uni_app/database"
	"uni_app/models"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type TokenRepository interface {
	Create(token *models.Token) (*models.Token, error)
	GetByID(ctx echo.Context, ID database.PID, useCache bool) (*models.Token, error)
}

type tokenRepository struct {
	db *gorm.DB
}

func NewTokenRepository(db *gorm.DB) TokenRepository {
	return &tokenRepository{db}
}

func (r *tokenRepository) Create(token *models.Token) (*models.Token, error) {
	if err := r.db.Create(token).Error; err != nil {
		return nil, err
	}
	return token, nil
}

func (r *tokenRepository) GetByID(ctx echo.Context, ID database.PID, useCache bool) (*models.Token, error) {
	var token models.Token
	if err := r.db.First(&token, ID).Error; err != nil {
		return nil, err
	}
	return &token, nil
}
