package domain

import (
	"errors"
)

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
	ErrUserNotFound = errors.New("존재하지 않는 회원 정보 입니다.")
	// ErrUserConflict will throw if the user already exists
	ErrUserConflict = errors.New("해당 휴대폰 번호로 가입된 계정이 이미 존재합니다.")
	// ErrWrongPassword will throw if the account is not authenticated
	ErrWrongPassword = errors.New("비밀번호가 올바르지 않습니다.")
	// ErrInvalidAccessToken will throw if the access token is not authorized
	ErrInvalidAccessToken = errors.New("토큰이 올바르지 않습니다.")
	// ErrRequiredAccessToken will throw if the access token is required
	ErrRequiredAccessToken = errors.New("토큰이 필요합니다.")
	// ErrProductNotFound will throw if the product is not exists
	ErrProductNotFound = errors.New("존재하지 않는 상품 입니다.")
	// ErrInvalidUser will throw if the user can't access the product
	ErrInvalidUser = errors.New("해당 상품에 접근할 수 없습니다.")
)
