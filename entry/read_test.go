package entry_test

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/jmhobbs/go-raP/entry"
)

func Test_Read(t *testing.T) {
	t.Run("class", func(t *testing.T) {
		input := []byte{
			0x00, // entry type (class)
			'W', 'I', 'L', 'D', 'L', 'A', 'N', 'D', 'Z', '_', 'N', 'B', 'C', '_', 'P', 'o', 'u', 'c', 'h', 0x00,
			0x19, 0x00, 0x00, 0x00, // offset to body (25 bytes)
			'F', 'i', 'r', 's', 't', 'A', 'i', 'd', 'K', 'i', 't', 0x00, // inherited classname
			0x00, // entry count
		}

		e, err := entry.Read(bytes.NewReader(input))
		require.NoError(t, err)
		assert.Equal(
			t,
			&entry.Class{
				Name:               "WILDLANDZ_NBC_Pouch",
				InheritedClassName: "FirstAidKit",
				Entries:            []entry.Entry{},
			},
			e,
		)
	})

	t.Run("assignment", func(t *testing.T) {
		input := []byte{
			0x01,                                    // entry type (assignment)
			0x02,                                    // subtype (long)
			'e', 'x', 'a', 'm', 'p', 'l', 'e', 0x00, // name
			0x39, 0x30, 0x00, 0x00, // value
		}

		e, err := entry.Read(bytes.NewReader(input))
		require.NoError(t, err)
		assert.Equal(
			t,
			&entry.Assignment{
				Name:    "example",
				Subtype: entry.AssignmentTypeLong,
				Value:   int32(12345),
			},
			e,
		)
	})

	t.Run("array", func(t *testing.T) {
		input := []byte{
			0x02,                     // entry type (array)
			'n', 'a', 'm', 'e', 0x00, // name
			0x00, // count
		}

		e, err := entry.Read(bytes.NewReader(input))
		require.NoError(t, err)
		assert.Equal(
			t,
			&entry.Array{
				Name:   "name",
				Values: []entry.ArrayValue{},
			},
			e,
		)
	})

	t.Run("extern", func(t *testing.T) {
		input := []byte{
			0x03,                                    // entry type (extern)
			'o', 'u', 't', 's', 'i', 'd', 'e', 0x00, // name
		}

		e, err := entry.Read(bytes.NewReader(input))
		require.NoError(t, err)
		extern := entry.Extern("outside")
		assert.Equal(t, &extern, e)
	})

	t.Run("delete", func(t *testing.T) {
		input := []byte{
			0x04,                                    // entry type (delete)
			'g', 'o', 'o', 'd', 'b', 'y', 'e', 0x00, // name
		}

		e, err := entry.Read(bytes.NewReader(input))
		require.NoError(t, err)
		del := entry.Delete("goodbye")
		assert.Equal(t, &del, e)
	})

	t.Run("array with flag", func(t *testing.T) {
		input := []byte{
			0x05,                   // entry type (array with flag)
			0x01, 0x00, 0x00, 0x00, // flag
			'n', 'a', 'm', 'e', 0x00, // name
			0x00, // count
		}

		e, err := entry.Read(bytes.NewReader(input))
		require.NoError(t, err)
		assert.Equal(
			t,
			&entry.ArrayWithFlag{
				Name:   "name",
				Flag:   1,
				Values: []entry.ArrayValue{},
			},
			e,
		)
	})
}
