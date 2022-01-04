package baseerr

import (
	"errors"
)

var (
	NotFoundErr   = errors.New("not found")
	DuplicatedErr = errors.New("duplicated")
)
