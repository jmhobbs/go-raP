package io

import (
	gio "io"
)

func ReadAsciiz(in gio.Reader) ([]byte, error) {
	var (
		buf        = make([]byte, 1)
		out []byte = []byte{}
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
