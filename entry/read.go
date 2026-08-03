package entry

import (
	"encoding/binary"
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
		return ReadExtern(in)
	case EntryTypeDelete:
		return ReadDelete(in)
	case EntryTypeArrayWithFlag:
		return ReadArrayWithFlag(in)
	}

	return nil, nil
}
