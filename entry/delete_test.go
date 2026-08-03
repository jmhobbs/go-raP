package entry_test

import (
	"bytes"
	"testing"

	"github.com/jmhobbs/go-raP/entry"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_ReadDelete(t *testing.T) {
	input := []byte{
		'g', 'o', 'o', 'd', 'b', 'y', 'e', 0x00, // name
	}

	del, err := entry.ReadDelete(bytes.NewReader(input))
	require.NoError(t, err)
	assert.Equal(
		t,
		entry.Delete("goodbye"),
		*del,
	)
}

func Test_ReadDelete_Error(t *testing.T) {
	input := []byte{
		'g', 'o', 'o', 'd', 'b', 'y', 'e', // name, missing terminator
	}

	del, err := entry.ReadDelete(bytes.NewReader(input))
	require.Error(t, err)
	assert.Nil(t, del)
}
