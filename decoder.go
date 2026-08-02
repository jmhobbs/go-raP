package raP

import (
	"encoding/binary"
	"fmt"
	"io"
	"os"
)

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

	root, err := ReadClassBody(in)
	if err != nil {
		return nil, err
	}

	return &File{Entries: root.Entries}, nil
}

type classBody struct {
	InheritedClassName string
	NumberOfEntries    int32
	Entries            []Entry
}

func ReadClassBody(in io.ReadSeeker) (*classBody, error) {
	inheritedClassName, err := readAsciiz(in)
	if err != nil {
		return nil, err
	}

	entryCount, err := readCompressedInteger(in)
	if err != nil {
		return nil, err
	}

	cb := &classBody{
		InheritedClassName: string(inheritedClassName),
		NumberOfEntries:    entryCount,
		Entries:            make([]Entry, entryCount),
	}

	for i := range entryCount {
		entry, err := ReadEntry(in)
		if err != nil {
			return nil, err
		}
		cb.Entries[i] = entry
	}

	return cb, nil
}

func ReadEntry(in io.ReadSeeker) (Entry, error) {
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
		name, err := readAsciiz(in)
		if err != nil {
			return nil, err
		}
		elementCount, err := readCompressedInteger(in)
		if err != nil {
			return nil, err
		}
		array := Array{Name: string(name), Values: make([]ArrayValue, elementCount)}

		var elementType ArrayValueType
		for i := range elementCount {
			err = binary.Read(in, binary.LittleEndian, &elementType)
			if err != nil {
				return nil, err
			}
			switch elementType {
			case ArrayValueTypeString:
				value, err := readAsciiz(in)
				if err != nil {
					return nil, err
				}
				array.Values[i] = ArrayValue{Type: ArrayValueTypeString, Value: string(value)}
			case ArrayValueTypeFloat:
				var value float32
				err = binary.Read(in, binary.LittleEndian, &value)
				if err != nil {
					return nil, err
				}
				array.Values[i] = ArrayValue{Type: ArrayValueTypeFloat, Value: value}
			case ArrayValueTypeLong:
				var value int32
				err = binary.Read(in, binary.LittleEndian, &value)
				if err != nil {
					return nil, err
				}
				array.Values[i] = ArrayValue{Type: ArrayValueTypeLong, Value: value}
			case ArrayValueTypeArray:
				return nil, fmt.Errorf("error: recursive arrays not implemented")
			case ArrayValueTypeVariable:
				value, err := readAsciiz(in)
				if err != nil {
					return nil, err
				}
				array.Values[i] = ArrayValue{Type: ArrayValueTypeVariable, Value: string(value)}
			default:
				return nil, fmt.Errorf("error: unknown array element type %d", elementType)
			}
		}
		return &array, nil
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

func ReadClass(in io.ReadSeeker) (*Class, error) {
	className, err := readAsciiz(in)
	if err != nil {
		return nil, err
	}
	var offsetToBody uint32
	err = binary.Read(in, binary.LittleEndian, &offsetToBody)
	if err != nil {
		return nil, err
	}

	// TODO: instead of seeking back and forth hydrate classes at the end?

	pos, err := in.Seek(0, io.SeekCurrent)
	if err != nil {
		return nil, err
	}

	_, err = in.Seek(int64(offsetToBody), io.SeekStart)
	if err != nil {
		return nil, err
	}

	cb, err := ReadClassBody(in)
	if err != nil {
		return nil, err
	}

	_, err = in.Seek(pos, io.SeekStart)
	if err != nil {
		return nil, err
	}

	return &Class{
		Name:               string(className),
		InheritedClassName: cb.InheritedClassName,
		Entries:            cb.Entries,
	}, nil
}

func ReadAssignment(in io.ReadSeeker) (*Assignment, error) {
	var subtype uint8
	err := binary.Read(in, binary.LittleEndian, &subtype)
	if err != nil {
		return nil, err
	}

	assignment := Assignment{}

	name, err := readAsciiz(in)
	if err != nil {
		return nil, err
	}
	assignment.Name = string(name)

	switch AssignmentType(subtype) {
	case AssignmentTypeString:
		value, err := readAsciiz(in)
		if err != nil {
			return nil, err
		}
		assignment.Value = string(value)
	case AssignmentTypeFloat:
		var value float32
		err = binary.Read(in, binary.LittleEndian, &value)
		if err != nil {
			return nil, err
		}
		assignment.Value = value
	case AssignmentTypeLong:
		var value int32
		err = binary.Read(in, binary.LittleEndian, &value)
		if err != nil {
			return nil, err
		}
		assignment.Value = value
	default:
		return nil, fmt.Errorf("error: unknown assignment type %d", subtype)
	}

	assignment.Subtype = AssignmentType(subtype)

	return &assignment, nil
}
