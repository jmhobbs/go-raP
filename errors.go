package raP

import (
	"errors"
	"fmt"
)

var (
	ErrInvalidFileHeader = errors.New("invalid file header")
	ErrUnsupportedType   = errors.New("unsupported type")
)

type UnknownPacketTypeError struct {
	Type uint8
}

func (u UnknownPacketTypeError) Error() string {
	return fmt.Sprintf("unknown packet type: %x", u.Type)
}
