package entry_test

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/jmhobbs/go-raP/entry"
)

func Test_ReadArray(t *testing.T) {
	t.Run("string array", func(t *testing.T) {
		input := []byte{
			'n', 'a', 'm', 'e', 0x00, // name
			0x02,                // count
			0x00,                // value 1 type
			'o', 'n', 'e', 0x00, // value 1
			0x00,                // value 2 type
			't', 'w', 'o', 0x00, // value 2
		}

		arr, err := entry.ReadArray(bytes.NewReader(input))
		require.NoError(t, err)
		t.Log(arr)
		assert.Equal(
			t,
			&entry.Array{
				Name: "name",
				Values: []entry.ArrayValue{
					{
						Type:  entry.ArrayValueTypeString,
						Value: "one",
					},
					{
						Type:  entry.ArrayValueTypeString,
						Value: "two",
					},
				},
			},
			arr,
		)
	})

	t.Run("float array", func(t *testing.T) {
		input := []byte{
			'n', 'a', 'm', 'e', 0x00, // name
			0x02,                   // count
			0x01,                   // value 1 type
			0xa4, 0x70, 0x9d, 0x3f, // 1.23 float32 little-endian
			0x01,                   // value 2 type
			0xa4, 0x70, 0x9d, 0x3f, // 1.23 float32 little-endian
		}

		arr, err := entry.ReadArray(bytes.NewReader(input))
		require.NoError(t, err)
		t.Log(arr)
		assert.Equal(
			t,
			&entry.Array{
				Name: "name",
				Values: []entry.ArrayValue{
					{
						Type:  entry.ArrayValueTypeFloat,
						Value: float32(1.23),
					},
					{
						Type:  entry.ArrayValueTypeFloat,
						Value: float32(1.23),
					},
				},
			},
			arr,
		)
	})

	t.Run("long array", func(t *testing.T) {
		input := []byte{
			'n', 'a', 'm', 'e', 0x00, // name
			0x02,                   // count
			0x02,                   // value 1 type
			0x01, 0x00, 0x00, 0x00, // value 1
			0x02,                   // value 2 type
			0x02, 0x00, 0x00, 0x00, // value 2
		}

		arr, err := entry.ReadArray(bytes.NewReader(input))
		require.NoError(t, err)
		t.Log(arr)
		assert.Equal(
			t,
			&entry.Array{
				Name: "name",
				Values: []entry.ArrayValue{
					{
						Type:  entry.ArrayValueTypeLong,
						Value: int32(1),
					},
					{
						Type:  entry.ArrayValueTypeLong,
						Value: int32(2),
					},
				},
			},
			arr,
		)
	})

	t.Run("array array", func(t *testing.T) {
		input := []byte{
			'n', 'a', 'm', 'e', 0x00, // name
			0x02,                   // count
			0x03,                   // value 1 type
			0x02,                   // nested array 1 count
			0x02,                   // nested array 1 value 1 type
			0x01, 0x00, 0x00, 0x00, // nested array 1 value 1
			0x02,                   // nested array 1 value 2 type
			0x02, 0x00, 0x00, 0x00, // nested array 1 value 2
			0x03,                   // value 2 type
			0x02,                   // nested array 2 count
			0x02,                   // nested array 2 value 1 type
			0x03, 0x00, 0x00, 0x00, // nested array 2 value 1
			0x02,                   // nested array 2 value 2 type
			0x04, 0x00, 0x00, 0x00, // nested array 2 value 2
		}

		arr, err := entry.ReadArray(bytes.NewReader(input))
		require.NoError(t, err)
		t.Log(arr)
		assert.Equal(
			t,
			&entry.Array{
				Name: "name",
				Values: []entry.ArrayValue{
					{
						Type: entry.ArrayValueTypeArray,
						Value: []entry.ArrayValue{
							{
								Type:  entry.ArrayValueTypeLong,
								Value: int32(1),
							},
							{
								Type:  entry.ArrayValueTypeLong,
								Value: int32(2),
							},
						},
					},
					{
						Type: entry.ArrayValueTypeArray,
						Value: []entry.ArrayValue{
							{
								Type:  entry.ArrayValueTypeLong,
								Value: int32(3),
							},
							{
								Type:  entry.ArrayValueTypeLong,
								Value: int32(4),
							},
						},
					},
				},
			},
			arr,
		)

	})

	t.Run("mixed arrray", func(t *testing.T) {
		input := []byte{
			'n', 'a', 'm', 'e', 0x00, // name
			0x04,                // count
			0x00,                // value 1 type (string)
			'o', 'n', 'e', 0x00, // value 1
			0x01,                   // value 2 type (float)
			0xa4, 0x70, 0x9d, 0x3f, // value 2 (1.23 float32 little-endian)
			0x02,                   // value 3 type (long)
			0x39, 0x30, 0x00, 0x00, // value 3 (12345 little-endian)
			0x03,                   // value 4 type
			0x01,                   // nested array count
			0x02,                   // nested array value 1 type
			0x01, 0x00, 0x00, 0x00, // nested array value 1
		}

		arr, err := entry.ReadArray(bytes.NewReader(input))
		require.NoError(t, err)
		t.Log(arr)
		assert.Equal(
			t,
			&entry.Array{
				Name: "name",
				Values: []entry.ArrayValue{
					{
						Type:  entry.ArrayValueTypeString,
						Value: "one",
					},
					{
						Type:  entry.ArrayValueTypeFloat,
						Value: float32(1.23),
					},
					{
						Type:  entry.ArrayValueTypeLong,
						Value: int32(12345),
					},
					{
						Type: entry.ArrayValueTypeArray,
						Value: []entry.ArrayValue{
							{
								Type:  entry.ArrayValueTypeLong,
								Value: int32(1)},
						},
					},
				},
			},
			arr,
		)
	})
}
