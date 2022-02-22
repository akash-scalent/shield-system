package constants

import (
	"errors"
)

var (
	ErrAvengerNameOrHeroNameEmpty = errors.New("avenger name or heroname cannot be empty")
	ErrDuplicateHeroName          = errors.New("choose different heroname")
	ErrDuplicateNameOfUser        = errors.New("choose different name")
	ErrDuplicateEntity            = errors.New("duplicate entity")
	ErrEntityNotFound             = errors.New("entity not found")
)
