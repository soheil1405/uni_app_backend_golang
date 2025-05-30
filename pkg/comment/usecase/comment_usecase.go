package usecase

import (
	"context"
	"errors"

	"uni_app/database"
	"uni_app/models"
	"uni_app/pkg/comment/repository"
)

type CommentUseCase interface {
	CreateComment(ctx context.Context, comment *models.Comment) error
	GetCommentByID(ctx context.Context, id uint) (*models.Comment, error)
	UpdateComment(ctx context.Context, comment *models.Comment) error
	DeleteComment(ctx context.Context, id uint) error
	GetCommentsByCommentable(ctx context.Context, commentableID database.PID, OwnerType string) ([]*models.Comment, error)
	GetCommentsByUser(ctx context.Context, userID database.PID) ([]*models.Comment, error)
	GetCommentReplies(ctx context.Context, parentID database.PID) ([]*models.Comment, error)
	GetAllComments(ctx context.Context, request models.FetchCommentRequest) ([]*models.Comment, error)
}

type commentUseCase struct {
	commentRepo repository.CommentRepository
}

func NewCommentUseCase(commentRepo repository.CommentRepository) CommentUseCase {
	return &commentUseCase{
		commentRepo: commentRepo,
	}
}

func (uc *commentUseCase) CreateComment(ctx context.Context, comment *models.Comment) error {
	if comment.Content == "" {
		return errors.New("comment content is required")
	}
	if comment.UserID == 0 {
		return errors.New("user ID is required")
	}
	if comment.OwnerID == 0 || comment.OwnerType == "" {
		return errors.New("owner ID and type are required")
	}
	return uc.commentRepo.Create(ctx, comment)
}

func (uc *commentUseCase) GetCommentByID(ctx context.Context, id uint) (*models.Comment, error) {
	return uc.commentRepo.GetByID(ctx, id)
}

func (uc *commentUseCase) UpdateComment(ctx context.Context, comment *models.Comment) error {
	if comment.Content == "" {
		return errors.New("comment content is required")
	}
	return uc.commentRepo.Update(ctx, comment)
}

func (uc *commentUseCase) DeleteComment(ctx context.Context, id uint) error {
	return uc.commentRepo.Delete(ctx, id)
}

func (uc *commentUseCase) GetCommentsByCommentable(ctx context.Context, commentableID database.PID, OwnerType string) ([]*models.Comment, error) {
	return uc.commentRepo.GetByCommentable(ctx, commentableID, OwnerType)
}

func (uc *commentUseCase) GetCommentsByUser(ctx context.Context, userID database.PID) ([]*models.Comment, error) {
	return uc.commentRepo.GetByUser(ctx, userID)
}

func (uc *commentUseCase) GetCommentReplies(ctx context.Context, parentID database.PID) ([]*models.Comment, error) {
	return uc.commentRepo.GetReplies(ctx, parentID)
}

func (uc *commentUseCase) GetAllComments(ctx context.Context, request models.FetchCommentRequest) ([]*models.Comment, error) {
	return uc.commentRepo.GetAll(ctx, request)
}
