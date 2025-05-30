package usecase

import (
	"context"
	"uni_app/pkg/gallery/repository"
	"uni_app/pkg/model"
)

type GalleryUsecase interface {
	CreateGallery(ctx context.Context, gallery *model.Gallery) error
	GetGalleryByID(ctx context.Context, id uint) (*model.Gallery, error)
	UpdateGallery(ctx context.Context, gallery *model.Gallery) error
	DeleteGallery(ctx context.Context, id uint) error
	GetGalleriesByImageable(ctx context.Context, imageableID uint, imageableType string) ([]*model.Gallery, error)
	GetMainImage(ctx context.Context, imageableID uint, imageableType string) (*model.Gallery, error)
	SetMainImage(ctx context.Context, galleryID uint) error
}

type galleryUsecase struct {
	galleryRepo repository.GalleryRepository
}

func NewGalleryUsecase(galleryRepo repository.GalleryRepository) GalleryUsecase {
	return &galleryUsecase{
		galleryRepo: galleryRepo,
	}
}

func (u *galleryUsecase) CreateGallery(ctx context.Context, gallery *model.Gallery) error {
	return u.galleryRepo.Create(ctx, gallery)
}

func (u *galleryUsecase) GetGalleryByID(ctx context.Context, id uint) (*model.Gallery, error) {
	return u.galleryRepo.GetByID(ctx, id)
}

func (u *galleryUsecase) UpdateGallery(ctx context.Context, gallery *model.Gallery) error {
	return u.galleryRepo.Update(ctx, gallery)
}

func (u *galleryUsecase) DeleteGallery(ctx context.Context, id uint) error {
	return u.galleryRepo.Delete(ctx, id)
}

func (u *galleryUsecase) GetGalleriesByImageable(ctx context.Context, imageableID uint, imageableType string) ([]*model.Gallery, error) {
	return u.galleryRepo.GetByImageable(ctx, imageableID, imageableType)
}

func (u *galleryUsecase) GetMainImage(ctx context.Context, imageableID uint, imageableType string) (*model.Gallery, error) {
	return u.galleryRepo.GetMainImage(ctx, imageableID, imageableType)
}

func (u *galleryUsecase) SetMainImage(ctx context.Context, galleryID uint) error {
	gallery, err := u.galleryRepo.GetByID(ctx, galleryID)
	if err != nil {
		return err
	}

	// First, unset any existing main image for this imageable
	existingMain, err := u.galleryRepo.GetMainImage(ctx, gallery.ImageableID, gallery.ImageableType)
	if err == nil && existingMain != nil {
		existingMain.IsMain = false
		if err := u.galleryRepo.Update(ctx, existingMain); err != nil {
			return err
		}
	}

	// Set the new main image
	gallery.IsMain = true
	return u.galleryRepo.Update(ctx, gallery)
}
