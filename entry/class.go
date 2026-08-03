package entry

import (
	"encoding/binary"
	gio "io"

	"github.com/jmhobbs/go-raP/io"
)

type Class struct {
	Name               string
	InheritedClassName string
	Entries            []Entry
}

func (c Class) Type() EntryType {
	return EntryTypeClass
}

func ReadClass(in gio.ReadSeeker) (*Class, error) {
	className, err := io.ReadAsciiz(in)
	if err != nil {
		return nil, err
	}
	var offsetToBody uint32
	err = binary.Read(in, binary.LittleEndian, &offsetToBody)
	if err != nil {
		return nil, err
	}

	pos, err := in.Seek(0, gio.SeekCurrent)
	if err != nil {
		return nil, err
	}

	_, err = in.Seek(int64(offsetToBody), gio.SeekStart)
	if err != nil {
		return nil, err
	}

	cb, err := ReadClassBody(in)
	if err != nil {
		return nil, err
	}

	_, err = in.Seek(pos, gio.SeekStart)
	if err != nil {
		return nil, err
	}

	return &Class{
		Name:               string(className),
		InheritedClassName: cb.InheritedClassName,
		Entries:            cb.Entries,
	}, nil
}
