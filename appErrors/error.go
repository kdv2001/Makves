package appErrors

import "github.com/pkg/errors"

var (
	ErrItemNotFound = errors.New("item not found")
)
