package consts

import "errors"

var (
	ErrParamsInvalid = errors.New("params error")
	ErrNotExpectError = errors.New("not expect error")

	ErrLoginError = errors.New("user name or password not correct")
)

