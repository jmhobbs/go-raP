package raP_test

import (
	"bytes"
	"os"
	"testing"

	"github.com/jmhobbs/go-raP"
	"github.com/jmhobbs/go-raP/printer"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_RoundTrip(t *testing.T) {
	expected, err := os.ReadFile("testdata/config.cpp")
	require.NoError(t, err)

	f, err := os.Open("testdata/config.bin")
	require.NoError(t, err)
	t.Cleanup(func() {
		if err := f.Close(); err != nil {
			t.Logf("error closing input file: %v", err)
		}
	})

	root, err := raP.Decode(f)
	require.NoError(t, err)

	var buf bytes.Buffer
	err = printer.New().File(&buf, root)
	require.NoError(t, err)

	assert.Equal(t, string(expected), buf.String())
}
