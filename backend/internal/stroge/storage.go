package stroge

import (
	"errors"
	"internal/internal/stroge/sqlite"
)

type Storage struct {
	Storage *sqlite.Storage
}

var (
	ErrUserExists   = errors.New("user already exists")
	ErrUserNotFound = errors.New("user not found")
	ErrAppNotFound  = errors.New("app not found")
)
