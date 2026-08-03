package entry

import (
	gio "io"

	"github.com/jmhobbs/go-raP/io"
)

type classBody struct {
	InheritedClassName string
	NumberOfEntries    int32
	Entries            []Entry
}

func ReadClassBody(in gio.ReadSeeker) (*classBody, error) {
	inheritedClassName, err := io.ReadAsciiz(in)
	if err != nil {
		return nil, err
	}

	entryCount, err := io.ReadCompressedInteger(in)
	if err != nil {
		return nil, err
	}

	cb := &classBody{
		InheritedClassName: string(inheritedClassName),
		NumberOfEntries:    entryCount,
		Entries:            make([]Entry, entryCount),
	}

	for i := range entryCount {
		entry, err := Read(in)
		if err != nil {
			return nil, err
		}
		cb.Entries[i] = entry
	}

	return cb, nil
}
