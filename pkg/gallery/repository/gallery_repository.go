package repository

import (
	"context"
	"uni_app/models"

	"gorm.io/gorm"
)

type GalleryRepository interface {
	Create(ctx context.Context, gallery *models.Gallery) error
	GetByID(ctx context.Context, id uint) (*models.Gallery, error)
	Update(ctx context.Context, gallery *models.Gallery) error
	Delete(ctx context.Context, id uint) error
	GetByImageable(ctx context.Context, imageableID uint, imageableType string) ([]*models.Gallery, error)
	GetMainImage(ctx context.Context, imageableID uint, imageableType string) (*models.Gallery, error)
}

type galleryRepository struct {
	db *gorm.DB
}

func NewGalleryRepository(db *gorm.DB) GalleryRepository {
	return &galleryRepository{db: db}
}

func (r *galleryRepository) Create(ctx context.Context, gallery *models.Gallery) error {
	return r.db.WithContext(ctx).Create(gallery).Error
}

func (r *galleryRepository) GetByID(ctx context.Context, id uint) (*models.Gallery, error) {
	var gallery models.Gallery
	err := r.db.WithContext(ctx).First(&gallery, id).Error
	if err != nil {
		return nil, err
	}
	return &gallery, nil
}

func (r *galleryRepository) Update(ctx context.Context, gallery *models.Gallery) error {
	return r.db.WithContext(ctx).Save(gallery).Error
}

func (r *galleryRepository) Delete(ctx context.Context, id uint) error {
	return r.db.WithContext(ctx).Delete(&models.Gallery{}, id).Error
}

func (r *galleryRepository) GetByImageable(ctx context.Context, OwnerID uint, OwnerType string) ([]*models.Gallery, error) {
	var galleries []*models.Gallery
	err := r.db.WithContext(ctx).
		Where("owner_id = ? AND owner_type = ?", OwnerID, OwnerType).
		Order("`order` ASC").
		Find(&galleries).Error
	if err != nil {
		return nil, err
	}
	return galleries, nil
}

func (r *galleryRepository) GetMainImage(ctx context.Context, OwnerID uint, OwnerType string) (*models.Gallery, error) {
	var gallery models.Gallery
	err := r.db.WithContext(ctx).
		Where("owner_id = ? AND owner_type = ? AND is_main = ?", OwnerID, OwnerType, true).
		First(&gallery).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &gallery, nil
}
