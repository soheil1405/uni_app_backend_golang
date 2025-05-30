package handler

import (
	"net/http"
	"uni_app/database"
	"uni_app/models"
	usecase "uni_app/pkg/place/usecase"
	"uni_app/utils/ctxHelper"
	"uni_app/utils/helpers"

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
		return helpers.Reply(c, http.StatusBadRequest, err, nil, nil)
	}
	if err := h.usecase.CreatePlace(&place); err != nil {
		return helpers.Reply(c, http.StatusInternalServerError, err, nil, nil)
	}
	return helpers.Reply(c, http.StatusCreated, nil, map[string]interface{}{"place": place}, nil)
}

func (h *PlaceHandler) GetPlaceByID(c echo.Context) error {
	var (
		ID  database.PID
		err error
	)
	if ID, err = ctxHelper.GetIDFromContxt(c); err != nil {
		return helpers.Reply(c, http.StatusBadRequest, err, nil, nil)
	}
	place, err := h.usecase.GetPlaceByID(c, ID, false)
	if err != nil {
		return helpers.Reply(c, http.StatusNotFound, err, nil, nil)
	}
	return helpers.Reply(c, http.StatusOK, nil, map[string]interface{}{"place": place}, nil)
}

func (h *PlaceHandler) UpdatePlace(c echo.Context) (err error) {
	var place models.Place
	if place.ID, err = ctxHelper.GetIDFromContxt(c); err != nil {
		return helpers.Reply(c, http.StatusBadRequest, err, nil, nil)
	}
	if err := h.usecase.UpdatePlace(&place); err != nil {
		return helpers.Reply(c, http.StatusInternalServerError, err, nil, nil)
	}
	return helpers.Reply(c, http.StatusOK, nil, map[string]interface{}{"place": place}, nil)
}

func (h *PlaceHandler) DeletePlace(c echo.Context) error {
	ID, err := ctxHelper.GetIDFromContxt(c)
	if err != nil {
		return helpers.Reply(c, http.StatusBadRequest, err, nil, nil)
	}

	if err := h.usecase.DeletePlace(ID); err != nil {
		return helpers.Reply(c, http.StatusInternalServerError, err, nil, nil)
	}
	return helpers.Reply(c, http.StatusOK, nil, map[string]interface{}{"message": "Place deleted"}, nil)
}

func (h *PlaceHandler) GetAllPlaces(c echo.Context) error {
	var request models.FetchPlaceRequest
	if err := c.Bind(&request); err != nil {
		return helpers.Reply(c, http.StatusBadRequest, err, nil, nil)
	}
	places, paginate, err := h.usecase.GetAllPlaces(c, request)
	if err != nil {
		return helpers.Reply(c, http.StatusInternalServerError, err, nil, nil)
	}
	return helpers.Reply(c, http.StatusOK, nil, map[string]interface{}{"places": places}, paginate)
}
