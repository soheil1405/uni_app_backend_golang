package handlers

import (
	"net/http"
	"strconv"
	"uni_app/models"
	usecases "uni_app/pkg/place_type/usecase"

	"github.com/labstack/echo/v4"
)

type PlaceTypeHandler struct {
	usecase usecases.PlaceTypeUsecase
}

func NewPlaceTypeHandler(usecase usecases.PlaceTypeUsecase, e echo.Group) {
	placeTypeHandler := &PlaceTypeHandler{usecase}
	e.POST("/place_types", placeTypeHandler.CreatePlaceType)
	e.GET("/place_types/:id", placeTypeHandler.GetPlaceTypeByID)
	e.PUT("/place_types/:id", placeTypeHandler.UpdatePlaceType)
	e.DELETE("/place_types/:id", placeTypeHandler.DeletePlaceType)
	e.GET("/place_types", placeTypeHandler.GetAllPlaceTypes)

}

func (h *PlaceTypeHandler) CreatePlaceType(c echo.Context) error {
	var placeType models.PlaceType
	if err := c.Bind(&placeType); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}
	if err := h.usecase.CreatePlaceType(&placeType); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusCreated, placeType)
}

func (h *PlaceTypeHandler) GetPlaceTypeByID(c echo.Context) error {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid id"})
	}
	placeType, err := h.usecase.GetPlaceTypeByID(uint(id))
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, placeType)
}

func (h *PlaceTypeHandler) UpdatePlaceType(c echo.Context) error {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid id"})
	}
	var placeType models.PlaceType
	if err := c.Bind(&placeType); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}
	placeType.ID = uint(id)
	if err := h.usecase.UpdatePlaceType(&placeType); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, placeType)
}

func (h *PlaceTypeHandler) DeletePlaceType(c echo.Context) error {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid id"})
	}
	if err := h.usecase.DeletePlaceType(uint(id)); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, map[string]string{"message": "PlaceType deleted"})
}

func (h *PlaceTypeHandler) GetAllPlaceTypes(c echo.Context) error {
	placeTypes, err := h.usecase.GetAllPlaceTypes()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, placeTypes)
}
