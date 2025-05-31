package repository

import (
	"context"
	"time"
	"uni_app/database"
	"uni_app/models"
	"uni_app/utils/helpers"

	"gorm.io/gorm"
)

type AttentionRepository interface {
	Create(ctx context.Context, attention *models.Attention) error
	GetByID(ctx context.Context, id uint) (*models.Attention, error)
	Update(ctx context.Context, attention *models.Attention) error
	Delete(ctx context.Context, id uint) error
	GetByRecipient(ctx context.Context, recipientID database.PID, recipientType string) ([]*models.Attention, error)
	GetActive(ctx context.Context) ([]*models.Attention, error)
	GetAll(ctx context.Context, request models.FetchAttentionRequest) ([]*models.Attention, *helpers.PaginateTemplate, error)
	UpdateStatus(ctx context.Context, id uint, status models.AttentionStatus) error
}

type attentionRepository struct {
	db *gorm.DB
}

func NewAttentionRepository(db *gorm.DB) AttentionRepository {
	return &attentionRepository{db: db}
}

func (r *attentionRepository) Create(ctx context.Context, attention *models.Attention) error {
	return r.db.WithContext(ctx).Create(attention).Error
}

func (r *attentionRepository) GetByID(ctx context.Context, id uint) (*models.Attention, error) {
	var attention models.Attention
	err := r.db.WithContext(ctx).First(&attention, id).Error
	if err != nil {
		return nil, err
	}
	return &attention, nil
}

func (r *attentionRepository) Update(ctx context.Context, attention *models.Attention) error {
	return r.db.WithContext(ctx).Save(attention).Error
}

func (r *attentionRepository) Delete(ctx context.Context, id uint) error {
	return r.db.WithContext(ctx).Delete(&models.Attention{}, id).Error
}

func (r *attentionRepository) GetByRecipient(ctx context.Context, recipientID database.PID, recipientType string) ([]*models.Attention, error) {
	var attentions []*models.Attention
	err := r.db.WithContext(ctx).
		Where("recipient_id = ? AND recipient_type = ?", recipientID, recipientType).
		Order("created_at DESC").
		Find(&attentions).Error
	if err != nil {
		return nil, err
	}
	return attentions, nil
}

func (r *attentionRepository) GetActive(ctx context.Context) ([]*models.Attention, error) {
	var attentions []*models.Attention
	err := r.db.WithContext(ctx).
		Where("status = ?", models.AttentionStatusActive).
		Order("created_at DESC").
		Find(&attentions).Error
	if err != nil {
		return nil, err
	}
	return attentions, nil
}

func (r *attentionRepository) GetAll(ctx context.Context, request models.FetchAttentionRequest) ([]*models.Attention, *helpers.PaginateTemplate, error) {
	var attentions []*models.Attention
	var total int64

	query := r.db.WithContext(ctx).Model(&models.Attention{})

	// Apply filters
	if request.Type != "" {
		query = query.Where("type = ?", request.Type)
	}
	if request.Status != "" {
		query = query.Where("status = ?", request.Status)
	}
	if request.RecipientID > 0 {
		query = query.Where("recipient_id = ?", request.RecipientID)
	}
	if request.RecipientType != "" {
		query = query.Where("recipient_type = ?", request.RecipientType)
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
	if err := query.Offset(offset).Limit(request.PerPage).Find(&attentions).Error; err != nil {
		return nil, nil, err
	}

	// Create pagination template
	paginate := &helpers.PaginateTemplate{
		Page:    request.Page,
		PerPage: request.PerPage,
		Total:   total,
	}

	return attentions, paginate, nil
}

func (r *attentionRepository) UpdateStatus(ctx context.Context, id uint, status models.AttentionStatus) error {
	updates := map[string]interface{}{
		"status": status,
	}

	if status == models.AttentionStatusRead {
		now := time.Now()
		updates["read_at"] = now
	} else if status == models.AttentionStatusArchived {
		now := time.Now()
		updates["archived_at"] = now
	}

	return r.db.WithContext(ctx).
		Model(&models.Attention{}).
		Where("id = ?", id).
		Updates(updates).Error
}
