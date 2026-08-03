package entry_test

import (
	"bytes"
	"testing"

	"github.com/jmhobbs/go-raP/entry"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_ReadExtern(t *testing.T) {
	input := []byte{
		'o', 'u', 't', 's', 'i', 'd', 'e', 0x00, // name
	}

	extern, err := entry.ReadExtern(bytes.NewReader(input))
	require.NoError(t, err)
	assert.Equal(
		t,
		entry.Extern("outside"),
		*extern,
	)
}

func Test_ReadExtern_Error(t *testing.T) {
	input := []byte{
		'o', 'u', 't', 's', 'i', 'd', 'e', // name, missing terminator
	}

	extern, err := entry.ReadExtern(bytes.NewReader(input))
	require.Error(t, err)
	assert.Nil(t, extern)
}
