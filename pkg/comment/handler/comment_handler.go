package handler

import (
	"net/http"
	"strconv"

	"uni_app/database"
	"uni_app/models"
	"uni_app/pkg/comment/usecase"
	"uni_app/utils/helpers"

	"github.com/labstack/echo/v4"
)

type CommentHandler struct {
	usecase usecase.CommentUseCase
}

func NewCommentHandler(usecase usecase.CommentUseCase) *CommentHandler {
	return &CommentHandler{
		usecase: usecase,
	}
}

func (h *CommentHandler) RegisterRoutes(e *echo.Echo) {
	comments := e.Group("/comments")
	comments.POST("", h.CreateComment)
	comments.GET("/:id", h.GetCommentByID)
	comments.PUT("/:id", h.UpdateComment)
	comments.DELETE("/:id", h.DeleteComment)
	comments.GET("/commentable/:type/:id", h.GetCommentsByCommentable)
	comments.GET("/user/:id", h.GetCommentsByUser)
	comments.GET("/replies/:id", h.GetCommentReplies)
	comments.GET("", h.GetAllComments)
}

func (h *CommentHandler) CreateComment(c echo.Context) error {
	var comment models.Comment
	if err := c.Bind(&comment); err != nil {
		return helpers.Reply(c, http.StatusBadRequest, err, nil, nil)
	}

	if err := h.usecase.CreateComment(c.Request().Context(), &comment); err != nil {
		return helpers.Reply(c, http.StatusInternalServerError, err, nil, nil)
	}

	return helpers.Reply(c, http.StatusCreated, nil, map[string]interface{}{"comment": comment}, nil)
}

func (h *CommentHandler) GetCommentByID(c echo.Context) error {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		return helpers.Reply(c, http.StatusBadRequest, err, nil, nil)
	}

	comment, err := h.usecase.GetCommentByID(c.Request().Context(), uint(id))
	if err != nil {
		return helpers.Reply(c, http.StatusNotFound, err, nil, nil)
	}

	return helpers.Reply(c, http.StatusOK, nil, map[string]interface{}{"comment": comment}, nil)
}

func (h *CommentHandler) UpdateComment(c echo.Context) error {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		return helpers.Reply(c, http.StatusBadRequest, err, nil, nil)
	}

	var comment models.Comment
	if err := c.Bind(&comment); err != nil {
		return helpers.Reply(c, http.StatusBadRequest, err, nil, nil)
	}

	comment.ID = database.PID(id)
	if err := h.usecase.UpdateComment(c.Request().Context(), &comment); err != nil {
		return helpers.Reply(c, http.StatusInternalServerError, err, nil, nil)
	}

	return helpers.Reply(c, http.StatusOK, nil, map[string]interface{}{"comment": comment}, nil)
}

func (h *CommentHandler) DeleteComment(c echo.Context) error {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		return helpers.Reply(c, http.StatusBadRequest, err, nil, nil)
	}

	if err := h.usecase.DeleteComment(c.Request().Context(), uint(id)); err != nil {
		return helpers.Reply(c, http.StatusInternalServerError, err, nil, nil)
	}

	return helpers.Reply(c, http.StatusOK, nil, map[string]interface{}{"message": "Comment deleted successfully"}, nil)
}

func (h *CommentHandler) GetCommentsByCommentable(c echo.Context) error {
	commentableType := c.Param("type")
	commentableID, err := database.ParsePID(c.Param("id"))
	if err != nil {
		return helpers.Reply(c, http.StatusBadRequest, err, nil, nil)
	}

	comments, err := h.usecase.GetCommentsByCommentable(c.Request().Context(), commentableID, commentableType)
	if err != nil {
		return helpers.Reply(c, http.StatusInternalServerError, err, nil, nil)
	}

	return helpers.Reply(c, http.StatusOK, nil, map[string]interface{}{"comments": comments}, nil)
}

func (h *CommentHandler) GetCommentsByUser(c echo.Context) error {
	userID, err := database.ParsePID(c.Param("id"))
	if err != nil {
		return helpers.Reply(c, http.StatusBadRequest, err, nil, nil)
	}

	comments, err := h.usecase.GetCommentsByUser(c.Request().Context(), userID)
	if err != nil {
		return helpers.Reply(c, http.StatusInternalServerError, err, nil, nil)
	}

	return helpers.Reply(c, http.StatusOK, nil, map[string]interface{}{"comments": comments}, nil)
}

func (h *CommentHandler) GetCommentReplies(c echo.Context) error {
	parentID, err := database.ParsePID(c.Param("id"))
	if err != nil {
		return helpers.Reply(c, http.StatusBadRequest, err, nil, nil)
	}

	replies, err := h.usecase.GetCommentReplies(c.Request().Context(), parentID)
	if err != nil {
		return helpers.Reply(c, http.StatusInternalServerError, err, nil, nil)
	}

	return helpers.Reply(c, http.StatusOK, nil, map[string]interface{}{"replies": replies}, nil)
}

func (h *CommentHandler) GetAllComments(c echo.Context) error {
	var request models.FetchCommentRequest
	if err := c.Bind(&request); err != nil {
		return helpers.Reply(c, http.StatusBadRequest, err, nil, nil)
	}

	comments, err := h.usecase.GetAllComments(c.Request().Context(), request)
	if err != nil {
		return helpers.Reply(c, http.StatusInternalServerError, err, nil, nil)
	}

	return helpers.Reply(c, http.StatusOK, nil, map[string]interface{}{"comments": comments}, nil)
}
