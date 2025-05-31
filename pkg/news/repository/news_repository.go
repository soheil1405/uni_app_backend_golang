package repository

import (
	"context"
	"time"
	"uni_app/database"
	"uni_app/models"
	"uni_app/utils/helpers"

	"gorm.io/gorm"
)

type NewsRepository interface {
	Create(ctx context.Context, news *models.News) error
	GetByID(ctx context.Context, id uint) (*models.News, error)
	Update(ctx context.Context, news *models.News) error
	Delete(ctx context.Context, id uint) error
	GetByOwner(ctx context.Context, ownerID database.PID, ownerType string) ([]*models.News, error)
	GetPublished(ctx context.Context) ([]*models.News, error)
	GetAll(ctx context.Context, request models.FetchNewsRequest) ([]*models.News, *helpers.PaginateTemplate, error)
	UpdateStatus(ctx context.Context, id uint, status models.NewsStatus) error
	GetPendingNotifications(ctx context.Context) ([]*models.News, error)
	MarkAsNotified(ctx context.Context, id uint) error
}

type newsRepository struct {
	db *gorm.DB
}

func NewNewsRepository(db *gorm.DB) NewsRepository {
	return &newsRepository{db: db}
}

func (r *newsRepository) Create(ctx context.Context, news *models.News) error {
	return r.db.WithContext(ctx).Create(news).Error
}

func (r *newsRepository) GetByID(ctx context.Context, id uint) (*models.News, error) {
	var news models.News
	err := r.db.WithContext(ctx).First(&news, id).Error
	if err != nil {
		return nil, err
	}
	return &news, nil
}

func (r *newsRepository) Update(ctx context.Context, news *models.News) error {
	return r.db.WithContext(ctx).Save(news).Error
}

func (r *newsRepository) Delete(ctx context.Context, id uint) error {
	return r.db.WithContext(ctx).Delete(&models.News{}, id).Error
}

func (r *newsRepository) GetByOwner(ctx context.Context, ownerID database.PID, ownerType string) ([]*models.News, error) {
	var news []*models.News
	err := r.db.WithContext(ctx).
		Where("owner_id = ? AND owner_type = ?", ownerID, ownerType).
		Order("created_at DESC").
		Find(&news).Error
	if err != nil {
		return nil, err
	}
	return news, nil
}

func (r *newsRepository) GetPublished(ctx context.Context) ([]*models.News, error) {
	var news []*models.News
	err := r.db.WithContext(ctx).
		Where("status = ?", models.NewsStatusPublished).
		Order("created_at DESC").
		Find(&news).Error
	if err != nil {
		return nil, err
	}
	return news, nil
}

func (r *newsRepository) GetAll(ctx context.Context, request models.FetchNewsRequest) ([]*models.News, *helpers.PaginateTemplate, error) {
	var news []*models.News
	var total int64

	query := r.db.WithContext(ctx).Model(&models.News{})

	// Apply filters
	if request.Status != "" {
		query = query.Where("status = ?", request.Status)
	}
	if request.OwnerID > 0 {
		query = query.Where("owner_id = ?", request.OwnerID)
	}
	if request.OwnerType != "" {
		query = query.Where("owner_type = ?", request.OwnerType)
	}
	if request.AuthorID > 0 {
		query = query.Where("author_id = ?", request.AuthorID)
	}
	if len(request.TagIDs) > 0 {
		query = query.Joins("JOIN news_tags ON news_tags.news_id = news.id").
			Where("news_tags.tag_id IN ?", request.TagIDs)
	}
	if request.Search != "" {
		query = query.Where("title ILIKE ? OR content ILIKE ?", "%"+request.Search+"%", "%"+request.Search+"%")
	}

	// Get total count
	if err := query.Count(&total).Error; err != nil {
		return nil, nil, err
	}

	// Apply pagination
	if request.Page < 1 {
		request.Page = 1
	}
	if request.PerPage < 1 {
		request.PerPage = 10
	}
	offset := (request.Page - 1) * request.PerPage

	// Apply sorting
	if request.SortBy != "" {
		order := "ASC"
		if request.SortDesc {
			order = "DESC"
		}
		query = query.Order(request.SortBy + " " + order)
	} else {
		query = query.Order("created_at DESC")
	}

	// Execute query
	if err := query.Offset(offset).Limit(request.PerPage).Find(&news).Error; err != nil {
		return nil, nil, err
	}

	// Create pagination template
	paginate := &helpers.PaginateTemplate{
		Page:    request.Page,
		PerPage: request.PerPage,
		Total:   total,
	}

	return news, paginate, nil
}

func (r *newsRepository) UpdateStatus(ctx context.Context, id uint, status models.NewsStatus) error {
	updates := map[string]interface{}{
		"status": status,
	}

	if status == models.NewsStatusPublished {
		now := time.Now()
		updates["published_at"] = now
	}

	return r.db.WithContext(ctx).
		Model(&models.News{}).
		Where("id = ?", id).
		Updates(updates).Error
}

func (r *newsRepository) GetPendingNotifications(ctx context.Context) ([]*models.News, error) {
	var news []*models.News
	now := time.Now()

	err := r.db.WithContext(ctx).
		Where("status = ? AND is_notified = ? AND notify_at <= ?",
			models.NewsStatusPublished, false, now).
		Find(&news).Error
	if err != nil {
		return nil, err
	}
	return news, nil
}

func (r *newsRepository) MarkAsNotified(ctx context.Context, id uint) error {
	return r.db.WithContext(ctx).
		Model(&models.News{}).
		Where("id = ?", id).
		Updates(map[string]interface{}{
			"is_notified": true,
			"notify_at":   time.Now(),
		}).Error
}
