package handler

import (
	"net/http"
	"uni_app/database"
	"uni_app/models"
	usecase "uni_app/pkg/city/usecase"
	"uni_app/utils/ctxHelper"
	"uni_app/utils/helpers"

	"github.com/labstack/echo/v4"
)

type CityHandler struct {
	usecase usecase.CityUsecase
}

func NewCityHandler(usecase usecase.CityUsecase, e echo.Group) {
	cityHandler := &CityHandler{usecase}

	citiesRouteGroup := e.Group("/cities")
	citiesRouteGroup.POST("", cityHandler.CreateCity)
	citiesRouteGroup.GET("/:id", cityHandler.GetCityByID)
	citiesRouteGroup.PUT("/:id", cityHandler.UpdateCity)
	citiesRouteGroup.DELETE("/:id", cityHandler.DeleteCity)
	citiesRouteGroup.GET("", cityHandler.GetAllCities)
}

func (h *CityHandler) CreateCity(c echo.Context) error {
	var city models.City
	if err := c.Bind(&city); err != nil {
		return helpers.Reply(c, http.StatusBadRequest, err, nil, nil)
	}
	if err := h.usecase.CreateCity(&city); err != nil {
		return helpers.Reply(c, http.StatusInternalServerError, err, nil, nil)
	}
	return helpers.Reply(c, http.StatusCreated, nil, map[string]interface{}{"city": city}, nil)
}

func (h *CityHandler) GetCityByID(c echo.Context) error {
	var (
		ID  database.PID
		err error
	)
	if ID, err = ctxHelper.GetIDFromContxt(c); err != nil {
		return helpers.Reply(c, http.StatusBadRequest, err, nil, nil)
	}
	city, err := h.usecase.GetCityByID(c, ID, false)
	if err != nil {
		return helpers.Reply(c, http.StatusNotFound, err, nil, nil)
	}
	return helpers.Reply(c, http.StatusOK, nil, map[string]interface{}{"city": city}, nil)
}

func (h *CityHandler) UpdateCity(c echo.Context) (err error) {
	var city models.City
	if city.ID, err = ctxHelper.GetIDFromContxt(c); err != nil {
		return helpers.Reply(c, http.StatusBadRequest, err, nil, nil)
	}
	if err := h.usecase.UpdateCity(&city); err != nil {
		return helpers.Reply(c, http.StatusInternalServerError, err, nil, nil)
	}
	return helpers.Reply(c, http.StatusOK, nil, map[string]interface{}{"city": city}, nil)
}

func (h *CityHandler) DeleteCity(c echo.Context) error {
	ID, err := ctxHelper.GetIDFromContxt(c)
	if err != nil {
		return helpers.Reply(c, http.StatusBadRequest, err, nil, nil)
	}

	if err := h.usecase.DeleteCity(ID); err != nil {
		return helpers.Reply(c, http.StatusInternalServerError, err, nil, nil)
	}
	return helpers.Reply(c, http.StatusOK, nil, map[string]interface{}{"message": "City deleted"}, nil)
}

func (h *CityHandler) GetAllCities(c echo.Context) error {
	var request models.FetchCityRequest
	if err := c.Bind(&request); err != nil {
		return helpers.Reply(c, http.StatusBadRequest, err, nil, nil)
	}
	cities, paginate, err := h.usecase.GetAllCities(c, request)
	if err != nil {
		return helpers.Reply(c, http.StatusInternalServerError, err, nil, nil)
	}
	return helpers.Reply(c, http.StatusOK, nil, map[string]interface{}{"cities": cities}, paginate)
}
