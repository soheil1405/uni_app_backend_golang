package handlers

import (
	"net/http"
	"uni_app/models"
	usecases "uni_app/pkg/rating/usecase"
	"uni_app/utils/helpers"

	"github.com/labstack/echo/v4"
)

type RatingHandler struct {
	usecase usecases.RatingUsecase
}

func NewRatingHandler(usecase usecases.RatingUsecase, e echo.Group) {
	handler := &RatingHandler{
		usecase: usecase,
	}

	e.POST("/ratings", handler.AddRating)
	e.GET("/ratings/:id", handler.GetRatingByID)
	e.PUT("/ratings/:id", handler.UpdateRating)
	e.DELETE("/ratings/:id", handler.DeleteRating)
	e.GET("/ratings", handler.GetAllRatings)
	e.GET("/ratable/:type/:id/ratings", handler.GetRatableRatings)
	e.GET("/ratable/:type/:id/average-rating", handler.GetRatableAverageRating)
	e.GET("/users/:user_id/ratable/:type/:id/rating", handler.GetUserRating)
}

func (h *RatingHandler) AddRating(c echo.Context) error {
	var rating models.Rating
	if err := c.Bind(&rating); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	if err := h.usecase.AddRating(&rating); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusCreated, rating)
}

func (h *RatingHandler) GetRatingByID(c echo.Context) error {
	ID, err := helpers.GetIDFromContxt(c)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	rating, err := h.usecase.GetRatingByID(c, ID, false)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, rating)
}

func (h *RatingHandler) UpdateRating(c echo.Context) error {
	ID, err := helpers.GetIDFromContxt(c)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	var rating models.Rating
	if err := c.Bind(&rating); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	rating.ID = ID
	if err := h.usecase.UpdateRating(&rating); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, rating)
}

func (h *RatingHandler) DeleteRating(c echo.Context) error {
	ID, err := helpers.GetIDFromContxt(c)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	if err := h.usecase.DeleteRating(ID); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.NoContent(http.StatusNoContent)
}

func (h *RatingHandler) GetAllRatings(c echo.Context) error {
	var request models.FetchRatingRequest
	if err := c.Bind(&request); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	ratings, err := h.usecase.GetAllRatings(c, request)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, ratings)
}

func (h *RatingHandler) GetRatableRatings(c echo.Context) error {
	ratableType := c.Param("type")
	ratableID, err := helpers.GetIDFromContxt(c)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	ratings, err := h.usecase.GetRatableRatings(ratableID, ratableType)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, ratings)
}

func (h *RatingHandler) GetRatableAverageRating(c echo.Context) error {
	ratableType := c.Param("type")
	ratableID, err := helpers.GetIDFromContxt(c)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	avgRating, err := h.usecase.GetRatableAverageRating(ratableID, ratableType)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, map[string]float64{"average_rating": avgRating})
}

func (h *RatingHandler) GetUserRating(c echo.Context) error {
	userID, err := helpers.GetIDFromContxt(c)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	ratableType := c.Param("type")
	ratableID, err := helpers.GetIDFromContxt(c)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	rating, err := h.usecase.GetUserRating(userID, ratableID, ratableType)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, rating)
}
