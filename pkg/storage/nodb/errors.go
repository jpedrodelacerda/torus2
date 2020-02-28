package nodb

import "errors"

var (
	ErrExistingUser = errors.New("user already exists")
	ErrUserNotFound = errors.New("user not found")
)
