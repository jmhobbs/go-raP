package entry_test

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/jmhobbs/go-raP/entry"
)

func Test_ReadAssignment(t *testing.T) {
	t.Run("string", func(t *testing.T) {
		input := []byte{
			0x00,                                    // subtype 0
			'e', 'x', 'a', 'm', 'p', 'l', 'e', 0x00, // name
			'H', 'e', 'l', 'l', 'o', ',', ' ', 'w', 'o', 'r', 'l', 'd', '!', 0x00, // value
		}
		assignment, err := entry.ReadAssignment(bytes.NewReader(input))
		require.NoError(t, err)
		assert.Equal(
			t,
			&entry.Assignment{
				Name:    "example",
				Subtype: entry.AssignmentTypeString,
				Value:   "Hello, world!",
			},
			assignment,
		)
	})

	t.Run("variable", func(t *testing.T) {
		input := []byte{
			0x04,                                    // subtype 4
			'e', 'x', 'a', 'm', 'p', 'l', 'e', 0x00, // name
			'a', 'n', 'o', 't', 'h', 'e', 'r', 'V', 'a', 'r', 'i', 'a', 'b', 'l', 'e', 0x00, // value
		}
		assignment, err := entry.ReadAssignment(bytes.NewReader(input))
		require.NoError(t, err)
		assert.Equal(
			t,
			&entry.Assignment{
				Name:    "example",
				Subtype: entry.AssignmentTypeVariable,
				Value:   "anotherVariable",
			},
			assignment,
		)
	})

	t.Run("float", func(t *testing.T) {
		input := []byte{
			0x01,                                    // subtype 1
			'e', 'x', 'a', 'm', 'p', 'l', 'e', 0x00, // name
			0xa4, 0x70, 0x9d, 0x3f, // 1.23 float32 little-endian
		}
		assignment, err := entry.ReadAssignment(bytes.NewReader(input))
		require.NoError(t, err)
		assert.Equal(
			t,
			&entry.Assignment{
				Name:    "example",
				Subtype: entry.AssignmentTypeFloat,
				Value:   float32(1.23),
			},
			assignment,
		)
	})

	t.Run("long", func(t *testing.T) {
		input := []byte{
			0x02,                                    // subtype 2
			'e', 'x', 'a', 'm', 'p', 'l', 'e', 0x00, // name
			0x39, 0x30, 0x00, 0x00, // 12345 little-endian
		}
		assignment, err := entry.ReadAssignment(bytes.NewReader(input))
		require.NoError(t, err)
		assert.Equal(
			t,
			&entry.Assignment{
				Name:    "example",
				Subtype: entry.AssignmentTypeLong,
				Value:   int32(12345),
			},
			assignment,
		)
	})

	t.Run("unknown subtype", func(t *testing.T) {
		input := []byte{
			0x09,                                    // subtype 9 (unknown)
			'e', 'x', 'a', 'm', 'p', 'l', 'e', 0x00, // name
		}
		assignment, err := entry.ReadAssignment(bytes.NewReader(input))
		require.Error(t, err)
		assert.Nil(t, assignment)
	})
}
