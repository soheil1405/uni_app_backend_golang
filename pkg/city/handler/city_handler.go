package handler

import (
	"net/http"
	"uni_app/database"
	"uni_app/models"
	usecase "uni_app/pkg/city/usecase"
	"uni_app/utils/ctxHelper"

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
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}
	if err := h.usecase.CreateCity(&city); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusCreated, city)
}

func (h *CityHandler) GetCityByID(c echo.Context) error {
	var (
		ID  database.PID
		err error
	)
	if ID, err = ctxHelper.GetIDFromContxt(c); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}
	city, err := h.usecase.GetCityByID(c, ID, false)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, city)
}

func (h *CityHandler) UpdateCity(c echo.Context) (err error) {
	var city models.City
	if city.ID, err = ctxHelper.GetIDFromContxt(c); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}
	if err := h.usecase.UpdateCity(&city); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, city)
}

func (h *CityHandler) DeleteCity(c echo.Context) error {
	ID, err := ctxHelper.GetIDFromContxt(c)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	if err := h.usecase.DeleteCity(ID); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, map[string]string{"message": "City deleted"})
}

func (h *CityHandler) GetAllCities(c echo.Context) error {
	var request models.FetchCityRequest
	if err := c.Bind(&request); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}
	cities, paginate, err := h.usecase.GetAllCities(c, request)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"cities": cities,
		"meta":   paginate,
	})
}
