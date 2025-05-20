package handler

import (
	"net/http"
	"uni_app/database"
	"uni_app/models"
	usecase "uni_app/pkg/request/usecase"
	"uni_app/utils/ctxHelper"

	"github.com/labstack/echo/v4"
)

type RequestHandler struct {
	usecase usecase.RequestUsecase
}

func NewRequestHandler(usecase usecase.RequestUsecase, e echo.Group) {
	requestHandler := &RequestHandler{usecase}

	requestsRouteGroup := e.Group("/requests")
	requestsRouteGroup.POST("", requestHandler.CreateRequest)
	requestsRouteGroup.GET("/:id", requestHandler.GetRequestByID)
	requestsRouteGroup.PUT("/:id", requestHandler.UpdateRequest)
	requestsRouteGroup.DELETE("/:id", requestHandler.DeleteRequest)
	requestsRouteGroup.GET("", requestHandler.GetAllRequests)

}

func (h *RequestHandler) CreateRequest(c echo.Context) error {
	var request models.FetchRequest
	if err := c.Bind(&request); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}
	if err := h.usecase.CreateRequest(&request); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusCreated, request)
}

func (h *RequestHandler) GetRequestByID(c echo.Context) error {
	var (
		ID  database.PID
		err error
	)
	if ID, err = ctxHelper.GetIDFromContxt(c); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}
	request, err := h.usecase.GetRequestByID(c, ID, false)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, request)
}

func (h *RequestHandler) UpdateRequest(c echo.Context) (err error) {
	var request models.FetchRequest
	if request.ID, err = ctxHelper.GetIDFromContxt(c); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}
	if err := h.usecase.UpdateRequest(&request); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, request)
}

func (h *RequestHandler) DeleteRequest(c echo.Context) error {
	ID, err := ctxHelper.GetIDFromContxt(c)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	if err := h.usecase.DeleteRequest(ID); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, map[string]string{"message": "Request deleted"})
}

func (h *RequestHandler) GetAllRequests(c echo.Context) error {
	requests, err := h.usecase.GetAllRequests()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, requests)
} 