package handler

import (
	"net/http"
	"uni_app/database"
	"uni_app/models"
	usecase "uni_app/pkg/degree/usecase"
	"uni_app/utils/ctxHelper"

	"github.com/labstack/echo/v4"
)

type DegreeHandler struct {
	usecase usecase.DegreeUsecase
}

func NewDegreeHandler(usecase usecase.DegreeUsecase, e echo.Group) {
	degreeHandler := &DegreeHandler{usecase}

	degreesRouteGroup := e.Group("/degrees")
	degreesRouteGroup.POST("", degreeHandler.CreateDegree)
	degreesRouteGroup.GET("/:id", degreeHandler.GetDegreeByID)
	degreesRouteGroup.PUT("/:id", degreeHandler.UpdateDegree)
	degreesRouteGroup.DELETE("/:id", degreeHandler.DeleteDegree)
	degreesRouteGroup.GET("", degreeHandler.GetAllDegrees)

}

func (h *DegreeHandler) CreateDegree(c echo.Context) error {
	var degree models.DegreeLevel
	if err := c.Bind(&degree); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}
	if err := h.usecase.CreateDegree(&degree); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusCreated, degree)
}

func (h *DegreeHandler) GetDegreeByID(c echo.Context) error {
	var (
		ID  database.PID
		err error
	)
	if ID, err = ctxHelper.GetIDFromContxt(c); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}
	degree, err := h.usecase.GetDegreeByID(c, ID, false)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, degree)
}

func (h *DegreeHandler) UpdateDegree(c echo.Context) (err error) {
	var degree models.DegreeLevel
	if degree.ID, err = ctxHelper.GetIDFromContxt(c); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}
	if err := h.usecase.UpdateDegree(&degree); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, degree)
}

func (h *DegreeHandler) DeleteDegree(c echo.Context) error {
	ID, err := ctxHelper.GetIDFromContxt(c)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	if err := h.usecase.DeleteDegree(ID); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, map[string]string{"message": "Degree deleted"})
}

func (h *DegreeHandler) GetAllDegrees(c echo.Context) error {
	degrees, err := h.usecase.GetAllDegrees()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, degrees)
} 