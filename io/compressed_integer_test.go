package io_test

import (
	"bytes"
	"testing"

	"github.com/jmhobbs/go-raP/io"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_ReadCompressedInteger(t *testing.T) {
	t.Run("one byte", func(t *testing.T) {
		input := []byte{0x01}
		i, err := io.ReadCompressedInteger(bytes.NewReader(input))
		require.NoError(t, err)
		assert.Equal(t, int32(1), i)
	})

	t.Run("two bytes", func(t *testing.T) {
		input := []byte{
			0b10000001, // 1 + contiinuation bit
			0b00000001, // 1 << 8
		}
		i, err := io.ReadCompressedInteger(bytes.NewReader(input))
		require.NoError(t, err)
		// TODO: Verify this value as the docs are confusing
		assert.Equal(t, int32(257), i)
	})

	t.Run("missing continuation byte", func(t *testing.T) {
		input := []byte{
			0b10000001, // 1 + contiinuation bit
		}
		_, err := io.ReadCompressedInteger(bytes.NewReader(input))
		require.Error(t, err)
	})
}
