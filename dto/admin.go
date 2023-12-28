package dto

import "errors"

var (
	ErrEmailNotFound    = errors.New("email not found")
	ErrNotAdminID       = errors.New("your role is not admin")
	ErrAdminNotFound    = errors.New("admin not found")
	ErrPasswordNotMatch = errors.New("password do not match")
)

type (
	AdminCreateDTO struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	AdminLoginDTO struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	GiveAccess struct {
		UserID string `json:"user_id"`
		BookID string `json:"book_id"`
	}
)
