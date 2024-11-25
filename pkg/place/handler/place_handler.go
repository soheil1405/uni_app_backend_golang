package handlers

import (
	"net/http"
	"strconv"
	"uni_app/models"
	usecases "uni_app/pkg/place/usecase"

	"github.com/labstack/echo/v4"
)

type PlaceHandler struct {
	usecase usecases.PlaceUsecase
}

func NewPlaceHandler(usecase usecases.PlaceUsecase, e echo.Group) {
	placeHandler := &PlaceHandler{usecase}
	e.POST("/places", placeHandler.CreatePlace)
	e.GET("/places/:id", placeHandler.GetPlaceByID)
	e.PUT("/places/:id", placeHandler.UpdatePlace)
	e.DELETE("/places/:id", placeHandler.DeletePlace)
	e.GET("/places", placeHandler.GetAllPlaces)

}

func (h *PlaceHandler) CreatePlace(c echo.Context) error {
	var place models.Place
	if err := c.Bind(&place); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}
	if err := h.usecase.CreatePlace(&place); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusCreated, place)
}

func (h *PlaceHandler) GetPlaceByID(c echo.Context) error {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid id"})
	}
	place, err := h.usecase.GetPlaceByID(uint(id))
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, place)
}

func (h *PlaceHandler) UpdatePlace(c echo.Context) error {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid id"})
	}
	var place models.Place
	if err := c.Bind(&place); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}
	place.ID = uint(id)
	if err := h.usecase.UpdatePlace(&place); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, place)
}

func (h *PlaceHandler) DeletePlace(c echo.Context) error {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid id"})
	}
	if err := h.usecase.DeletePlace(uint(id)); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, map[string]string{"message": "Place deleted"})
}

func (h *PlaceHandler) GetAllPlaces(c echo.Context) error {
	places, err := h.usecase.GetAllPlaces()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, places)
}
