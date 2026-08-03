package raP

import (
	"errors"
)

var (
	ErrInvalidFileHeader        = errors.New("invalid file header")
	ErrAuthenticatedUnsupported = errors.New("authenticated raP not supported")
)
