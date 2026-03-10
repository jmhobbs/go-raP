package raP_test

import (
	"bytes"
	"testing"

	"github.com/jmhobbs/go-raP"
	"github.com/stretchr/testify/require"
)

func Test_Decode(t *testing.T) {
	t.Run("valid", func(t *testing.T) {
		buf := bytes.NewBuffer([]byte{
			0, 'r', 'a', 'P',
			4, 0, 0,
			0,
		})
		require.NoError(t, raP.Decode(buf, nil))
	})

	t.Run("invalid header", func(t *testing.T) {
		buf := bytes.NewBuffer([]byte{0, 'r', 'a', 'p'})
		require.Error(t, raP.Decode(buf, nil))
	})

	t.Run("short header read", func(t *testing.T) {
		buf := bytes.NewBuffer([]byte{0, 'r', 'a'})
		require.Error(t, raP.Decode(buf, nil))
	})

	t.Run("invalid type", func(t *testing.T) {
		buf := bytes.NewBuffer([]byte{
			0, 'r', 'a', 'P',
			5, 0, 0,
		})
		require.Error(t, raP.Decode(buf, nil))
	})

	t.Run("short type read", func(t *testing.T) {
		buf := bytes.NewBuffer([]byte{
			0, 'r', 'a', 'P',
			4, 0,
		})
		require.Error(t, raP.Decode(buf, nil))
	})

	t.Run("invalid packet type", func(t *testing.T) {
		buf := bytes.NewBuffer([]byte{
			0, 'r', 'a', 'P',
			4, 0, 0,
			7,
		})
		require.Error(t, raP.Decode(buf, nil))
	})

}
