package raP

import (
	"encoding/binary"
	"io"

	"github.com/jmhobbs/go-raP/entry"
)

type File struct {
	Entries []entry.Entry
}

type rapHeader struct {
	Signature     [4]byte
	Always0       uint32
	Always8       uint32
	OffsetToEnums uint32
}

func Decode(in io.ReadSeeker) (*File, error) {
	var header rapHeader
	err := binary.Read(in, binary.LittleEndian, &header)
	if err != nil {
		return nil, err
	}

	if header.Signature != [4]byte{0, 'r', 'a', 'P'} {
		return nil, ErrInvalidFileHeader
	}

	if header.Always0 != 0 || header.Always8 != 8 {
		return nil, ErrAuthenticatedUnsupported
	}

	root, err := entry.ReadClassBody(in)
	if err != nil {
		return nil, err
	}

	return &File{Entries: root.Entries}, nil
}
