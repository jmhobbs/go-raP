package entry

import (
	gio "io"

	"github.com/jmhobbs/go-raP/io"
)

type Delete string

func (e Delete) Type() EntryType {
	return EntryTypeDelete
}

func ReadDelete(in gio.Reader) (*Delete, error) {
	classname, err := io.ReadAsciiz(in)
	if err != nil {
		return nil, err
	}
	del := Delete(classname)
	return &del, nil
}
