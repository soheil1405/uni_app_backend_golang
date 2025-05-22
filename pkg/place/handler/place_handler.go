package handler

import (
	"net/http"
	"uni_app/database"
	"uni_app/models"
	usecase "uni_app/pkg/place/usecase"
	"uni_app/utils/ctxHelper"

	"github.com/labstack/echo/v4"
)

type PlaceHandler struct {
	usecase usecase.PlaceUsecase
}

func NewPlaceHandler(usecase usecase.PlaceUsecase, e echo.Group) {
	placeHandler := &PlaceHandler{usecase}

	placesRouteGroup := e.Group("/places")
	placesRouteGroup.POST("", placeHandler.CreatePlace)
	placesRouteGroup.GET("/:id", placeHandler.GetPlaceByID)
	placesRouteGroup.PUT("/:id", placeHandler.UpdatePlace)
	placesRouteGroup.DELETE("/:id", placeHandler.DeletePlace)
	placesRouteGroup.GET("", placeHandler.GetAllPlaces)
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
	var (
		ID  database.PID
		err error
	)
	if ID, err = ctxHelper.GetIDFromContxt(c); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}
	place, err := h.usecase.GetPlaceByID(c, ID, false)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, place)
}

func (h *PlaceHandler) UpdatePlace(c echo.Context) (err error) {
	var place models.Place
	if place.ID, err = ctxHelper.GetIDFromContxt(c); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}
	if err := h.usecase.UpdatePlace(&place); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, place)
}

func (h *PlaceHandler) DeletePlace(c echo.Context) error {
	ID, err := ctxHelper.GetIDFromContxt(c)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	if err := h.usecase.DeletePlace(ID); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, map[string]string{"message": "Place deleted"})
}

func (h *PlaceHandler) GetAllPlaces(c echo.Context) error {
	var request models.FetchPlaceRequest
	if err := c.Bind(&request); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}
	places, paginate, err := h.usecase.GetAllPlaces(c, request)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"places": places,
		"meta":   paginate,
	})
}
