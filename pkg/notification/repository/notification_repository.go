package repository

import (
	"context"
	"time"
	"uni_app/database"
	"uni_app/models"
	"uni_app/utils/helpers"

	"gorm.io/gorm"
)

type NotificationRepository interface {
	Create(ctx context.Context, notification *models.Notification) error
	GetByID(ctx context.Context, id uint) (*models.Notification, error)
	Update(ctx context.Context, notification *models.Notification) error
	Delete(ctx context.Context, id uint) error
	GetByRecipient(ctx context.Context, recipientID database.PID, recipientType string) ([]*models.Notification, error)
	GetPending(ctx context.Context) ([]*models.Notification, error)
	GetAll(ctx context.Context, request models.FetchNotificationRequest) ([]*models.Notification, *helpers.PaginateTemplate, error)
	UpdateStatus(ctx context.Context, id uint, status models.NotificationStatus, error string) error
	CreateTemplate(ctx context.Context, template *models.NotificationTemplate) error
	GetTemplateByName(ctx context.Context, name string) (*models.NotificationTemplate, error)
	GetPreference(ctx context.Context, ownerID database.PID, ownerType string) (*models.NotificationPreference, error)
	UpdatePreference(ctx context.Context, preference *models.NotificationPreference) error
}

type notificationRepository struct {
	db *gorm.DB
}

func NewNotificationRepository(db *gorm.DB) NotificationRepository {
	return &notificationRepository{db: db}
}

func (r *notificationRepository) Create(ctx context.Context, notification *models.Notification) error {
	return r.db.WithContext(ctx).Create(notification).Error
}

func (r *notificationRepository) GetByID(ctx context.Context, id uint) (*models.Notification, error) {
	var notification models.Notification
	err := r.db.WithContext(ctx).First(&notification, id).Error
	if err != nil {
		return nil, err
	}
	return &notification, nil
}

func (r *notificationRepository) Update(ctx context.Context, notification *models.Notification) error {
	return r.db.WithContext(ctx).Save(notification).Error
}

func (r *notificationRepository) Delete(ctx context.Context, id uint) error {
	return r.db.WithContext(ctx).Delete(&models.Notification{}, id).Error
}

func (r *notificationRepository) GetByRecipient(ctx context.Context, recipientID database.PID, recipientType string) ([]*models.Notification, error) {
	var notifications []*models.Notification
	err := r.db.WithContext(ctx).
		Where("recipient_id = ? AND recipient_type = ?", recipientID, recipientType).
		Order("created_at DESC").
		Find(&notifications).Error
	if err != nil {
		return nil, err
	}
	return notifications, nil
}

func (r *notificationRepository) GetPending(ctx context.Context) ([]*models.Notification, error) {
	var notifications []*models.Notification
	err := r.db.WithContext(ctx).
		Where("status = ?", models.NotificationStatusPending).
		Find(&notifications).Error
	if err != nil {
		return nil, err
	}
	return notifications, nil
}

func (r *notificationRepository) GetAll(ctx context.Context, request models.FetchNotificationRequest) ([]*models.Notification, *helpers.PaginateTemplate, error) {
	var notifications []*models.Notification
	var total int64

	query := r.db.WithContext(ctx).Model(&models.Notification{})

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
	if err := query.Offset(offset).Limit(request.PerPage).Find(&notifications).Error; err != nil {
		return nil, nil, err
	}

	// Create pagination template
	paginate := &helpers.PaginateTemplate{
		Page:    request.Page,
		PerPage: request.PerPage,
		Total:   total,
	}

	return notifications, paginate, nil
}

func (r *notificationRepository) UpdateStatus(ctx context.Context, id uint, status models.NotificationStatus, error string) error {
	updates := map[string]interface{}{
		"status": status,
		"error":  error,
	}

	if status == models.NotificationStatusSent {
		now := time.Now()
		updates["sent_at"] = now
	} else if status == models.NotificationStatusDelivered {
		now := time.Now()
		updates["delivered_at"] = now
	}

	return r.db.WithContext(ctx).
		Model(&models.Notification{}).
		Where("id = ?", id).
		Updates(updates).Error
}

func (r *notificationRepository) CreateTemplate(ctx context.Context, template *models.NotificationTemplate) error {
	return r.db.WithContext(ctx).Create(template).Error
}

func (r *notificationRepository) GetTemplateByName(ctx context.Context, name string) (*models.NotificationTemplate, error) {
	var template models.NotificationTemplate
	err := r.db.WithContext(ctx).
		Where("name = ?", name).
		First(&template).Error
	if err != nil {
		return nil, err
	}
	return &template, nil
}

func (r *notificationRepository) GetPreference(ctx context.Context, ownerID database.PID, ownerType string) (*models.NotificationPreference, error) {
	var preference models.NotificationPreference
	err := r.db.WithContext(ctx).
		Where("owner_id = ? AND owner_type = ?", ownerID, ownerType).
		First(&preference).Error
	if err != nil {
		return nil, err
	}
	return &preference, nil
}

func (r *notificationRepository) UpdatePreference(ctx context.Context, preference *models.NotificationPreference) error {
	return r.db.WithContext(ctx).Save(preference).Error
}
