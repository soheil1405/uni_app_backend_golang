package models

import "errors"

var (
	ErrorBadRequest             = generateErr("bad request")
	ErrorAccessDenied           = generateErr("access denied")
	ErrorInvalidUserPass        = generateErr("invalid username or password")
	ErrorInvalidUserRoles       = generateErr("user dosent have role")
	ErrUserIsNotActive          = generateErr("user is not active")
	ErrUnAuthorizedInValidToken = generateErr("invalid token")
	ErrUnAuthorizedTokenExpired = generateErr("token has been expired")
	ErrInvalidUserID            = generateErr("invalid user id")
	ErrInvalidStudentID         = generateErr("invalid student id")
)

func generateErr(errMsg string) error {
	return errors.New(errMsg)
}
