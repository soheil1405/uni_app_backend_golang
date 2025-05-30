package repository

import (
	"context"

	"uni_app/models"

	"gorm.io/gorm"
)

type SocialMediaLinkRepository interface {
	Create(ctx context.Context, link *models.SocialMediaLink) error
	GetByID(ctx context.Context, id uint) (*models.SocialMediaLink, error)
	Update(ctx context.Context, link *models.SocialMediaLink) error
	Delete(ctx context.Context, id uint) error
	GetByLinkable(ctx context.Context, linkableID uint, linkableType string) ([]*models.SocialMediaLink, error)
	GetByPlatform(ctx context.Context, linkableID uint, linkableType string, platform string) (*models.SocialMediaLink, error)
	GetActiveLinks(ctx context.Context) ([]*models.SocialMediaLink, error)
	GetLinksByPlatform(ctx context.Context, platform string) ([]*models.SocialMediaLink, error)
	UpdateOrder(ctx context.Context, id uint, order int) error
	BulkUpdateOrder(ctx context.Context, orders map[uint]int) error
}

type socialMediaLinkRepository struct {
	db *gorm.DB
}

func NewSocialMediaLinkRepository(db *gorm.DB) SocialMediaLinkRepository {
	return &socialMediaLinkRepository{db: db}
}

func (r *socialMediaLinkRepository) Create(ctx context.Context, link *models.SocialMediaLink) error {
	return r.db.WithContext(ctx).Create(link).Error
}

func (r *socialMediaLinkRepository) GetByID(ctx context.Context, id uint) (*models.SocialMediaLink, error) {
	var link models.SocialMediaLink
	err := r.db.WithContext(ctx).First(&link, id).Error
	if err != nil {
		return nil, err
	}
	return &link, nil
}

func (r *socialMediaLinkRepository) Update(ctx context.Context, link *models.SocialMediaLink) error {
	return r.db.WithContext(ctx).Save(link).Error
}

func (r *socialMediaLinkRepository) Delete(ctx context.Context, id uint) error {
	return r.db.WithContext(ctx).Delete(&models.SocialMediaLink{}, id).Error
}

func (r *socialMediaLinkRepository) GetByLinkable(ctx context.Context, linkableID uint, linkableType string) ([]*models.SocialMediaLink, error) {
	var links []*models.SocialMediaLink
	err := r.db.WithContext(ctx).
		Where("linkable_id = ? AND linkable_type = ? AND is_active = ?", linkableID, linkableType, true).
		Order("`order` ASC").
		Find(&links).Error
	if err != nil {
		return nil, err
	}
	return links, nil
}

func (r *socialMediaLinkRepository) GetByPlatform(ctx context.Context, linkableID uint, linkableType string, platform string) (*models.SocialMediaLink, error) {
	var link models.SocialMediaLink
	err := r.db.WithContext(ctx).
		Where("linkable_id = ? AND linkable_type = ? AND platform = ? AND is_active = ?",
			linkableID, linkableType, platform, true).
		First(&link).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &link, nil
}

func (r *socialMediaLinkRepository) GetActiveLinks(ctx context.Context) ([]*models.SocialMediaLink, error) {
	var links []*models.SocialMediaLink
	err := r.db.WithContext(ctx).
		Where("is_active = ?", true).
		Order("`order` ASC").
		Find(&links).Error
	if err != nil {
		return nil, err
	}
	return links, nil
}

func (r *socialMediaLinkRepository) GetLinksByPlatform(ctx context.Context, platform string) ([]*models.SocialMediaLink, error) {
	var links []*models.SocialMediaLink
	err := r.db.WithContext(ctx).
		Where("platform = ? AND is_active = ?", platform, true).
		Order("`order` ASC").
		Find(&links).Error
	if err != nil {
		return nil, err
	}
	return links, nil
}

func (r *socialMediaLinkRepository) UpdateOrder(ctx context.Context, id uint, order int) error {
	return r.db.WithContext(ctx).
		Model(&models.SocialMediaLink{}).
		Where("id = ?", id).
		Update("order", order).Error
}

func (r *socialMediaLinkRepository) BulkUpdateOrder(ctx context.Context, orders map[uint]int) error {
	return r.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		for id, order := range orders {
			if err := tx.Model(&models.SocialMediaLink{}).
				Where("id = ?", id).
				Update("order", order).Error; err != nil {
				return err
			}
		}
		return nil
	})
}
