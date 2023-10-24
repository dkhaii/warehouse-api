package helpers

import "errors"

var (
	ErrRequestTimedOut = errors.New("request timed out, operation took too long")
	ErrRoleNotFound = errors.New("role not found")
	ErrUserNotFound = errors.New("user not found")
	ErrItemNotFound = errors.New("item not found")
	ErrCategoryNotFound = errors.New("category not found")
	ErrLocationNotFound = errors.New("location not found")
	ErrOrderNotFound = errors.New("order not found")
)