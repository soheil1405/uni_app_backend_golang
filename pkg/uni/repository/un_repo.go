package repositories

import (
	"fmt"
	"uni_app/database"
	"uni_app/models"
	"uni_app/utils/templates"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type UniRepository interface {
	Create(uni *models.Uni) error
	GetByID(ctx echo.Context, ID database.PID, useCache bool) (*models.Uni, error)
	Update(uni *models.Uni) error
	Delete(ID database.PID) error
	GetAll(ctx echo.Context, request models.FetchRequest) ([]models.Uni, *templates.PaginateTemplate, error)
}

type uniRepository struct {
	db *gorm.DB
}

func NewUniRepository(db *gorm.DB) UniRepository {
	return &uniRepository{db}
}

func (r *uniRepository) Create(uni *models.Uni) error {
	return r.db.Create(uni).Error
}

func (r *uniRepository) GetByID(ctx echo.Context, ID database.PID, useCache bool) (*models.Uni, error) {
	var uni models.Uni
	if err := r.db.First(&uni, ID).Error; err != nil {
		return nil, err
	}
	return &uni, nil
}

func (r *uniRepository) Update(uni *models.Uni) error {
	return r.db.Save(uni).Error
}

func (r *uniRepository) Delete(ID database.PID) error {
	return r.db.Delete(&models.Uni{}, ID).Error
}

func (r *uniRepository) GetAll(ctx echo.Context, request models.FetchRequest) ([]models.Uni, *templates.PaginateTemplate, error) {
	var (
		unis     []models.Uni
		total    int64
		err      error
		limit    = request.Limit
		offset   = request.Offset
		search   = request.Search
		ids      = request.IDs
		filters  = request.Filters
		includes = request.Includes
		sorts    = request.Sorts
		meta     *templates.PaginateTemplate
		query    = r.db.Debug()
	)

	// Apply Filters
	for key, value := range filters {
		query = query.Where(fmt.Sprintf("%s = ?", key), value)
	}

	if len(ids) > 0 {
		query = query.Where("id in (?)", ids)
	}

	if search != "" {
		searchStr := "%" + search + "%"
		query = query.Where("name like ? ", searchStr)
	}

	// Count total records
	if err = query.Model(&models.Uni{}).Count(&total).Error; err != nil {
		return nil, nil, err
	}

	for _, sort := range sorts {
		query = query.Order(sort)
	}

	// Apply Includes (Relations)
	for _, include := range includes {
		query = query.Preload(include)
	}

	for {
		err = query.
			Limit(limit).
			Offset(offset).Find(&unis).Error

		if limit > len(unis) && int(total) > offset+limit {
			offset += limit
		} else {
			break
		}
	}
	meta = templates.CreatePaginateTemplate(int(total), offset, limit)

	return unis, meta, nil

}
