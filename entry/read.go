package entry

import (
	"encoding/binary"
	"fmt"
	"io"
	"os"
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
		fmt.Fprintln(os.Stderr, "extern")
		return nil, nil
	case EntryTypeDelete:
		fmt.Fprintln(os.Stderr, "delete")
		return nil, nil
	case EntryTypeArrayWithFlag:
		fmt.Fprintln(os.Stderr, "array with flag")
		return nil, nil
	}

	return nil, nil
}
