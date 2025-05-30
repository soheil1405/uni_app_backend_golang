package repository

import (
	"context"

	"uni_app/models"
	"uni_app/utils/helpers"

	"gorm.io/gorm"
)

type ArticleRepository interface {
	Create(ctx context.Context, article *models.Article) error
	GetByID(ctx context.Context, id uint) (*models.Article, error)
	GetBySlug(ctx context.Context, slug string) (*models.Article, error)
	Update(ctx context.Context, article *models.Article) error
	Delete(ctx context.Context, id uint) error
	GetByAuthor(ctx context.Context, authorID uint) ([]*models.Article, error)
	GetPublished(ctx context.Context) ([]*models.Article, error)
	GetByCategory(ctx context.Context, categoryID uint) ([]*models.Article, error)
	GetByTag(ctx context.Context, tagID uint) ([]*models.Article, error)
	Search(ctx context.Context, query string) ([]*models.Article, error)
	IncrementViews(ctx context.Context, id uint) error
	IncrementLikes(ctx context.Context, id uint) error
	DecrementLikes(ctx context.Context, id uint) error
	GetAll(ctx context.Context, request models.FetchArticleRequest) ([]*models.Article, *helpers.PaginateTemplate, error)
}

type articleRepository struct {
	db *gorm.DB
}

func NewArticleRepository(db *gorm.DB) ArticleRepository {
	return &articleRepository{db: db}
}

func (r *articleRepository) Create(ctx context.Context, article *models.Article) error {
	return r.db.WithContext(ctx).Create(article).Error
}

func (r *articleRepository) GetByID(ctx context.Context, id uint) (*models.Article, error) {
	var article models.Article
	err := r.db.WithContext(ctx).
		Preload("Author").
		Preload("Tags").
		Preload("Categories").
		Preload("Comments").
		First(&article, id).Error
	if err != nil {
		return nil, err
	}
	return &article, nil
}

func (r *articleRepository) GetBySlug(ctx context.Context, slug string) (*models.Article, error) {
	var article models.Article
	err := r.db.WithContext(ctx).
		Preload("Author").
		Preload("Tags").
		Preload("Categories").
		Preload("Comments").
		Where("slug = ?", slug).
		First(&article).Error
	if err != nil {
		return nil, err
	}
	return &article, nil
}

func (r *articleRepository) Update(ctx context.Context, article *models.Article) error {
	return r.db.WithContext(ctx).Save(article).Error
}

func (r *articleRepository) Delete(ctx context.Context, id uint) error {
	return r.db.WithContext(ctx).Delete(&models.Article{}, id).Error
}

func (r *articleRepository) GetByAuthor(ctx context.Context, authorID uint) ([]*models.Article, error) {
	var articles []*models.Article
	err := r.db.WithContext(ctx).
		Where("author_id = ? AND status = ?", authorID, "published").
		Order("published_at DESC").
		Find(&articles).Error
	if err != nil {
		return nil, err
	}
	return articles, nil
}

func (r *articleRepository) GetPublished(ctx context.Context) ([]*models.Article, error) {
	var articles []*models.Article
	err := r.db.WithContext(ctx).
		Where("status = ?", "published").
		Order("published_at DESC").
		Find(&articles).Error
	if err != nil {
		return nil, err
	}
	return articles, nil
}

func (r *articleRepository) GetByCategory(ctx context.Context, categoryID uint) ([]*models.Article, error) {
	var articles []*models.Article
	err := r.db.WithContext(ctx).
		Joins("JOIN article_categories ON article_categories.article_id = articles.id").
		Where("article_categories.category_id = ? AND articles.status = ?", categoryID, "published").
		Order("articles.published_at DESC").
		Find(&articles).Error
	if err != nil {
		return nil, err
	}
	return articles, nil
}

func (r *articleRepository) GetByTag(ctx context.Context, tagID uint) ([]*models.Article, error) {
	var articles []*models.Article
	err := r.db.WithContext(ctx).
		Joins("JOIN article_tags ON article_tags.article_id = articles.id").
		Where("article_tags.tag_id = ? AND articles.status = ?", tagID, "published").
		Order("articles.published_at DESC").
		Find(&articles).Error
	if err != nil {
		return nil, err
	}
	return articles, nil
}

func (r *articleRepository) Search(ctx context.Context, query string) ([]*models.Article, error) {
	var articles []*models.Article
	err := r.db.WithContext(ctx).
		Where("(title LIKE ? OR content LIKE ?) AND status = ?",
			"%"+query+"%", "%"+query+"%", "published").
		Order("published_at DESC").
		Find(&articles).Error
	if err != nil {
		return nil, err
	}
	return articles, nil
}

func (r *articleRepository) IncrementViews(ctx context.Context, id uint) error {
	return r.db.WithContext(ctx).
		Model(&models.Article{}).
		Where("id = ?", id).
		UpdateColumn("views", gorm.Expr("views + ?", 1)).Error
}

func (r *articleRepository) IncrementLikes(ctx context.Context, id uint) error {
	return r.db.WithContext(ctx).
		Model(&models.Article{}).
		Where("id = ?", id).
		UpdateColumn("likes", gorm.Expr("likes + ?", 1)).Error
}

func (r *articleRepository) DecrementLikes(ctx context.Context, id uint) error {
	return r.db.WithContext(ctx).
		Model(&models.Article{}).
		Where("id = ?", id).
		UpdateColumn("likes", gorm.Expr("likes - ?", 1)).Error
}

func (r *articleRepository) GetAll(ctx context.Context, request models.FetchArticleRequest) ([]*models.Article, *helpers.PaginateTemplate, error) {
	var articles []*models.Article
	var total int64

	query := r.db.WithContext(ctx).Model(&models.Article{})

	// Apply filters
	if request.Search != "" {
		query = query.Where("title LIKE ? OR content LIKE ?", "%"+request.Search+"%", "%"+request.Search+"%")
	}
	if request.Status != "" {
		query = query.Where("status = ?", request.Status)
	}
	if request.AuthorID != 0 {
		query = query.Where("author_id = ?", request.AuthorID)
	}
	if request.TagID != 0 {
		query = query.Joins("JOIN article_tags ON article_tags.article_id = articles.id").
			Where("article_tags.tag_id = ?", request.TagID)
	}
	if request.CategoryID != 0 {
		query = query.Joins("JOIN article_categories ON article_categories.article_id = articles.id").
			Where("article_categories.category_id = ?", request.CategoryID)
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
	if err := query.Offset(offset).Limit(request.PerPage).
		Preload("Author").
		Preload("Tags").
		Preload("Categories").
		Find(&articles).Error; err != nil {
		return nil, nil, err
	}

	// Create pagination template
	paginate := &helpers.PaginateTemplate{
		Page:    request.Page,
		PerPage: request.PerPage,
		Total:   total,
	}

	return articles, paginate, nil
}
