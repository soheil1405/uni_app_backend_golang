package models

import "errors"

var (
	ErrorInvalidUserPass = generateErr("invalid username or password")
)

func generateErr(errMsg string) error {
	return errors.New(errMsg)
}
