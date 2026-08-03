package entry

import (
	gio "io"

	"github.com/jmhobbs/go-raP/io"
)

type Extern string

func (e Extern) Type() EntryType {
	return EntryTypeExtern
}

func ReadExtern(in gio.Reader) (*Extern, error) {
	classname, err := io.ReadAsciiz(in)
	if err != nil {
		return nil, err
	}
	extern := Extern(classname)
	return &extern, nil
}
