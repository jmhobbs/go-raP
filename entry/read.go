package entry

import (
	"encoding/binary"
	"errors"
	"io"
)

func Read(in io.ReadSeeker) (Entry, error) {
	var (
		entryType EntryType
		err       error
	)
	err = binary.Read(in, binary.LittleEndian, &entryType)
	if err != nil {
		return nil, err
	}
	switch entryType {
	case EntryTypeClass:
		return ReadClass(in)
	case EntryTypeAssignment:
		return ReadAssignment(in)
	case EntryTypeArray:
		return ReadArray(in)
	case EntryTypeExtern:
		return nil, errors.New("extern entry type not implemented")
	case EntryTypeDelete:
		return nil, errors.New("delete entry type not implemented")
	case EntryTypeArrayWithFlag:
		return nil, errors.New("array with flag entry type not implemented")
	}

	return nil, nil
}
