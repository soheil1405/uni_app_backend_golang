package handler

import (
	"net/http"
	"uni_app/database"
	"uni_app/models"
	usecase "uni_app/pkg/place_type/usecase"
	"uni_app/utils/ctxHelper"
	"uni_app/utils/helpers"

	"github.com/labstack/echo/v4"
)

type PlaceTypeHandler struct {
	usecase usecase.PlaceTypeUsecase
}

func NewPlaceTypeHandler(usecase usecase.PlaceTypeUsecase, e echo.Group) {
	placeTypeHandler := &PlaceTypeHandler{usecase}

	placeTypesRouteGroup := e.Group("/place-types")
	placeTypesRouteGroup.POST("", placeTypeHandler.CreatePlaceType)
	placeTypesRouteGroup.GET("/:id", placeTypeHandler.GetPlaceTypeByID)
	placeTypesRouteGroup.PUT("/:id", placeTypeHandler.UpdatePlaceType)
	placeTypesRouteGroup.DELETE("/:id", placeTypeHandler.DeletePlaceType)
	placeTypesRouteGroup.GET("", placeTypeHandler.GetAllPlaceTypes)
}

func (h *PlaceTypeHandler) CreatePlaceType(c echo.Context) error {
	var placeType models.PlaceType
	if err := c.Bind(&placeType); err != nil {
		return helpers.Reply(c, http.StatusBadRequest, err, nil, nil)
	}
	if err := h.usecase.CreatePlaceType(&placeType); err != nil {
		return helpers.Reply(c, http.StatusInternalServerError, err, nil, nil)
	}
	return helpers.Reply(c, http.StatusCreated, nil, map[string]interface{}{"place_type": placeType}, nil)
}

func (h *PlaceTypeHandler) GetPlaceTypeByID(c echo.Context) error {
	var (
		ID  database.PID
		err error
	)
	if ID, err = ctxHelper.GetIDFromContxt(c); err != nil {
		return helpers.Reply(c, http.StatusBadRequest, err, nil, nil)
	}
	placeType, err := h.usecase.GetPlaceTypeByID(c, ID, false)
	if err != nil {
		return helpers.Reply(c, http.StatusNotFound, err, nil, nil)
	}
	return helpers.Reply(c, http.StatusOK, nil, map[string]interface{}{"place_type": placeType}, nil)
}

func (h *PlaceTypeHandler) UpdatePlaceType(c echo.Context) (err error) {
	var placeType models.PlaceType
	if placeType.ID, err = ctxHelper.GetIDFromContxt(c); err != nil {
		return helpers.Reply(c, http.StatusBadRequest, err, nil, nil)
	}
	if err := h.usecase.UpdatePlaceType(&placeType); err != nil {
		return helpers.Reply(c, http.StatusInternalServerError, err, nil, nil)
	}
	return helpers.Reply(c, http.StatusOK, nil, map[string]interface{}{"place_type": placeType}, nil)
}

func (h *PlaceTypeHandler) DeletePlaceType(c echo.Context) error {
	ID, err := ctxHelper.GetIDFromContxt(c)
	if err != nil {
		return helpers.Reply(c, http.StatusBadRequest, err, nil, nil)
	}

	if err := h.usecase.DeletePlaceType(ID); err != nil {
		return helpers.Reply(c, http.StatusInternalServerError, err, nil, nil)
	}
	return helpers.Reply(c, http.StatusOK, nil, map[string]interface{}{"message": "Place type deleted"}, nil)
}

func (h *PlaceTypeHandler) GetAllPlaceTypes(c echo.Context) error {
	placeTypes, err := h.usecase.GetAllPlaceTypes()
	if err != nil {
		return helpers.Reply(c, http.StatusInternalServerError, err, nil, nil)
	}
	return helpers.Reply(c, http.StatusOK, nil, map[string]interface{}{"place_types": placeTypes}, nil)
}
