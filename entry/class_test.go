package entry_test

import (
	"bytes"
	gio "io"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/jmhobbs/go-raP/entry"
)

func Test_ReadClass(t *testing.T) {
	t.Run("valid", func(t *testing.T) {
		input := []byte{
			'W', 'I', 'L', 'D', 'L', 'A', 'N', 'D', 'Z', '_', 'N', 'B', 'C', '_', 'P', 'o', 'u', 'c', 'h', 0x00,
			0x18, 0x00, 0x00, 0x00, // offset to body (24 bytes)
			'F', 'i', 'r', 's', 't', 'A', 'i', 'd', 'K', 'i', 't', 0x00, // inherited classname
			0x00, // entry count
		}

		c, err := entry.ReadClass(bytes.NewReader(input))
		require.NoError(t, err)
		assert.Equal(
			t,
			&entry.Class{
				Name:               "WILDLANDZ_NBC_Pouch",
				InheritedClassName: "FirstAidKit",
				Entries:            []entry.Entry{},
			},
			c,
		)
	})

	t.Run("body can come before the header", func(t *testing.T) {
		input := []byte{
			// body
			'F', 'i', 'r', 's', 't', 'A', 'i', 'd', 'K', 'i', 't', 0x00, // inherited classname
			0x00, // entry count

			// class entry
			'W', 'I', 'L', 'D', 'L', 'A', 'N', 'D', 'Z', '_', 'N', 'B', 'C', '_', 'P', 'o', 'u', 'c', 'h', 0x00,
			0x00, 0x00, 0x00, 0x00, // offset to body (0 bytes)

			// marker so we know reader position was restored
			0xFF,
		}

		r := bytes.NewReader(input)
		_, err := r.Seek(13, gio.SeekStart) // go to start of class entry
		require.NoError(t, err)

		c, err := entry.ReadClass(r)
		require.NoError(t, err)
		assert.Equal(
			t,
			&entry.Class{
				Name:               "WILDLANDZ_NBC_Pouch",
				InheritedClassName: "FirstAidKit",
				Entries:            []entry.Entry{},
			},
			c,
		)

		marker, err := r.ReadByte()
		require.NoError(t, err)
		assert.Equal(t, byte(0xFF), marker)
	})
}
