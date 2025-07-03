package handler

import (
	"net/http"
	"strconv"
	"uni_app/database"
	"uni_app/models"
	"uni_app/pkg/room/usecase"

	"github.com/labstack/echo/v4"
)

type RoomHandler struct {
	usecase usecase.RoomUsecase
	group   echo.Group
}

func NewRoomHandler(usecase usecase.RoomUsecase, group echo.Group) {
	handler := &RoomHandler{
		usecase: usecase,
		group:   group,
	}
	handler.initRoutes()
}

func (h *RoomHandler) initRoutes() {
	roomGroup := h.group.Group("/rooms")

	roomGroup.POST("", h.Create)
	roomGroup.PUT("/:id", h.Update)
	roomGroup.DELETE("/:id", h.Delete)
	roomGroup.GET("/:id", h.GetByID)
	roomGroup.GET("", h.List)
}

func (h *RoomHandler) Create(c echo.Context) error {
	var room models.Room
	if err := c.Bind(&room); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	if err := h.usecase.Create(&room); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusCreated, room)
}

func (h *RoomHandler) Update(c echo.Context) error {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid ID")
	}

	var room models.Room
	if err := c.Bind(&room); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	room.ID = database.PID(id)
	if err := h.usecase.Update(&room); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, room)
}

func (h *RoomHandler) Delete(c echo.Context) error {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid ID")
	}

	if err := h.usecase.Delete(database.PID(id)); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.NoContent(http.StatusNoContent)
}

func (h *RoomHandler) GetByID(c echo.Context) error {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid ID")
	}

	room, err := h.usecase.GetByID(database.PID(id))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, room)
}

func (h *RoomHandler) List(c echo.Context) error {
	var filters models.FetchRoomRequest
	if err := c.Bind(&filters); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	rooms, err := h.usecase.List(&filters)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, rooms)
}
