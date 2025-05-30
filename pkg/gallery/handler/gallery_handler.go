package handler

import (
	"net/http"
	"strconv"

	"uni_app/models"
	"uni_app/pkg/gallery/usecase"
	"uni_app/utils/ctxHelper"
	"uni_app/utils/helpers"

	"github.com/labstack/echo/v4"
)

type GalleryHandler struct {
	galleryUsecase usecase.GalleryUsecase
}

func NewGalleryHandler(galleryUsecase usecase.GalleryUsecase, e echo.Group) {
	galleryHandler := &GalleryHandler{galleryUsecase}

	galleryRouteGroup := e.Group("/gallery")
	galleryRouteGroup.POST("", galleryHandler.CreateGallery)
	galleryRouteGroup.GET("/:id", galleryHandler.GetGalleryByID)
	galleryRouteGroup.PUT("/:id", galleryHandler.UpdateGallery)
	galleryRouteGroup.DELETE("/:id", galleryHandler.DeleteGallery)
	galleryRouteGroup.GET("/imageable", galleryHandler.GetGalleriesByImageable)
	galleryRouteGroup.PUT("/:id/main", galleryHandler.SetMainImage)
}

func (h *GalleryHandler) CreateGallery(c echo.Context) error {
	var gallery models.Gallery
	if err := c.Bind(&gallery); err != nil {
		return helpers.Reply(c, http.StatusBadRequest, err, nil, nil)
	}

	if err := h.galleryUsecase.CreateGallery(c.Request().Context(), &gallery); err != nil {
		return helpers.Reply(c, http.StatusInternalServerError, err, nil, nil)
	}

	return helpers.Reply(c, http.StatusCreated, nil, map[string]interface{}{"gallery": gallery}, nil)
}

func (h *GalleryHandler) GetGalleryByID(c echo.Context) error {
	ID, err := ctxHelper.GetIDFromContxt(c)
	if err != nil {
		return helpers.Reply(c, http.StatusBadRequest, err, nil, nil)
	}

	gallery, err := h.galleryUsecase.GetGalleryByID(c.Request().Context(), uint(ID))
	if err != nil {
		return helpers.Reply(c, http.StatusNotFound, err, nil, nil)
	}

	return helpers.Reply(c, http.StatusOK, nil, map[string]interface{}{"gallery": gallery}, nil)
}

func (h *GalleryHandler) UpdateGallery(c echo.Context) error {
	var gallery models.Gallery
	if err := c.Bind(&gallery); err != nil {
		return helpers.Reply(c, http.StatusBadRequest, err, nil, nil)
	}

	ID, err := ctxHelper.GetIDFromContxt(c)
	if err != nil {
		return helpers.Reply(c, http.StatusBadRequest, err, nil, nil)
	}
	gallery.ID = ID

	if err := h.galleryUsecase.UpdateGallery(c.Request().Context(), &gallery); err != nil {
		return helpers.Reply(c, http.StatusInternalServerError, err, nil, nil)
	}

	return helpers.Reply(c, http.StatusOK, nil, map[string]interface{}{"gallery": gallery}, nil)
}

func (h *GalleryHandler) DeleteGallery(c echo.Context) error {
	ID, err := ctxHelper.GetIDFromContxt(c)
	if err != nil {
		return helpers.Reply(c, http.StatusBadRequest, err, nil, nil)
	}

	if err := h.galleryUsecase.DeleteGallery(c.Request().Context(), uint(ID)); err != nil {
		return helpers.Reply(c, http.StatusInternalServerError, err, nil, nil)
	}

	return helpers.Reply(c, http.StatusOK, nil, map[string]interface{}{"message": "Gallery deleted successfully"}, nil)
}

func (h *GalleryHandler) GetGalleriesByImageable(c echo.Context) error {
	imageableID, err := strconv.ParseUint(c.QueryParam("imageable_id"), 10, 32)
	if err != nil {
		return helpers.Reply(c, http.StatusBadRequest, err, nil, nil)
	}

	imageableType := c.QueryParam("imageable_type")
	if imageableType == "" {
		return helpers.Reply(c, http.StatusBadRequest, nil, nil, map[string]string{"error": "Imageable type is required"})
	}

	galleries, err := h.galleryUsecase.GetGalleriesByImageable(c.Request().Context(), uint(imageableID), imageableType)
	if err != nil {
		return helpers.Reply(c, http.StatusInternalServerError, err, nil, nil)
	}

	return helpers.Reply(c, http.StatusOK, nil, map[string]interface{}{"galleries": galleries}, nil)
}

func (h *GalleryHandler) SetMainImage(c echo.Context) error {
	ID, err := ctxHelper.GetIDFromContxt(c)
	if err != nil {
		return helpers.Reply(c, http.StatusBadRequest, err, nil, nil)
	}

	if err := h.galleryUsecase.SetMainImage(c.Request().Context(), uint(ID)); err != nil {
		return helpers.Reply(c, http.StatusInternalServerError, err, nil, nil)
	}

	return helpers.Reply(c, http.StatusOK, nil, map[string]interface{}{"message": "Main image updated successfully"}, nil)
}
