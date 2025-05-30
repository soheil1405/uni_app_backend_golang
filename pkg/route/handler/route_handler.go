package handler

import (
	"net/http"
	"uni_app/database"
	"uni_app/models"
	usecase "uni_app/pkg/route/usecase"
	"uni_app/utils/ctxHelper"
	"uni_app/utils/helpers"

	"github.com/labstack/echo/v4"
)

type RouteHandler struct {
	usecase usecase.RouteUsecase
}

func NewRouteHandler(usecase usecase.RouteUsecase, e echo.Group) {
	routeHandler := &RouteHandler{usecase}

	routesRouteGroup := e.Group("/routes")
	routesRouteGroup.POST("", routeHandler.CreateRoute)
	routesRouteGroup.GET("/:id", routeHandler.GetRouteByID)
	routesRouteGroup.PUT("/:id", routeHandler.UpdateRoute)
	routesRouteGroup.DELETE("/:id", routeHandler.DeleteRoute)
	routesRouteGroup.GET("", routeHandler.GetAllRoutes)
}

func (h *RouteHandler) CreateRoute(c echo.Context) error {
	var route models.Route
	if err := c.Bind(&route); err != nil {
		return helpers.Reply(c, http.StatusBadRequest, err, nil, nil)
	}
	if err := h.usecase.CreateRoute(&route); err != nil {
		return helpers.Reply(c, http.StatusInternalServerError, err, nil, nil)
	}
	return helpers.Reply(c, http.StatusCreated, nil, map[string]interface{}{"route": route}, nil)
}

func (h *RouteHandler) GetRouteByID(c echo.Context) error {
	var (
		ID  database.PID
		err error
	)
	if ID, err = ctxHelper.GetIDFromContxt(c); err != nil {
		return helpers.Reply(c, http.StatusBadRequest, err, nil, nil)
	}
	route, err := h.usecase.GetRouteByID(c, ID, false)
	if err != nil {
		return helpers.Reply(c, http.StatusNotFound, err, nil, nil)
	}
	return helpers.Reply(c, http.StatusOK, nil, map[string]interface{}{"route": route}, nil)
}

func (h *RouteHandler) UpdateRoute(c echo.Context) (err error) {
	var route models.Route
	if route.ID, err = ctxHelper.GetIDFromContxt(c); err != nil {
		return helpers.Reply(c, http.StatusBadRequest, err, nil, nil)
	}
	if err := h.usecase.UpdateRoute(&route); err != nil {
		return helpers.Reply(c, http.StatusInternalServerError, err, nil, nil)
	}
	return helpers.Reply(c, http.StatusOK, nil, map[string]interface{}{"route": route}, nil)
}

func (h *RouteHandler) DeleteRoute(c echo.Context) error {
	ID, err := ctxHelper.GetIDFromContxt(c)
	if err != nil {
		return helpers.Reply(c, http.StatusBadRequest, err, nil, nil)
	}

	if err := h.usecase.DeleteRoute(ID); err != nil {
		return helpers.Reply(c, http.StatusInternalServerError, err, nil, nil)
	}
	return helpers.Reply(c, http.StatusOK, nil, map[string]interface{}{"message": "Route deleted"}, nil)
}

func (h *RouteHandler) GetAllRoutes(c echo.Context) error {
	var request models.FetchRouteRequest
	if err := c.Bind(&request); err != nil {
		return helpers.Reply(c, http.StatusBadRequest, err, nil, nil)
	}
	routes, paginate, err := h.usecase.GetAllRoutes(c, request)
	if err != nil {
		return helpers.Reply(c, http.StatusInternalServerError, err, nil, nil)
	}
	return helpers.Reply(c, http.StatusOK, nil, map[string]interface{}{"routes": routes}, paginate)
}
