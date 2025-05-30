package usecase

import (
	"context"
	"fmt"
	"uni_app/database"
	"uni_app/models"
	"uni_app/pkg/notification/repository"
	"uni_app/utils/helpers"
)

type NotificationUseCase interface {
	CreateNotification(ctx context.Context, notification *models.Notification) error
	GetNotificationByID(ctx context.Context, id database.PID) (*models.Notification, error)
	UpdateNotification(ctx context.Context, notification *models.Notification) error
	DeleteNotification(ctx context.Context, id database.PID) error
	GetNotificationsByRecipient(ctx context.Context, recipientID database.PID, recipientType string) ([]*models.Notification, error)
	GetPendingNotifications(ctx context.Context) ([]*models.Notification, error)
	GetAllNotifications(ctx context.Context, request models.FetchNotificationRequest) ([]*models.Notification, *helpers.PaginateTemplate, error)
	SendNotification(ctx context.Context, notification *models.Notification) error
	CreateTemplate(ctx context.Context, template *models.NotificationTemplate) error
	GetTemplateByName(ctx context.Context, name string) (*models.NotificationTemplate, error)
	GetPreference(ctx context.Context, ownerID database.PID, ownerType string) (*models.NotificationPreference, error)
	UpdatePreference(ctx context.Context, preference *models.NotificationPreference) error
}

type notificationUseCase struct {
	notificationRepo repository.NotificationRepository
	// Add other dependencies like SMS service, push notification service, etc.
}

func NewNotificationUseCase(notificationRepo repository.NotificationRepository) NotificationUseCase {
	return &notificationUseCase{
		notificationRepo: notificationRepo,
	}
}

func (uc *notificationUseCase) CreateNotification(ctx context.Context, notification *models.Notification) error {
	if notification.Type == "" {
		return fmt.Errorf("notification type is required")
	}
	if notification.Title == "" {
		return fmt.Errorf("notification title is required")
	}
	if notification.Body == "" {
		return fmt.Errorf("notification body is required")
	}
	if notification.RecipientID == 0 {
		return fmt.Errorf("recipient ID is required")
	}
	if notification.RecipientType == "" {
		return fmt.Errorf("recipient type is required")
	}

	// Set default status
	notification.Status = models.NotificationStatusPending

	return uc.notificationRepo.Create(ctx, notification)
}

func (uc *notificationUseCase) GetNotificationByID(ctx context.Context, id database.PID) (*models.Notification, error) {
	return uc.notificationRepo.GetByID(ctx, uint(id))
}

func (uc *notificationUseCase) UpdateNotification(ctx context.Context, notification *models.Notification) error {
	if notification.ID == 0 {
		return fmt.Errorf("notification ID is required")
	}
	return uc.notificationRepo.Update(ctx, notification)
}

func (uc *notificationUseCase) DeleteNotification(ctx context.Context, id database.PID) error {
	return uc.notificationRepo.Delete(ctx, uint(id))
}

func (uc *notificationUseCase) GetNotificationsByRecipient(ctx context.Context, recipientID database.PID, recipientType string) ([]*models.Notification, error) {
	return uc.notificationRepo.GetByRecipient(ctx, recipientID, recipientType)
}

func (uc *notificationUseCase) GetPendingNotifications(ctx context.Context) ([]*models.Notification, error) {
	return uc.notificationRepo.GetPending(ctx)
}

func (uc *notificationUseCase) GetAllNotifications(ctx context.Context, request models.FetchNotificationRequest) ([]*models.Notification, *helpers.PaginateTemplate, error) {
	return uc.notificationRepo.GetAll(ctx, request)
}

func (uc *notificationUseCase) SendNotification(ctx context.Context, notification *models.Notification) error {
	// Check if recipient has notification preferences
	preference, err := uc.notificationRepo.GetPreference(ctx, notification.RecipientID, notification.RecipientType)
	if err == nil && preference != nil {
		// Check if this type of notification is enabled for the recipient
		switch notification.Type {
		case models.NotificationTypePush:
			if !preference.PushEnabled {
				return fmt.Errorf("push notifications are disabled for recipient")
			}
		case models.NotificationTypeSMS:
			if !preference.SMSEnabled {
				return fmt.Errorf("SMS notifications are disabled for recipient")
			}
		default:
			return fmt.Errorf("unsupported notification type: %s", notification.Type)
		}
	}

	// Send notification based on type
	switch notification.Type {
	case models.NotificationTypePush:
		return uc.sendPushNotification(ctx, notification)
	case models.NotificationTypeSMS:
		return uc.sendSMSNotification(ctx, notification)
	default:
		return fmt.Errorf("unsupported notification type: %s", notification.Type)
	}
}

func (uc *notificationUseCase) sendPushNotification(ctx context.Context, notification *models.Notification) error {
	// TODO: Implement push notification sending logic
	// This would typically involve:
	// 1. Getting the recipient's device tokens
	// 2. Sending the notification to a push notification service (e.g., Firebase Cloud Messaging)
	// 3. Updating the notification status based on the result

	// For now, just mark it as sent
	return uc.notificationRepo.UpdateStatus(ctx, uint(notification.ID), models.NotificationStatusSent, "")
}

func (uc *notificationUseCase) sendSMSNotification(ctx context.Context, notification *models.Notification) error {
	// TODO: Implement SMS sending logic
	// This would typically involve:
	// 1. Getting the recipient's phone number
	// 2. Sending the SMS through an SMS service provider
	// 3. Updating the notification status based on the result

	// For now, just mark it as sent
	return uc.notificationRepo.UpdateStatus(ctx, uint(notification.ID), models.NotificationStatusSent, "")
}

func (uc *notificationUseCase) CreateTemplate(ctx context.Context, template *models.NotificationTemplate) error {
	if template.Name == "" {
		return fmt.Errorf("template name is required")
	}
	if template.Title == "" {
		return fmt.Errorf("template title is required")
	}
	if template.Body == "" {
		return fmt.Errorf("template body is required")
	}
	if template.Type == "" {
		return fmt.Errorf("template type is required")
	}

	return uc.notificationRepo.CreateTemplate(ctx, template)
}

func (uc *notificationUseCase) GetTemplateByName(ctx context.Context, name string) (*models.NotificationTemplate, error) {
	return uc.notificationRepo.GetTemplateByName(ctx, name)
}

func (uc *notificationUseCase) GetPreference(ctx context.Context, ownerID database.PID, ownerType string) (*models.NotificationPreference, error) {
	return uc.notificationRepo.GetPreference(ctx, ownerID, ownerType)
}

func (uc *notificationUseCase) UpdatePreference(ctx context.Context, preference *models.NotificationPreference) error {
	if preference.OwnerID == 0 {
		return fmt.Errorf("owner ID is required")
	}
	if preference.OwnerType == "" {
		return fmt.Errorf("owner type is required")
	}

	return uc.notificationRepo.UpdatePreference(ctx, preference)
}
