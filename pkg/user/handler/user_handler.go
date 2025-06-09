package handlers

import (
	"net/http"
	"uni_app/database"
	"uni_app/models"
	usecases "uni_app/pkg/user/usecase"
	"uni_app/utils/ctxHelper"
	"uni_app/utils/helpers"

	"github.com/labstack/echo/v4"
)

type UserHandler struct {
	usecase usecases.UserUsecase
}

func NewUserHandler(usecase usecases.UserUsecase, e echo.Group) {
	userHandler := &UserHandler{usecase}

	// Public routes
	// Protected routes
	users := e.Group("/users")
	users.POST("", userHandler.CreateUser)
	users.POST("/register", userHandler.RegisterUser)
	users.POST("/login", userHandler.LoginUser)
	users.GET("/:id", userHandler.GetUserByID)
	users.PUT("/:id", userHandler.UpdateUser)
	users.DELETE("/:id", userHandler.DeleteUser)
	users.GET("", userHandler.GetAllUsers)
}

func (h *UserHandler) CreateUser(c echo.Context) error {
	var user models.User
	if err := c.Bind(&user); err != nil {
		return helpers.Reply(c, http.StatusBadRequest, err, nil, nil)
	}
	if err := h.usecase.CreateUser(&user); err != nil {
		return helpers.Reply(c, http.StatusInternalServerError, err, nil, nil)
	}
	return helpers.Reply(c, http.StatusCreated, nil, map[string]interface{}{"user": user}, nil)
}

func (h *UserHandler) GetUserByID(c echo.Context) error {
	var (
		ID  database.PID
		err error
	)
	if ID, err = ctxHelper.GetIDFromContxt(c); err != nil {
		return helpers.Reply(c, http.StatusBadRequest, err, nil, nil)
	}
	user, err := h.usecase.GetUserByID(c, ID, false)
	if err != nil {
		return helpers.Reply(c, http.StatusNotFound, err, nil, nil)
	}
	return helpers.Reply(c, http.StatusOK, nil, map[string]interface{}{"user": user}, nil)
}

func (h *UserHandler) UpdateUser(c echo.Context) error {
	var (
		user models.User
		err  error
	)
	if err := c.Bind(&user); err != nil {
		return helpers.Reply(c, http.StatusBadRequest, err, nil, nil)
	}
	if user.ID, err = ctxHelper.GetIDFromContxt(c); err != nil {
		return helpers.Reply(c, http.StatusBadRequest, err, nil, nil)
	}
	if err := h.usecase.UpdateUser(&user); err != nil {
		return helpers.Reply(c, http.StatusInternalServerError, err, nil, nil)
	}
	return helpers.Reply(c, http.StatusOK, nil, map[string]interface{}{"user": user}, nil)
}

func (h *UserHandler) DeleteUser(c echo.Context) error {
	var (
		ID  database.PID
		err error
	)
	if ID, err = ctxHelper.GetIDFromContxt(c); err != nil {
		return helpers.Reply(c, http.StatusBadRequest, err, nil, nil)
	}
	if err := h.usecase.DeleteUser(ID); err != nil {
		return helpers.Reply(c, http.StatusInternalServerError, err, nil, nil)
	}
	return helpers.Reply(c, http.StatusOK, nil, map[string]interface{}{"message": "User deleted"}, nil)
}

func (h *UserHandler) GetAllUsers(c echo.Context) error {
	var request models.FetchUserRequest
	if err := c.Bind(&request); err != nil {
		return helpers.Reply(c, http.StatusBadRequest, err, nil, nil)
	}
	users, paginate, err := h.usecase.GetAllUsers(c, request)
	if err != nil {
		return helpers.Reply(c, http.StatusInternalServerError, err, nil, nil)
	}
	return helpers.Reply(c, http.StatusOK, nil, map[string]interface{}{"users": users}, paginate)
}

func (h *UserHandler) RegisterUser(c echo.Context) error {
	var request models.UserRegisterRequest
	if err := c.Bind(&request); err != nil {
		return helpers.Reply(c, http.StatusBadRequest, err, nil, nil)
	}

	if err := request.IsValid(); err != nil {
		return helpers.Reply(c, http.StatusBadRequest, err, nil, nil)
	}

	if err := h.usecase.RegisterUser(c, &request); err != nil {
		return helpers.Reply(c, http.StatusBadRequest, err, nil, nil)
	}
	return helpers.Reply(c, http.StatusCreated, nil, map[string]interface{}{"user": request}, nil)
}

func (h *UserHandler) LoginUser(c echo.Context) error {
	var loginRequest struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	if err := c.Bind(&loginRequest); err != nil {
		return helpers.Reply(c, http.StatusBadRequest, err, nil, nil)
	}
	user, err := h.usecase.LoginUser(loginRequest.Username, loginRequest.Password)
	if err != nil {
		return helpers.Reply(c, http.StatusUnauthorized, err, nil, nil)
	}
	return helpers.Reply(c, http.StatusOK, nil, map[string]interface{}{"user": user}, nil)
}
