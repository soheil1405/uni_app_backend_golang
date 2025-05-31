package usecase

import (
	"context"
	"fmt"
	"time"
	"uni_app/database"
	"uni_app/models"
	"uni_app/pkg/news/repository"
	"uni_app/utils/helpers"
)

type NewsUseCase interface {
	CreateNews(ctx context.Context, news *models.News) error
	GetNewsByID(ctx context.Context, id database.PID) (*models.News, error)
	UpdateNews(ctx context.Context, news *models.News) error
	DeleteNews(ctx context.Context, id database.PID) error
	GetNewsByOwner(ctx context.Context, ownerID database.PID, ownerType string) ([]*models.News, error)
	GetPublishedNews(ctx context.Context) ([]*models.News, error)
	GetAllNews(ctx context.Context, request models.FetchNewsRequest) ([]*models.News, *helpers.PaginateTemplate, error)
	PublishNews(ctx context.Context, id database.PID) error
	ArchiveNews(ctx context.Context, id database.PID) error
	GetPendingNotifications(ctx context.Context) ([]*models.News, error)
	ProcessNotifications(ctx context.Context) error
}

type newsUseCase struct {
	newsRepo repository.NewsRepository
}

func NewNewsUsecase(newsRepo repository.NewsRepository) NewsUseCase {
	return &newsUseCase{
		newsRepo: newsRepo,
	}
}

func (uc *newsUseCase) CreateNews(ctx context.Context, news *models.News) error {
	if news.Title == "" {
		return fmt.Errorf("news title is required")
	}
	if news.Content == "" {
		return fmt.Errorf("news content is required")
	}
	if news.OwnerID == 0 {
		return fmt.Errorf("owner ID is required")
	}
	if news.OwnerType == "" {
		return fmt.Errorf("owner type is required")
	}
	if news.AuthorID == 0 {
		return fmt.Errorf("author ID is required")
	}

	// Set default status
	if news.Status == "" {
		news.Status = models.NewsStatusDraft
	}

	// Set notification timing
	if news.Status == models.NewsStatusPublished {
		now := time.Now()
		news.PublishedAt = &now
		if news.NotifyAt == nil {
			news.NotifyAt = &now
		}
	}

	return uc.newsRepo.Create(ctx, news)
}

func (uc *newsUseCase) GetNewsByID(ctx context.Context, id database.PID) (*models.News, error) {
	return uc.newsRepo.GetByID(ctx, uint(id))
}

func (uc *newsUseCase) UpdateNews(ctx context.Context, news *models.News) error {
	if news.ID == 0 {
		return fmt.Errorf("news ID is required")
	}

	// If status is changing to published, set published_at
	if news.Status == models.NewsStatusPublished {
		now := time.Now()
		news.PublishedAt = &now
		if news.NotifyAt == nil {
			news.NotifyAt = &now
		}
	}

	return uc.newsRepo.Update(ctx, news)
}

func (uc *newsUseCase) DeleteNews(ctx context.Context, id database.PID) error {
	return uc.newsRepo.Delete(ctx, uint(id))
}

func (uc *newsUseCase) GetNewsByOwner(ctx context.Context, ownerID database.PID, ownerType string) ([]*models.News, error) {
	return uc.newsRepo.GetByOwner(ctx, ownerID, ownerType)
}

func (uc *newsUseCase) GetPublishedNews(ctx context.Context) ([]*models.News, error) {
	return uc.newsRepo.GetPublished(ctx)
}

func (uc *newsUseCase) GetAllNews(ctx context.Context, request models.FetchNewsRequest) ([]*models.News, *helpers.PaginateTemplate, error) {
	return uc.newsRepo.GetAll(ctx, request)
}

func (uc *newsUseCase) PublishNews(ctx context.Context, id database.PID) error {
	return uc.newsRepo.UpdateStatus(ctx, uint(id), models.NewsStatusPublished)
}

func (uc *newsUseCase) ArchiveNews(ctx context.Context, id database.PID) error {
	return uc.newsRepo.UpdateStatus(ctx, uint(id), models.NewsStatusArchived)
}

func (uc *newsUseCase) GetPendingNotifications(ctx context.Context) ([]*models.News, error) {
	return uc.newsRepo.GetPendingNotifications(ctx)
}

func (uc *newsUseCase) ProcessNotifications(ctx context.Context) error {
	// Get all pending notifications
	news, err := uc.newsRepo.GetPendingNotifications(ctx)
	if err != nil {
		return err
	}

	// Process each notification
	for _, n := range news {
		// TODO: Implement actual notification sending logic here
		// This could involve:
		// 1. Creating notifications for each recipient
		// 2. Sending emails
		// 3. Sending push notifications
		// 4. Sending SMS

		// Mark as notified after processing
		if err := uc.newsRepo.MarkAsNotified(ctx, uint(n.ID)); err != nil {
			return err
		}
	}

	return nil
}
