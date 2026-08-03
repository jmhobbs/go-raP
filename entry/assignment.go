package entry

import (
	"encoding/binary"
	"fmt"
	gio "io"

	"github.com/jmhobbs/go-raP/io"
)

type AssignmentType uint8

const (
	AssignmentTypeString   AssignmentType = 0
	AssignmentTypeFloat    AssignmentType = 1
	AssignmentTypeLong     AssignmentType = 2
	AssignmentTypeVariable AssignmentType = 4
)

type Assignment struct {
	Name    string
	Subtype AssignmentType
	Value   any
}

func (a Assignment) Type() EntryType {
	return EntryTypeAssignment
}

func ReadAssignment(in gio.Reader) (*Assignment, error) {
	var subtype uint8
	err := binary.Read(in, binary.LittleEndian, &subtype)
	if err != nil {
		return nil, err
	}

	assignment := Assignment{}

	name, err := io.ReadAsciiz(in)
	if err != nil {
		return nil, err
	}
	assignment.Name = string(name)

	switch AssignmentType(subtype) {
	case AssignmentTypeString:
		fallthrough
	case AssignmentTypeVariable:
		value, err := io.ReadAsciiz(in)
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
