package entry_test

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/jmhobbs/go-raP/entry"
)

func Test_ReadClassBody(t *testing.T) {
	t.Run("valid", func(t *testing.T) {
		input := []byte{
			'A', 'p', 'p', 'l', 'e', 0x00, // inherited class name
			0x02,            // entry count
			0x04, 'a', 0x00, // entry 1: delete "a"
			0x03, 'b', 0x00, // entry 2: extern "b"
		}

		del := entry.Delete("a")
		ext := entry.Extern("b")

		cb, err := entry.ReadClassBody(bytes.NewReader(input))
		require.NoError(t, err)
		assert.Equal(t, "Apple", cb.InheritedClassName)
		assert.Equal(t, []entry.Entry{&del, &ext}, cb.Entries)
	})

	t.Run("no inherited class or entries", func(t *testing.T) {
		input := []byte{
			0x00, // inherited class name (empty)
			0x00, // entry count
		}

		cb, err := entry.ReadClassBody(bytes.NewReader(input))
		require.NoError(t, err)
		assert.Equal(t, "", cb.InheritedClassName)
		assert.Equal(t, []entry.Entry{}, cb.Entries)
	})
}
