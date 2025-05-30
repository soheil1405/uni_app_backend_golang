package repository

import (
	"context"
	"uni_app/database"
	"uni_app/models"

	"gorm.io/gorm"
)

type CommentRepository interface {
	Create(ctx context.Context, comment *models.Comment) error
	GetByID(ctx context.Context, id uint) (*models.Comment, error)
	Update(ctx context.Context, comment *models.Comment) error
	Delete(ctx context.Context, id uint) error
	GetByCommentable(ctx context.Context, commentableID database.PID, commentableType string) ([]*models.Comment, error)
	GetByUser(ctx context.Context, userID database.PID) ([]*models.Comment, error)
	GetReplies(ctx context.Context, parentID database.PID) ([]*models.Comment, error)
	GetAll(ctx context.Context, request models.FetchCommentRequest) ([]*models.Comment, error)
}

type commentRepository struct {
	db *gorm.DB
}

func NewCommentRepository(db *gorm.DB) CommentRepository {
	return &commentRepository{db: db}
}

func (r *commentRepository) Create(ctx context.Context, comment *models.Comment) error {
	return r.db.WithContext(ctx).Create(comment).Error
}

func (r *commentRepository) GetByID(ctx context.Context, id uint) (*models.Comment, error) {
	var comment models.Comment
	err := r.db.WithContext(ctx).
		Preload("User").
		Preload("Parent").
		Preload("Replies").
		First(&comment, id).Error
	if err != nil {
		return nil, err
	}
	return &comment, nil
}

func (r *commentRepository) Update(ctx context.Context, comment *models.Comment) error {
	return r.db.WithContext(ctx).Save(comment).Error
}

func (r *commentRepository) Delete(ctx context.Context, id uint) error {
	return r.db.WithContext(ctx).Delete(&models.Comment{}, id).Error
}

func (r *commentRepository) GetByCommentable(ctx context.Context, commentableID database.PID, commentableType string) ([]*models.Comment, error) {
	var comments []*models.Comment
	err := r.db.WithContext(ctx).
		Where("commentable_id = ? AND commentable_type = ? AND parent_id IS NULL", commentableID, commentableType).
		Preload("User").
		Preload("Replies").
		Find(&comments).Error
	if err != nil {
		return nil, err
	}
	return comments, nil
}

func (r *commentRepository) GetByUser(ctx context.Context, userID database.PID) ([]*models.Comment, error) {
	var comments []*models.Comment
	err := r.db.WithContext(ctx).
		Where("user_id = ?", userID).
		Preload("Parent").
		Find(&comments).Error
	if err != nil {
		return nil, err
	}
	return comments, nil
}

func (r *commentRepository) GetReplies(ctx context.Context, parentID database.PID) ([]*models.Comment, error) {
	var replies []*models.Comment
	err := r.db.WithContext(ctx).
		Where("parent_id = ?", parentID).
		Preload("User").
		Find(&replies).Error
	if err != nil {
		return nil, err
	}
	return replies, nil
}

func (r *commentRepository) GetAll(ctx context.Context, request models.FetchCommentRequest) ([]*models.Comment, error) {
	var comments []*models.Comment
	query := r.db.WithContext(ctx).Model(&models.Comment{})

	if request.UserID > 0 {
		query = query.Where("user_id = ?", request.UserID)
	}
	if request.CommentableID > 0 {
		query = query.Where("commentable_id = ?", request.CommentableID)
	}
	if request.CommentableType != "" {
		query = query.Where("commentable_type = ?", request.CommentableType)
	}
	if request.ParentID != nil {
		query = query.Where("parent_id = ?", request.ParentID)
	} else {
		query = query.Where("parent_id IS NULL")
	}

	// Apply includes
	if len(request.Includes) > 0 {
		for _, include := range request.Includes {
			query = query.Preload(include)
		}
	}

	if err := query.Find(&comments).Error; err != nil {
		return nil, err
	}
	return comments, nil
}
