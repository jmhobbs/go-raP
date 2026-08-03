package io

import (
	gio "io"
)

func ReadCompressedInteger(in gio.Reader) (int32, error) {
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
