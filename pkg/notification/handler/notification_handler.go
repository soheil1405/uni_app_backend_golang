package handler

import (
	"net/http"
	"uni_app/database"
	"uni_app/models"
	"uni_app/pkg/notification/usecase"
	"uni_app/utils/helpers"

	"github.com/labstack/echo/v4"
)

type NotificationHandler struct {
	notificationUseCase usecase.NotificationUseCase
}

func NewNotificationHandler(notificationUseCase usecase.NotificationUseCase) *NotificationHandler {
	return &NotificationHandler{
		notificationUseCase: notificationUseCase,
	}
}

func (h *NotificationHandler) RegisterRoutes(e *echo.Echo) {
	notifications := e.Group("/api/notifications")
	notifications.POST("", h.CreateNotification)
	notifications.GET("/:id", h.GetNotificationByID)
	notifications.PUT("/:id", h.UpdateNotification)
	notifications.DELETE("/:id", h.DeleteNotification)
	notifications.GET("/recipient/:id", h.GetNotificationsByRecipient)
	notifications.GET("/pending", h.GetPendingNotifications)
	notifications.GET("", h.GetAllNotifications)
	notifications.POST("/:id/send", h.SendNotification)

	templates := notifications.Group("/templates")
	templates.POST("", h.CreateTemplate)
	templates.GET("/:name", h.GetTemplateByName)

	preferences := notifications.Group("/preferences")
	preferences.GET("/:id", h.GetPreference)
	preferences.PUT("/:id", h.UpdatePreference)
}

func (h *NotificationHandler) CreateNotification(c echo.Context) error {
	var notification models.Notification
	if err := c.Bind(&notification); err != nil {
		return helpers.Reply(c, http.StatusBadRequest, err, nil, nil)
	}

	if err := h.notificationUseCase.CreateNotification(c.Request().Context(), &notification); err != nil {
		return helpers.Reply(c, http.StatusInternalServerError, err, nil, nil)
	}

	return helpers.Reply(c, http.StatusCreated, nil, map[string]interface{}{
		"notification": notification,
	}, nil)
}

func (h *NotificationHandler) GetNotificationByID(c echo.Context) error {
	id, err := database.ParsePID(c.Param("id"))
	if err != nil {
		return helpers.Reply(c, http.StatusBadRequest, err, nil, nil)
	}

	notification, err := h.notificationUseCase.GetNotificationByID(c.Request().Context(), id)
	if err != nil {
		return helpers.Reply(c, http.StatusInternalServerError, err, nil, nil)
	}

	return helpers.Reply(c, http.StatusOK, nil, map[string]interface{}{
		"notification": notification,
	}, nil)
}

func (h *NotificationHandler) UpdateNotification(c echo.Context) error {
	id, err := database.ParsePID(c.Param("id"))
	if err != nil {
		return helpers.Reply(c, http.StatusBadRequest, err, nil, nil)
	}

	var notification models.Notification
	if err := c.Bind(&notification); err != nil {
		return helpers.Reply(c, http.StatusBadRequest, err, nil, nil)
	}

	notification.ID = id
	if err := h.notificationUseCase.UpdateNotification(c.Request().Context(), &notification); err != nil {
		return helpers.Reply(c, http.StatusInternalServerError, err, nil, nil)
	}

	return helpers.Reply(c, http.StatusOK, nil, map[string]interface{}{
		"notification": notification,
	}, nil)
}

func (h *NotificationHandler) DeleteNotification(c echo.Context) error {
	id, err := database.ParsePID(c.Param("id"))
	if err != nil {
		return helpers.Reply(c, http.StatusBadRequest, err, nil, nil)
	}

	if err := h.notificationUseCase.DeleteNotification(c.Request().Context(), id); err != nil {
		return helpers.Reply(c, http.StatusInternalServerError, err, nil, nil)
	}

	return helpers.Reply(c, http.StatusOK, nil, nil, nil)
}

func (h *NotificationHandler) GetNotificationsByRecipient(c echo.Context) error {
	recipientID, err := database.ParsePID(c.Param("id"))
	if err != nil {
		return helpers.Reply(c, http.StatusBadRequest, err, nil, nil)
	}

	recipientType := c.QueryParam("type")
	if recipientType == "" {
		return helpers.Reply(c, http.StatusBadRequest, echo.NewHTTPError(http.StatusBadRequest, "recipient type is required"), nil, nil)
	}

	notifications, err := h.notificationUseCase.GetNotificationsByRecipient(c.Request().Context(), recipientID, recipientType)
	if err != nil {
		return helpers.Reply(c, http.StatusInternalServerError, err, nil, nil)
	}

	return helpers.Reply(c, http.StatusOK, nil, map[string]interface{}{
		"notifications": notifications,
	}, nil)
}

func (h *NotificationHandler) GetPendingNotifications(c echo.Context) error {
	notifications, err := h.notificationUseCase.GetPendingNotifications(c.Request().Context())
	if err != nil {
		return helpers.Reply(c, http.StatusInternalServerError, err, nil, nil)
	}

	return helpers.Reply(c, http.StatusOK, nil, map[string]interface{}{
		"notifications": notifications,
	}, nil)
}

func (h *NotificationHandler) GetAllNotifications(c echo.Context) error {
	var request models.FetchNotificationRequest
	if err := c.Bind(&request); err != nil {
		return helpers.Reply(c, http.StatusBadRequest, err, nil, nil)
	}

	notifications, paginate, err := h.notificationUseCase.GetAllNotifications(c.Request().Context(), request)
	if err != nil {
		return helpers.Reply(c, http.StatusInternalServerError, err, nil, nil)
	}

	return helpers.Reply(c, http.StatusOK, nil, map[string]interface{}{
		"notifications": notifications,
	}, paginate)
}

func (h *NotificationHandler) SendNotification(c echo.Context) error {
	id, err := database.ParsePID(c.Param("id"))
	if err != nil {
		return helpers.Reply(c, http.StatusBadRequest, err, nil, nil)
	}

	notification, err := h.notificationUseCase.GetNotificationByID(c.Request().Context(), id)
	if err != nil {
		return helpers.Reply(c, http.StatusInternalServerError, err, nil, nil)
	}

	if err := h.notificationUseCase.SendNotification(c.Request().Context(), notification); err != nil {
		return helpers.Reply(c, http.StatusInternalServerError, err, nil, nil)
	}

	return helpers.Reply(c, http.StatusOK, nil, map[string]interface{}{
		"notification": notification,
	}, nil)
}

func (h *NotificationHandler) CreateTemplate(c echo.Context) error {
	var template models.NotificationTemplate
	if err := c.Bind(&template); err != nil {
		return helpers.Reply(c, http.StatusBadRequest, err, nil, nil)
	}

	if err := h.notificationUseCase.CreateTemplate(c.Request().Context(), &template); err != nil {
		return helpers.Reply(c, http.StatusInternalServerError, err, nil, nil)
	}

	return helpers.Reply(c, http.StatusCreated, nil, map[string]interface{}{
		"template": template,
	}, nil)
}

func (h *NotificationHandler) GetTemplateByName(c echo.Context) error {
	name := c.Param("name")
	if name == "" {
		return helpers.Reply(c, http.StatusBadRequest, echo.NewHTTPError(http.StatusBadRequest, "template name is required"), nil, nil)
	}

	template, err := h.notificationUseCase.GetTemplateByName(c.Request().Context(), name)
	if err != nil {
		return helpers.Reply(c, http.StatusInternalServerError, err, nil, nil)
	}

	return helpers.Reply(c, http.StatusOK, nil, map[string]interface{}{
		"template": template,
	}, nil)
}

func (h *NotificationHandler) GetPreference(c echo.Context) error {
	ownerID, err := database.ParsePID(c.Param("id"))
	if err != nil {
		return helpers.Reply(c, http.StatusBadRequest, err, nil, nil)
	}

	ownerType := c.QueryParam("type")
	if ownerType == "" {
		return helpers.Reply(c, http.StatusBadRequest, echo.NewHTTPError(http.StatusBadRequest, "owner type is required"), nil, nil)
	}

	preference, err := h.notificationUseCase.GetPreference(c.Request().Context(), ownerID, ownerType)
	if err != nil {
		return helpers.Reply(c, http.StatusInternalServerError, err, nil, nil)
	}

	return helpers.Reply(c, http.StatusOK, nil, map[string]interface{}{
		"preference": preference,
	}, nil)
}

func (h *NotificationHandler) UpdatePreference(c echo.Context) error {
	ownerID, err := database.ParsePID(c.Param("id"))
	if err != nil {
		return helpers.Reply(c, http.StatusBadRequest, err, nil, nil)
	}

	var preference models.NotificationPreference
	if err := c.Bind(&preference); err != nil {
		return helpers.Reply(c, http.StatusBadRequest, err, nil, nil)
	}

	preference.OwnerID = ownerID
	if err := h.notificationUseCase.UpdatePreference(c.Request().Context(), &preference); err != nil {
		return helpers.Reply(c, http.StatusInternalServerError, err, nil, nil)
	}

	return helpers.Reply(c, http.StatusOK, nil, map[string]interface{}{
		"preference": preference,
	}, nil)
}
