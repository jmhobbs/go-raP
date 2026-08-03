package entry

import (
	"encoding/binary"
	"fmt"
	gio "io"

	"github.com/jmhobbs/go-raP/io"
)

type ArrayValueType uint8

const (
	ArrayValueTypeString   ArrayValueType = 0
	ArrayValueTypeFloat    ArrayValueType = 1
	ArrayValueTypeLong     ArrayValueType = 2
	ArrayValueTypeArray    ArrayValueType = 3
	ArrayValueTypeVariable ArrayValueType = 4
)

type Array struct {
	Name   string
	Values []ArrayValue
}

func (a Array) Type() EntryType {
	return EntryTypeArray
}

type ArrayValue struct {
	Type  ArrayValueType
	Value any
}

func ReadArray(in gio.Reader) (*Array, error) {
	name, err := io.ReadAsciiz(in)
	if err != nil {
		return nil, err
	}

	values, err := readArrayValues(in)
	if err != nil {
		return nil, err
	}
	return &Array{Name: string(name), Values: values}, nil

}

func readArrayValues(in gio.Reader) ([]ArrayValue, error) {
	elementCount, err := io.ReadCompressedInteger(in)
	if err != nil {
		return nil, err
	}
	values := make([]ArrayValue, elementCount)

	var elementType ArrayValueType
	for i := range elementCount {
		err = binary.Read(in, binary.LittleEndian, &elementType)
		if err != nil {
			return nil, err
		}
		switch elementType {
		case ArrayValueTypeString:
			value, err := io.ReadAsciiz(in)
			if err != nil {
				return nil, err
			}
			values[i] = ArrayValue{Type: ArrayValueTypeString, Value: string(value)}
		case ArrayValueTypeFloat:
			var value float32
			err = binary.Read(in, binary.LittleEndian, &value)
			if err != nil {
				return nil, err
			}
			values[i] = ArrayValue{Type: ArrayValueTypeFloat, Value: value}
		case ArrayValueTypeLong:
			var value int32
			err = binary.Read(in, binary.LittleEndian, &value)
			if err != nil {
				return nil, err
			}
			values[i] = ArrayValue{Type: ArrayValueTypeLong, Value: value}
		case ArrayValueTypeArray:
			value, err := readArrayValues(in)
			if err != nil {
				return nil, err
			}
			values[i] = ArrayValue{Type: ArrayValueTypeArray, Value: value}
		case ArrayValueTypeVariable:
			value, err := io.ReadAsciiz(in)
			if err != nil {
				return nil, err
			}
			values[i] = ArrayValue{Type: ArrayValueTypeVariable, Value: string(value)}
		default:
			return nil, fmt.Errorf("error: unknown array element type %d", elementType)
		}
	}
	return values, nil
}
