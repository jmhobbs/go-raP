package raP

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
)

func Decode(in io.Reader, out io.Writer) error {
	buf, err := readN(in, 4)
	if err != nil {
		return err
	}

	if !bytes.Equal([]byte{0, 'r', 'a', 'P'}, buf) {
		return ErrInvalidFileHeader
	}

	buf, err = readN(in, 3)
	if err != nil {
		return err
	}

	if !bytes.Equal([]byte{4, 0, 0}, buf) {
		return ErrUnsupportedType
	}

	var packetType uint8
	err = binary.Read(in, binary.LittleEndian, &packetType)
	if err != nil {
		return err
	}

	switch packetType {
	case PacketTypeClassname:
	case PacketTypeTokenNames:
	case PacketTypeArrays:
	default:
		return UnknownPacketTypeError{Type: packetType}
	}

	return nil
}

func readN(in io.Reader, count int) ([]byte, error) {
	buf := make([]byte, count)
	n, err := in.Read(buf)
	if err != nil {
		return nil, err
	}
	if n != count {
		return nil, fmt.Errorf("bad read, wanted %d bytes, got %d", count, n)
	}
	return buf, nil
}
