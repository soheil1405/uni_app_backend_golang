package handlers

import (
	"net/http"
	"strconv"
	"uni_app/models"
	usecases "uni_app/pkg/city/usecase"

	"github.com/labstack/echo/v4"
)

type CityHandler struct {
	usecase usecases.CityUsecase
}

func NewCityHandler(usecase usecases.CityUsecase, e echo.Group) {
	cityHandler := &CityHandler{usecase}

	e.POST("/cities", cityHandler.CreateCity)
	e.GET("/cities/:id", cityHandler.GetCityByID)
	e.PUT("/cities/:id", cityHandler.UpdateCity)
	e.DELETE("/cities/:id", cityHandler.DeleteCity)
	e.GET("/cities", cityHandler.GetAllCities)

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
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid id"})
	}
	city, err := h.usecase.GetCityByID(uint(id))
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, city)
}

func (h *CityHandler) UpdateCity(c echo.Context) error {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid id"})
	}
	var city models.City
	if err := c.Bind(&city); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}
	city.ID = uint(id)
	if err := h.usecase.UpdateCity(&city); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, city)
}

func (h *CityHandler) DeleteCity(c echo.Context) error {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid id"})
	}
	if err := h.usecase.DeleteCity(uint(id)); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, map[string]string{"message": "City deleted"})
}

func (h *CityHandler) GetAllCities(c echo.Context) error {
	cities, err := h.usecase.GetAllCities()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, cities)
}
