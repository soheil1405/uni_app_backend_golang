package helpers

import (
	"errors"
	"net/http"
	"uni_app/utils/templates"
)

var (
	ErrNotFound                 = errors.New("not found")
	ErrAlreadyExist             = errors.New("already exist")
	ErrInvalidRequest           = errors.New("invalid request")
	ErrSystemItemDelete         = errors.New("you cant delete system items")
	ErrInvalidRestClientInfo    = errors.New("invalid rest client info")
	ErrCacheInfo                = errors.New("invalid cache info")
	ErrLocationNotCovered       = errors.New("selected location isn't in branch covered area")
	ErrWrongPassword            = errors.New("wrong password, try another")
	ErrInvalidAdapter           = errors.New("invalid adapter")
	ErrorUnAuthorized           = generateErr("unauthorized")
	ErrUnAuthorizedInValidToken = generateErr("invalid token")
	ErrUnAuthorizedTokenExpired = generateErr("token has been expired")
	ErrorBadRequest             = generateErr("bad request")
	ErrorWrongPassword          = generateErr("wrong passwrod")
	ErrorAccessDenied           = generateErr("access denied")
	ErrorInvalidUserPass        = generateErr("invalid username or password")
	ErrorInvalidUserRoles       = generateErr("user dosent have role")
	ErrUserIsNotActive          = generateErr("user is not active")
	ErrInvalidUserID            = generateErr("invalid user id")
	ErrInvalidStudentID         = generateErr("invalid student id")
)

func generateErr(errMsg string) error {
	return errors.New(errMsg)
}

// MyError ...
type MyError struct {
	Err  error
	Code int
}

// NewMyError ...
func NewMyError() MyError {
	return MyError{
		Err:  nil,
		Code: http.StatusOK,
	}
}

// Default ...
func (m *MyError) Default() {
	m.Err = nil
	m.Code = http.StatusOK
}

// SetError ...
func (m *MyError) SetError(err error, code int) {
	m.Err = err
	m.Code = code
}

// Models for RESTFUL Data transmitions

// MessageTemplate ...
type MessageTemplate struct {
	Message string `json:"message"`
}

// MessageResponse ... -> Customer Response Template witch Receives From ecommerce server
type MessageResponse struct {
	*templates.ResponseTemplate
	Data MessageTemplate `json:"data"`
}
