package handlers

import (
	"net/http"
	"uni_app/database"
	"uni_app/models"
	usecases "uni_app/pkg/city/usecase"
	"uni_app/utils/ctxHelper"

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
	cities, err := h.usecase.GetAllCities()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, cities)
}
