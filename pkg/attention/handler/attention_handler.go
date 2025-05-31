package handler

import (
	"net/http"
	"uni_app/database"
	"uni_app/models"
	usecase "uni_app/pkg/attention/usecase"
	"uni_app/utils/ctxHelper"
	"uni_app/utils/helpers"

	"github.com/labstack/echo/v4"
)

type AttentionHandler struct {
	usecase usecase.AttentionUseCase
}

func NewAttentionHandler(usecase usecase.AttentionUseCase, e echo.Group) {
	attentionHandler := &AttentionHandler{usecase}

	attentionsRouteGroup := e.Group("/attentions")
	attentionsRouteGroup.POST("", attentionHandler.CreateAttention)
	attentionsRouteGroup.GET("/:id", attentionHandler.GetAttentionByID)
	attentionsRouteGroup.PUT("/:id", attentionHandler.UpdateAttention)
	attentionsRouteGroup.DELETE("/:id", attentionHandler.DeleteAttention)
	attentionsRouteGroup.GET("", attentionHandler.GetAllAttentions)
	attentionsRouteGroup.POST("/:id/read", attentionHandler.MarkAsRead)
	attentionsRouteGroup.POST("/:id/archive", attentionHandler.ArchiveAttention)
	attentionsRouteGroup.GET("/recipient/:id", attentionHandler.GetAttentionsByRecipient)
}

func (h *AttentionHandler) CreateAttention(c echo.Context) error {
	var attention models.Attention
	if err := c.Bind(&attention); err != nil {
		return helpers.Reply(c, http.StatusBadRequest, err, nil, nil)
	}
	if err := h.usecase.CreateAttention(c.Request().Context(), &attention); err != nil {
		return helpers.Reply(c, http.StatusInternalServerError, err, nil, nil)
	}
	return helpers.Reply(c, http.StatusCreated, nil, map[string]interface{}{"attention": attention}, nil)
}

func (h *AttentionHandler) GetAttentionByID(c echo.Context) error {
	var (
		ID  database.PID
		err error
	)
	if ID, err = ctxHelper.GetIDFromContxt(c); err != nil {
		return helpers.Reply(c, http.StatusBadRequest, err, nil, nil)
	}
	attention, err := h.usecase.GetAttentionByID(c.Request().Context(), ID)
	if err != nil {
		return helpers.Reply(c, http.StatusNotFound, err, nil, nil)
	}
	return helpers.Reply(c, http.StatusOK, nil, map[string]interface{}{"attention": attention}, nil)
}

func (h *AttentionHandler) UpdateAttention(c echo.Context) error {
	var (
		attention models.Attention
		err       error
	)
	if attention.ID, err = ctxHelper.GetIDFromContxt(c); err != nil {
		return helpers.Reply(c, http.StatusBadRequest, err, nil, nil)
	}
	if err := c.Bind(&attention); err != nil {
		return helpers.Reply(c, http.StatusBadRequest, err, nil, nil)
	}
	if err := h.usecase.UpdateAttention(c.Request().Context(), &attention); err != nil {
		return helpers.Reply(c, http.StatusInternalServerError, err, nil, nil)
	}
	return helpers.Reply(c, http.StatusOK, nil, map[string]interface{}{"attention": attention}, nil)
}

func (h *AttentionHandler) DeleteAttention(c echo.Context) error {
	ID, err := ctxHelper.GetIDFromContxt(c)
	if err != nil {
		return helpers.Reply(c, http.StatusBadRequest, err, nil, nil)
	}
	if err := h.usecase.DeleteAttention(c.Request().Context(), ID); err != nil {
		return helpers.Reply(c, http.StatusInternalServerError, err, nil, nil)
	}
	return helpers.Reply(c, http.StatusOK, nil, map[string]interface{}{"message": "Attention deleted"}, nil)
}

func (h *AttentionHandler) GetAllAttentions(c echo.Context) error {
	var request models.FetchAttentionRequest
	if err := c.Bind(&request); err != nil {
		return helpers.Reply(c, http.StatusBadRequest, err, nil, nil)
	}
	attentions, paginate, err := h.usecase.GetAllAttentions(c.Request().Context(), request)
	if err != nil {
		return helpers.Reply(c, http.StatusInternalServerError, err, nil, nil)
	}
	return helpers.Reply(c, http.StatusOK, nil, map[string]interface{}{"attentions": attentions}, paginate)
}

func (h *AttentionHandler) MarkAsRead(c echo.Context) error {
	ID, err := ctxHelper.GetIDFromContxt(c)
	if err != nil {
		return helpers.Reply(c, http.StatusBadRequest, err, nil, nil)
	}
	if err := h.usecase.MarkAsRead(c.Request().Context(), ID); err != nil {
		return helpers.Reply(c, http.StatusInternalServerError, err, nil, nil)
	}
	attention, err := h.usecase.GetAttentionByID(c.Request().Context(), ID)
	if err != nil {
		return helpers.Reply(c, http.StatusNotFound, err, nil, nil)
	}
	return helpers.Reply(c, http.StatusOK, nil, map[string]interface{}{"attention": attention}, nil)
}

func (h *AttentionHandler) ArchiveAttention(c echo.Context) error {
	ID, err := ctxHelper.GetIDFromContxt(c)
	if err != nil {
		return helpers.Reply(c, http.StatusBadRequest, err, nil, nil)
	}
	if err := h.usecase.MarkAsArchived(c.Request().Context(), ID); err != nil {
		return helpers.Reply(c, http.StatusInternalServerError, err, nil, nil)
	}
	attention, err := h.usecase.GetAttentionByID(c.Request().Context(), ID)
	if err != nil {
		return helpers.Reply(c, http.StatusNotFound, err, nil, nil)
	}
	return helpers.Reply(c, http.StatusOK, nil, map[string]interface{}{"attention": attention}, nil)
}

func (h *AttentionHandler) GetAttentionsByRecipient(c echo.Context) error {
	recipientID, err := ctxHelper.GetIDFromContxt(c)
	if err != nil {
		return helpers.Reply(c, http.StatusBadRequest, err, nil, nil)
	}
	recipientType := c.QueryParam("type")
	if recipientType == "" {
		return helpers.Reply(c, http.StatusBadRequest, echo.NewHTTPError(http.StatusBadRequest, "recipient type is required"), nil, nil)
	}
	attentions, err := h.usecase.GetAttentionsByRecipient(c.Request().Context(), recipientID, recipientType)
	if err != nil {
		return helpers.Reply(c, http.StatusInternalServerError, err, nil, nil)
	}
	return helpers.Reply(c, http.StatusOK, nil, map[string]interface{}{"attentions": attentions}, nil)
}
