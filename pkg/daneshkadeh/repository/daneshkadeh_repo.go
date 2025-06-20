package repositories

import (
	"errors"
	"fmt"
	"uni_app/database"
	"uni_app/models"
	"uni_app/utils/templates"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type FacultyRepository interface {
	Create(faculty *models.DaneshKadeh) error
	GetByID(ctx echo.Context, ID database.PID, useCache bool) (*models.DaneshKadeh, error)
	Update(faculty *models.DaneshKadeh) error
	Delete(ID database.PID) error
	GetAll(ctx echo.Context, request models.FetchDaneshKadehRequest) (models.DaneshKadeha, *templates.PaginateTemplate, error)
}

type facultyRepository struct {
	db *gorm.DB
}

func NewDaneshKadehRepository(db *gorm.DB) FacultyRepository {
	return &facultyRepository{db}
}

func (r *facultyRepository) Create(faculty *models.DaneshKadeh) error {
	return r.db.Create(faculty).Error
}

func (r *facultyRepository) GetByID(ctx echo.Context, ID database.PID, useCache bool) (*models.DaneshKadeh, error) {
	var faculty models.DaneshKadeh
	if err := r.db.First(&faculty, ID).Error; err != nil {
		return nil, err
	}
	return &faculty, nil
}

func (r *facultyRepository) Update(faculty *models.DaneshKadeh) error {
	return r.db.Save(faculty).Error
}

func (r *facultyRepository) Delete(ID database.PID) error {
	return r.db.Delete(&models.DaneshKadeha{}, ID).Error
}

func (r *facultyRepository) GetAll(ctx echo.Context, request models.FetchDaneshKadehRequest) (models.DaneshKadeha, *templates.PaginateTemplate, error) {
	var (
		daneshkadeha models.DaneshKadeha
		total        int64
		err          error
		limit        = request.Limit
		offset       = request.Offset
		search       = request.Search
		ids          = request.IDs
		filters      = request.Filters
		includes     = request.Includes
		sorts        = request.Sorts
		meta         *templates.PaginateTemplate
		query        = r.db
	)

	for key, value := range filters {
		switch v := value.(type) {
		case string:
			query = query.Where(fmt.Sprintf("%s = ?", key), v)
		case []string:
			query = query.Where(fmt.Sprintf("%s IN (?)", key), v)
		default:
			fmt.Printf("Unexpected type for key %s: %T\n", key, v)
		}
	}

	if len(ids) > 0 {
		query = query.Where("id in (?)", ids)
	}

	if search != "" {
		searchStr := "%" + search + "%"
		query = query.Where("name like ? ", searchStr)
	}

	if err = query.Model(&models.DaneshKadeh{}).Count(&total).Error; err != nil {
		return nil, nil, err
	}

	for _, sort := range sorts {
		query = query.Order(sort)
	}

	for _, include := range includes {
		query = query.Preload(include)
	}

	for {
		query.Limit(limit).Offset(offset).Find(&daneshkadeha)
		if limit > len(daneshkadeha) && int(total) > offset+limit {
			offset += limit
		} else {
			break
		}
	}

	if len(daneshkadeha) == 0 {
		return nil, nil, errors.New("no data found")
	}

	meta = templates.CreatePaginateTemplate(int(total), offset, limit)
	return daneshkadeha, meta, nil
}
