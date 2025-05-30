package usecase

import (
	"context"
	"errors"
	"time"

	"uni_app/models"
	"uni_app/pkg/article/repository"
	"uni_app/utils/helpers"
)

type ArticleUseCase interface {
	CreateArticle(ctx context.Context, article *models.Article) error
	GetArticleByID(ctx context.Context, id uint) (*models.Article, error)
	GetArticleBySlug(ctx context.Context, slug string) (*models.Article, error)
	UpdateArticle(ctx context.Context, article *models.Article) error
	DeleteArticle(ctx context.Context, id uint) error
	GetArticlesByAuthor(ctx context.Context, authorID uint) ([]*models.Article, error)
	GetPublishedArticles(ctx context.Context) ([]*models.Article, error)
	GetArticlesByCategory(ctx context.Context, categoryID uint) ([]*models.Article, error)
	GetArticlesByTag(ctx context.Context, tagID uint) ([]*models.Article, error)
	SearchArticles(ctx context.Context, query string) ([]*models.Article, error)
	PublishArticle(ctx context.Context, id uint) error
	ArchiveArticle(ctx context.Context, id uint) error
	LikeArticle(ctx context.Context, id uint) error
	UnlikeArticle(ctx context.Context, id uint) error
	ViewArticle(ctx context.Context, id uint) error
	GetAllArticles(ctx context.Context, request models.FetchArticleRequest) ([]*models.Article, *helpers.PaginateTemplate, error)
}

type articleUseCase struct {
	articleRepo repository.ArticleRepository
}

func NewArticleUseCase(articleRepo repository.ArticleRepository) ArticleUseCase {
	return &articleUseCase{
		articleRepo: articleRepo,
	}
}

func (uc *articleUseCase) CreateArticle(ctx context.Context, article *models.Article) error {
	if article.Title == "" || article.Content == "" {
		return errors.New("title and content are required")
	}
	return uc.articleRepo.Create(ctx, article)
}

func (uc *articleUseCase) GetArticleByID(ctx context.Context, id uint) (*models.Article, error) {
	return uc.articleRepo.GetByID(ctx, id)
}

func (uc *articleUseCase) GetArticleBySlug(ctx context.Context, slug string) (*models.Article, error) {
	return uc.articleRepo.GetBySlug(ctx, slug)
}

func (uc *articleUseCase) UpdateArticle(ctx context.Context, article *models.Article) error {
	if article.Title == "" || article.Content == "" {
		return errors.New("title and content are required")
	}
	return uc.articleRepo.Update(ctx, article)
}

func (uc *articleUseCase) DeleteArticle(ctx context.Context, id uint) error {
	return uc.articleRepo.Delete(ctx, id)
}

func (uc *articleUseCase) GetArticlesByAuthor(ctx context.Context, authorID uint) ([]*models.Article, error) {
	return uc.articleRepo.GetByAuthor(ctx, authorID)
}

func (uc *articleUseCase) GetPublishedArticles(ctx context.Context) ([]*models.Article, error) {
	return uc.articleRepo.GetPublished(ctx)
}

func (uc *articleUseCase) GetArticlesByCategory(ctx context.Context, categoryID uint) ([]*models.Article, error) {
	return uc.articleRepo.GetByCategory(ctx, categoryID)
}

func (uc *articleUseCase) GetArticlesByTag(ctx context.Context, tagID uint) ([]*models.Article, error) {
	return uc.articleRepo.GetByTag(ctx, tagID)
}

func (uc *articleUseCase) SearchArticles(ctx context.Context, query string) ([]*models.Article, error) {
	if query == "" {
		return nil, errors.New("search query is required")
	}
	return uc.articleRepo.Search(ctx, query)
}

func (uc *articleUseCase) PublishArticle(ctx context.Context, id uint) error {
	article, err := uc.articleRepo.GetByID(ctx, id)
	if err != nil {
		return err
	}

	article.Status = "published"
	article.PublishedAt = time.Now()
	return uc.articleRepo.Update(ctx, article)
}

func (uc *articleUseCase) ArchiveArticle(ctx context.Context, id uint) error {
	article, err := uc.articleRepo.GetByID(ctx, id)
	if err != nil {
		return err
	}

	article.Status = "archived"
	return uc.articleRepo.Update(ctx, article)
}

func (uc *articleUseCase) LikeArticle(ctx context.Context, id uint) error {
	return uc.articleRepo.IncrementLikes(ctx, id)
}

func (uc *articleUseCase) UnlikeArticle(ctx context.Context, id uint) error {
	return uc.articleRepo.DecrementLikes(ctx, id)
}

func (uc *articleUseCase) ViewArticle(ctx context.Context, id uint) error {
	return uc.articleRepo.IncrementViews(ctx, id)
}

func (uc *articleUseCase) GetAllArticles(ctx context.Context, request models.FetchArticleRequest) ([]*models.Article, *helpers.PaginateTemplate, error) {
	return uc.articleRepo.GetAll(ctx, request)
}
