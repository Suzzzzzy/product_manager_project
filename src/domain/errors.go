package domain

import "errors"

var (
	// ErrInternalServerError will throw if any the Internal Server Error happen
	ErrInternalServerError = errors.New("internal Server Error")
	// ErrNotFound will throw if the requested item is not exists
	ErrNotFound = errors.New("your requested Item is not found")
	// ErrConflict will throw if the current action already exists
	ErrConflict = errors.New("your Item already exist")
	// ErrBadParamInput will throw if the given request-body or params is not valid
	ErrBadParamInput = errors.New("given Param is not valid")
	// ErrUserNotFound will throw if the requested item is not exists
	ErrUserNotFound = errors.New("requested User is not found")
	// ErrUserConflict will throw if the user already exists
	ErrUserConflict = errors.New("해당 휴대폰 번호로 가입된 계정이 이미 존재합니다.")
)
