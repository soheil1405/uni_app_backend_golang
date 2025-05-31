package usecase

import (
	"context"
	"fmt"
	"uni_app/database"
	"uni_app/models"
	"uni_app/pkg/attention/repository"
	"uni_app/utils/helpers"
)

type AttentionUseCase interface {
	CreateAttention(ctx context.Context, attention *models.Attention) error
	GetAttentionByID(ctx context.Context, id database.PID) (*models.Attention, error)
	UpdateAttention(ctx context.Context, attention *models.Attention) error
	DeleteAttention(ctx context.Context, id database.PID) error
	GetAttentionsByRecipient(ctx context.Context, recipientID database.PID, recipientType string) ([]*models.Attention, error)
	GetActiveAttentions(ctx context.Context) ([]*models.Attention, error)
	GetAllAttentions(ctx context.Context, request models.FetchAttentionRequest) ([]*models.Attention, *helpers.PaginateTemplate, error)
	MarkAsRead(ctx context.Context, id database.PID) error
	MarkAsArchived(ctx context.Context, id database.PID) error
}

type attentionUseCase struct {
	attentionRepo repository.AttentionRepository
}

func NewAttentionUseCase(attentionRepo repository.AttentionRepository) AttentionUseCase {
	return &attentionUseCase{
		attentionRepo: attentionRepo,
	}
}

func (uc *attentionUseCase) CreateAttention(ctx context.Context, attention *models.Attention) error {
	if attention.Type == "" {
		return fmt.Errorf("attention type is required")
	}
	if attention.Title == "" {
		return fmt.Errorf("attention title is required")
	}
	if attention.Message == "" {
		return fmt.Errorf("attention message is required")
	}
	if attention.RecipientID == 0 {
		return fmt.Errorf("recipient ID is required")
	}
	if attention.RecipientType == "" {
		return fmt.Errorf("recipient type is required")
	}

	// Set default status
	attention.Status = models.AttentionStatusActive

	return uc.attentionRepo.Create(ctx, attention)
}

func (uc *attentionUseCase) GetAttentionByID(ctx context.Context, id database.PID) (*models.Attention, error) {
	return uc.attentionRepo.GetByID(ctx, uint(id))
}

func (uc *attentionUseCase) UpdateAttention(ctx context.Context, attention *models.Attention) error {
	if attention.ID == 0 {
		return fmt.Errorf("attention ID is required")
	}
	return uc.attentionRepo.Update(ctx, attention)
}

func (uc *attentionUseCase) DeleteAttention(ctx context.Context, id database.PID) error {
	return uc.attentionRepo.Delete(ctx, uint(id))
}

func (uc *attentionUseCase) GetAttentionsByRecipient(ctx context.Context, recipientID database.PID, recipientType string) ([]*models.Attention, error) {
	return uc.attentionRepo.GetByRecipient(ctx, recipientID, recipientType)
}

func (uc *attentionUseCase) GetActiveAttentions(ctx context.Context) ([]*models.Attention, error) {
	return uc.attentionRepo.GetActive(ctx)
}

func (uc *attentionUseCase) GetAllAttentions(ctx context.Context, request models.FetchAttentionRequest) ([]*models.Attention, *helpers.PaginateTemplate, error) {
	return uc.attentionRepo.GetAll(ctx, request)
}

func (uc *attentionUseCase) MarkAsRead(ctx context.Context, id database.PID) error {
	return uc.attentionRepo.UpdateStatus(ctx, uint(id), models.AttentionStatusRead)
}

func (uc *attentionUseCase) MarkAsArchived(ctx context.Context, id database.PID) error {
	return uc.attentionRepo.UpdateStatus(ctx, uint(id), models.AttentionStatusArchived)
}
