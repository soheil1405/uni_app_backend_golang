package handler

import (
	"net/http"
	"strconv"
	"uni_app/database"
	"uni_app/models"
	"uni_app/pkg/news/usecase"
	"uni_app/utils/ctxHelper"
	"uni_app/utils/helpers"

	"github.com/labstack/echo/v4"
)

type NewsHandler struct {
	newsUseCase usecase.NewsUseCase
}

func NewNewsHandler(newsUseCase usecase.NewsUseCase, e echo.Group) {
	handler := &NewsHandler{
		newsUseCase: newsUseCase,
	}

	// News CRUD endpoints
	e.POST("/news", handler.CreateNews)
	e.GET("/news/:id", handler.GetNewsByID)
	e.PUT("/news/:id", handler.UpdateNews)
	e.DELETE("/news/:id", handler.DeleteNews)

	// News listing and filtering endpoints
	e.GET("/news", handler.GetAllNews)
	e.GET("/news/owner/:ownerType/:ownerID", handler.GetNewsByOwner)
	e.GET("/news/published", handler.GetPublishedNews)

	// News status management endpoints
	e.PUT("/news/:id/publish", handler.PublishNews)
	e.PUT("/news/:id/archive", handler.ArchiveNews)

	// Notification management endpoints
	e.GET("/news/notifications/pending", handler.GetPendingNotifications)
	e.POST("/news/notifications/process", handler.ProcessNotifications)
}

// CreateNews handles the creation of a new news item
func (h *NewsHandler) CreateNews(c echo.Context) error {
	var news models.News
	if err := c.Bind(&news); err != nil {
		return helpers.Reply(c, http.StatusBadRequest, err, nil, nil)
	}

	// Set author ID from context (assuming it's set by auth middleware)
	if authorID, ok := c.Get("user_id").(uint); ok {
		news.AuthorID = database.PID(authorID)
	}

	if err := h.newsUseCase.CreateNews(c.Request().Context(), &news); err != nil {
		return helpers.Reply(c, http.StatusInternalServerError, err, nil, nil)
	}

	return helpers.Reply(c, http.StatusCreated, nil, map[string]interface{}{"news": news}, nil)
}

// GetNewsByID retrieves a specific news item by ID
func (h *NewsHandler) GetNewsByID(c echo.Context) error {
	id, err := ctxHelper.GetIDFromContxt(c)
	if err != nil {
		return helpers.Reply(c, http.StatusBadRequest, err, nil, nil)
	}

	news, err := h.newsUseCase.GetNewsByID(c.Request().Context(), database.PID(id))
	if err != nil {
		return helpers.Reply(c, http.StatusNotFound, err, nil, nil)
	}

	return helpers.Reply(c, http.StatusOK, nil, map[string]interface{}{"news": news}, nil)
}

// UpdateNews updates an existing news item
func (h *NewsHandler) UpdateNews(c echo.Context) error {
	id, err := ctxHelper.GetIDFromContxt(c)
	if err != nil {
		return helpers.Reply(c, http.StatusBadRequest, err, nil, nil)
	}

	var news models.News
	if err := c.Bind(&news); err != nil {
		return helpers.Reply(c, http.StatusBadRequest, err, nil, nil)
	}

	news.ID = database.PID(id)
	if err := h.newsUseCase.UpdateNews(c.Request().Context(), &news); err != nil {
		return helpers.Reply(c, http.StatusInternalServerError, err, nil, nil)
	}

	return helpers.Reply(c, http.StatusOK, nil, map[string]interface{}{"news": news}, nil)
}

// DeleteNews deletes a news item
func (h *NewsHandler) DeleteNews(c echo.Context) error {
	id, err := ctxHelper.GetIDFromContxt(c)
	if err != nil {
		return helpers.Reply(c, http.StatusBadRequest, err, nil, nil)
	}

	if err := h.newsUseCase.DeleteNews(c.Request().Context(), database.PID(id)); err != nil {
		return helpers.Reply(c, http.StatusInternalServerError, err, nil, nil)
	}

	return helpers.Reply(c, http.StatusOK, nil, map[string]interface{}{"message": "News deleted successfully"}, nil)
}

// GetAllNews retrieves all news items with filtering and pagination
func (h *NewsHandler) GetAllNews(c echo.Context) error {
	var request models.FetchNewsRequest
	if err := c.Bind(&request); err != nil {
		return helpers.Reply(c, http.StatusBadRequest, err, nil, nil)
	}

	news, paginate, err := h.newsUseCase.GetAllNews(c.Request().Context(), request)
	if err != nil {
		return helpers.Reply(c, http.StatusInternalServerError, err, nil, nil)
	}

	return helpers.Reply(c, http.StatusOK, nil, map[string]interface{}{"news": news}, paginate)
}

// GetNewsByOwner retrieves news items for a specific owner
func (h *NewsHandler) GetNewsByOwner(c echo.Context) error {
	ownerType := c.Param("ownerType")
	ownerID, err := strconv.ParseUint(c.Param("ownerID"), 10, 64)
	if err != nil {
		return helpers.Reply(c, http.StatusBadRequest, err, nil, nil)
	}

	news, err := h.newsUseCase.GetNewsByOwner(c.Request().Context(), database.PID(ownerID), ownerType)
	if err != nil {
		return helpers.Reply(c, http.StatusInternalServerError, err, nil, nil)
	}

	return helpers.Reply(c, http.StatusOK, nil, map[string]interface{}{"news": news}, nil)
}

// GetPublishedNews retrieves all published news items
func (h *NewsHandler) GetPublishedNews(c echo.Context) error {
	news, err := h.newsUseCase.GetPublishedNews(c.Request().Context())
	if err != nil {
		return helpers.Reply(c, http.StatusInternalServerError, err, nil, nil)
	}

	return helpers.Reply(c, http.StatusOK, nil, map[string]interface{}{"news": news}, nil)
}

// PublishNews publishes a news item
func (h *NewsHandler) PublishNews(c echo.Context) error {
	id, err := ctxHelper.GetIDFromContxt(c)
	if err != nil {
		return helpers.Reply(c, http.StatusBadRequest, err, nil, nil)
	}

	if err := h.newsUseCase.PublishNews(c.Request().Context(), database.PID(id)); err != nil {
		return helpers.Reply(c, http.StatusInternalServerError, err, nil, nil)
	}

	return helpers.Reply(c, http.StatusOK, nil, map[string]interface{}{"message": "News published successfully"}, nil)
}

// ArchiveNews archives a news item
func (h *NewsHandler) ArchiveNews(c echo.Context) error {
	id, err := ctxHelper.GetIDFromContxt(c)
	if err != nil {
		return helpers.Reply(c, http.StatusBadRequest, err, nil, nil)
	}

	if err := h.newsUseCase.ArchiveNews(c.Request().Context(), database.PID(id)); err != nil {
		return helpers.Reply(c, http.StatusInternalServerError, err, nil, nil)
	}

	return helpers.Reply(c, http.StatusOK, nil, map[string]interface{}{"message": "News archived successfully"}, nil)
}

// GetPendingNotifications retrieves all pending notifications
func (h *NewsHandler) GetPendingNotifications(c echo.Context) error {
	news, err := h.newsUseCase.GetPendingNotifications(c.Request().Context())
	if err != nil {
		return helpers.Reply(c, http.StatusInternalServerError, err, nil, nil)
	}

	return helpers.Reply(c, http.StatusOK, nil, map[string]interface{}{"news": news}, nil)
}

// ProcessNotifications processes all pending notifications
func (h *NewsHandler) ProcessNotifications(c echo.Context) error {
	if err := h.newsUseCase.ProcessNotifications(c.Request().Context()); err != nil {
		return helpers.Reply(c, http.StatusInternalServerError, err, nil, nil)
	}

	return helpers.Reply(c, http.StatusOK, nil, map[string]interface{}{"message": "Notifications processed successfully"}, nil)
}
