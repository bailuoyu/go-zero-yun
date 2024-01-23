package model

import (
	"errors"

	mon "go-zero-yun/pkg/monkit"
)

var (
	ErrNotFound        = mon.ErrNotFound
	ErrInvalidObjectId = errors.New("invalid objectId")
)
