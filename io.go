package raP

import (
	"encoding/binary"
	"errors"
	"io"
)

func readAsciiz(in io.Reader) ([]byte, error) {
	var (
		buf = make([]byte, 1)
		out []byte
		err error
	)

	for {
		_, err = in.Read(buf)
		if err != nil {
			return nil, err
		}
		if buf[0] == 0 {
			break
		}
		out = append(out, buf[0])
	}

	return out, nil
}

func readCompressedInteger(in io.Reader) (int32, error) {
	var (
		buf       = make([]byte, 1)
		out int32 = 0
		err error
	)

	for {
		_, err = in.Read(buf)
		if err != nil {
			return out, err
		}

		if buf[0]&0x80 == 0x80 {
			// continue reading
			out |= int32(buf[0] & 0x7F)
			out <<= 8
		} else {
			// done
			out |= int32(buf[0])
			break
		}
	}

	return out, nil
}

var ErrStringIndexNotFound = errors.New("indexed string empty but not found in index")

type StringIndex map[uint16]string

func NewStringIndex() StringIndex {
	return make(StringIndex)
}

func (i StringIndex) Read(in io.Reader) (string, error) {
	var index uint16
	err := binary.Read(in, binary.LittleEndian, &index)
	if err != nil {
		return "", err
	}
	str, err := readAsciiz(in)
	if err != nil {
		return "", err
	}
	if len(str) == 0 {
		str, ok := i[index]
		if !ok {
			return "", ErrStringIndexNotFound
		}
		return string(str), nil
	}

	i[index] = string(str)
	return string(str), nil
}
