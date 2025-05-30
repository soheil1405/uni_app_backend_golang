package repository

import (
	"context"
	"uni_app/pkg/model"

	"gorm.io/gorm"
)

type GalleryRepository interface {
	Create(ctx context.Context, gallery *model.Gallery) error
	GetByID(ctx context.Context, id uint) (*model.Gallery, error)
	Update(ctx context.Context, gallery *model.Gallery) error
	Delete(ctx context.Context, id uint) error
	GetByImageable(ctx context.Context, imageableID uint, imageableType string) ([]*model.Gallery, error)
	GetMainImage(ctx context.Context, imageableID uint, imageableType string) (*model.Gallery, error)
}

type galleryRepository struct {
	db *gorm.DB
}

func NewGalleryRepository(db *gorm.DB) GalleryRepository {
	return &galleryRepository{db: db}
}

func (r *galleryRepository) Create(ctx context.Context, gallery *model.Gallery) error {
	return r.db.WithContext(ctx).Create(gallery).Error
}

func (r *galleryRepository) GetByID(ctx context.Context, id uint) (*model.Gallery, error) {
	var gallery model.Gallery
	err := r.db.WithContext(ctx).First(&gallery, id).Error
	if err != nil {
		return nil, err
	}
	return &gallery, nil
}

func (r *galleryRepository) Update(ctx context.Context, gallery *model.Gallery) error {
	return r.db.WithContext(ctx).Save(gallery).Error
}

func (r *galleryRepository) Delete(ctx context.Context, id uint) error {
	return r.db.WithContext(ctx).Delete(&model.Gallery{}, id).Error
}

func (r *galleryRepository) GetByImageable(ctx context.Context, imageableID uint, imageableType string) ([]*model.Gallery, error) {
	var galleries []*model.Gallery
	err := r.db.WithContext(ctx).
		Where("imageable_id = ? AND imageable_type = ?", imageableID, imageableType).
		Order("`order` ASC").
		Find(&galleries).Error
	if err != nil {
		return nil, err
	}
	return galleries, nil
}

func (r *galleryRepository) GetMainImage(ctx context.Context, imageableID uint, imageableType string) (*model.Gallery, error) {
	var gallery model.Gallery
	err := r.db.WithContext(ctx).
		Where("imageable_id = ? AND imageable_type = ? AND is_main = ?", imageableID, imageableType, true).
		First(&gallery).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &gallery, nil
}
