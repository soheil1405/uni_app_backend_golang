package handler

import (
	"net/http"
	"strconv"

	"uni_app/models"
	"uni_app/pkg/social_media_link/usecase"
	"uni_app/utils/ctxHelper"
	"uni_app/utils/helpers"

	"github.com/labstack/echo/v4"
)

type SocialMediaLinkHandler struct {
	socialMediaLinkUsecase usecase.SocialMediaLinkUsecase
}

func NewSocialMediaLinkHandler(socialMediaLinkUsecase usecase.SocialMediaLinkUsecase, e echo.Group) {
	socialMediaLinkHandler := &SocialMediaLinkHandler{socialMediaLinkUsecase}

	socialMediaLinkRouteGroup := e.Group("/social-media-links")
	socialMediaLinkRouteGroup.POST("", socialMediaLinkHandler.CreateLink)
	socialMediaLinkRouteGroup.GET("/:id", socialMediaLinkHandler.GetLinkByID)
	socialMediaLinkRouteGroup.PUT("/:id", socialMediaLinkHandler.UpdateLink)
	socialMediaLinkRouteGroup.DELETE("/:id", socialMediaLinkHandler.DeleteLink)
	socialMediaLinkRouteGroup.GET("/linkable", socialMediaLinkHandler.GetLinksByLinkable)
	socialMediaLinkRouteGroup.GET("/platform", socialMediaLinkHandler.GetLinkByPlatform)
}

func (h *SocialMediaLinkHandler) CreateLink(c echo.Context) error {
	var link models.SocialMediaLink
	if err := c.Bind(&link); err != nil {
		return helpers.Reply(c, http.StatusBadRequest, err, nil, nil)
	}

	if err := h.socialMediaLinkUsecase.CreateLink(c.Request().Context(), &link); err != nil {
		return helpers.Reply(c, http.StatusInternalServerError, err, nil, nil)
	}

	return helpers.Reply(c, http.StatusCreated, nil, map[string]interface{}{"link": link}, nil)
}

func (h *SocialMediaLinkHandler) GetLinkByID(c echo.Context) error {
	ID, err := ctxHelper.GetIDFromContxt(c)
	if err != nil {
		return helpers.Reply(c, http.StatusBadRequest, err, nil, nil)
	}

	link, err := h.socialMediaLinkUsecase.GetLinkByID(c.Request().Context(), uint(ID))
	if err != nil {
		return helpers.Reply(c, http.StatusNotFound, err, nil, nil)
	}

	return helpers.Reply(c, http.StatusOK, nil, map[string]interface{}{"link": link}, nil)
}

func (h *SocialMediaLinkHandler) UpdateLink(c echo.Context) error {
	var link models.SocialMediaLink
	if err := c.Bind(&link); err != nil {
		return helpers.Reply(c, http.StatusBadRequest, err, nil, nil)
	}

	ID, err := ctxHelper.GetIDFromContxt(c)
	if err != nil {
		return helpers.Reply(c, http.StatusBadRequest, err, nil, nil)
	}
	link.ID = uint(ID)

	if err := h.socialMediaLinkUsecase.UpdateLink(c.Request().Context(), &link); err != nil {
		return helpers.Reply(c, http.StatusInternalServerError, err, nil, nil)
	}

	return helpers.Reply(c, http.StatusOK, nil, map[string]interface{}{"link": link}, nil)
}

func (h *SocialMediaLinkHandler) DeleteLink(c echo.Context) error {
	ID, err := ctxHelper.GetIDFromContxt(c)
	if err != nil {
		return helpers.Reply(c, http.StatusBadRequest, err, nil, nil)
	}

	if err := h.socialMediaLinkUsecase.DeleteLink(c.Request().Context(), uint(ID)); err != nil {
		return helpers.Reply(c, http.StatusInternalServerError, err, nil, nil)
	}

	return helpers.Reply(c, http.StatusOK, nil, map[string]interface{}{"message": "Social media link deleted successfully"}, nil)
}

func (h *SocialMediaLinkHandler) GetLinksByLinkable(c echo.Context) error {
	linkableID, err := strconv.ParseUint(c.QueryParam("linkable_id"), 10, 32)
	if err != nil {
		return helpers.Reply(c, http.StatusBadRequest, err, nil, nil)
	}

	linkableType := c.QueryParam("linkable_type")
	if linkableType == "" {
		return helpers.Reply(c, http.StatusBadRequest, nil, nil, map[string]string{"error": "Linkable type is required"})
	}

	links, err := h.socialMediaLinkUsecase.GetLinksByLinkable(c.Request().Context(), uint(linkableID), linkableType)
	if err != nil {
		return helpers.Reply(c, http.StatusInternalServerError, err, nil, nil)
	}

	return helpers.Reply(c, http.StatusOK, nil, map[string]interface{}{"links": links}, nil)
}

func (h *SocialMediaLinkHandler) GetLinkByPlatform(c echo.Context) error {
	linkableID, err := strconv.ParseUint(c.QueryParam("linkable_id"), 10, 32)
	if err != nil {
		return helpers.Reply(c, http.StatusBadRequest, err, nil, nil)
	}

	linkableType := c.QueryParam("linkable_type")
	if linkableType == "" {
		return helpers.Reply(c, http.StatusBadRequest, nil, nil, map[string]string{"error": "Linkable type is required"})
	}

	platform := c.QueryParam("platform")
	if platform == "" {
		return helpers.Reply(c, http.StatusBadRequest, nil, nil, map[string]string{"error": "Platform is required"})
	}

	link, err := h.socialMediaLinkUsecase.GetLinkByPlatform(c.Request().Context(), uint(linkableID), linkableType, platform)
	if err != nil {
		return helpers.Reply(c, http.StatusInternalServerError, err, nil, nil)
	}

	return helpers.Reply(c, http.StatusOK, nil, map[string]interface{}{"link": link}, nil)
}
