package usecase

import (
	"context"

	"uni_app/models"
	"uni_app/pkg/social_media_link/repository"
)

type SocialMediaLinkUsecase interface {
	CreateLink(ctx context.Context, link *models.SocialMediaLink) error
	GetLinkByID(ctx context.Context, id uint) (*models.SocialMediaLink, error)
	UpdateLink(ctx context.Context, link *models.SocialMediaLink) error
	DeleteLink(ctx context.Context, id uint) error
	GetLinksByLinkable(ctx context.Context, linkableID uint, linkableType string) ([]*models.SocialMediaLink, error)
	GetLinkByPlatform(ctx context.Context, linkableID uint, linkableType string, platform string) (*models.SocialMediaLink, error)
}

type socialMediaLinkUsecase struct {
	socialMediaLinkRepo repository.SocialMediaLinkRepository
}

func NewSocialMediaLinkUsecase(socialMediaLinkRepo repository.SocialMediaLinkRepository) SocialMediaLinkUsecase {
	return &socialMediaLinkUsecase{
		socialMediaLinkRepo: socialMediaLinkRepo,
	}
}

func (u *socialMediaLinkUsecase) CreateLink(ctx context.Context, link *models.SocialMediaLink) error {
	return u.socialMediaLinkRepo.Create(ctx, link)
}

func (u *socialMediaLinkUsecase) GetLinkByID(ctx context.Context, id uint) (*models.SocialMediaLink, error) {
	return u.socialMediaLinkRepo.GetByID(ctx, id)
}

func (u *socialMediaLinkUsecase) UpdateLink(ctx context.Context, link *models.SocialMediaLink) error {
	return u.socialMediaLinkRepo.Update(ctx, link)
}

func (u *socialMediaLinkUsecase) DeleteLink(ctx context.Context, id uint) error {
	return u.socialMediaLinkRepo.Delete(ctx, id)
}

func (u *socialMediaLinkUsecase) GetLinksByLinkable(ctx context.Context, linkableID uint, linkableType string) ([]*models.SocialMediaLink, error) {
	return u.socialMediaLinkRepo.GetByLinkable(ctx, linkableID, linkableType)
}

func (u *socialMediaLinkUsecase) GetLinkByPlatform(ctx context.Context, linkableID uint, linkableType string, platform string) (*models.SocialMediaLink, error) {
	return u.socialMediaLinkRepo.GetByPlatform(ctx, linkableID, linkableType, platform)
}
