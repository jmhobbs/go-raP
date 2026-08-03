package io_test

import (
	"bytes"
	"testing"

	"github.com/jmhobbs/go-raP/io"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestReadAsciiz(t *testing.T) {
	t.Run("valid", func(t *testing.T) {
		input := []byte{'h', 'e', 'l', 'l', 'o', 0x00}
		str, err := io.ReadAsciiz(bytes.NewReader(input))
		require.NoError(t, err)
		assert.Equal(t, []byte("hello"), str)
	})

	t.Run("empty", func(t *testing.T) {
		input := []byte{0x00}
		str, err := io.ReadAsciiz(bytes.NewReader(input))
		require.NoError(t, err)
		assert.Equal(t, []byte{}, str)
	})

	t.Run("missing null terminator", func(t *testing.T) {
		input := []byte{'h', 'e', 'l', 'l', 'o'}
		_, err := io.ReadAsciiz(bytes.NewReader(input))
		assert.Error(t, err)
	})
}
