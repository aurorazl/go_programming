package model

import "errors"

var (
	ERROR_USER_NOTEXISTS = errors.New("user not exists!")
	ERROR_USER_TEXISTS   = errors.New("user exists!")
	ERROR_PWD            = errors.New("pwd is wrong!")
)
