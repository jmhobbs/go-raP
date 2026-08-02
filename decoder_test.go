package raP_test

import (
	"bytes"
	"testing"

	"github.com/jmhobbs/go-raP"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_ReadAssignment(t *testing.T) {
	t.Run("string", func(t *testing.T) {
		input := []byte{
			0x00,                                    // subtype 0
			'e', 'x', 'a', 'm', 'p', 'l', 'e', 0x00, // name
			'H', 'e', 'l', 'l', 'o', ',', ' ', 'w', 'o', 'r', 'l', 'd', '!', 0x00, // value
		}
		assignment, err := raP.ReadAssignment(bytes.NewReader(input))
		require.NoError(t, err)
		assert.Equal(
			t,
			&raP.Assignment{
				Name:    "example",
				Subtype: raP.AssignmentTypeString,
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
		assignment, err := raP.ReadAssignment(bytes.NewReader(input))
		require.NoError(t, err)
		assert.Equal(
			t,
			&raP.Assignment{
				Name:    "example",
				Subtype: raP.AssignmentTypeVariable,
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
		assignment, err := raP.ReadAssignment(bytes.NewReader(input))
		require.NoError(t, err)
		assert.Equal(
			t,
			&raP.Assignment{
				Name:    "example",
				Subtype: raP.AssignmentTypeFloat,
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
		assignment, err := raP.ReadAssignment(bytes.NewReader(input))
		require.NoError(t, err)
		assert.Equal(
			t,
			&raP.Assignment{
				Name:    "example",
				Subtype: raP.AssignmentTypeLong,
				Value:   int32(12345),
			},
			assignment,
		)
	})
}

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

		arr, err := raP.ReadArray(bytes.NewReader(input))
		require.NoError(t, err)
		t.Log(arr)
		assert.Equal(
			t,
			&raP.Array{
				Name: "name",
				Values: []raP.ArrayValue{
					{
						Type:  raP.ArrayValueTypeString,
						Value: "one",
					},
					{
						Type:  raP.ArrayValueTypeString,
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

		arr, err := raP.ReadArray(bytes.NewReader(input))
		require.NoError(t, err)
		t.Log(arr)
		assert.Equal(
			t,
			&raP.Array{
				Name: "name",
				Values: []raP.ArrayValue{
					{
						Type:  raP.ArrayValueTypeFloat,
						Value: float32(1.23),
					},
					{
						Type:  raP.ArrayValueTypeFloat,
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

		arr, err := raP.ReadArray(bytes.NewReader(input))
		require.NoError(t, err)
		t.Log(arr)
		assert.Equal(
			t,
			&raP.Array{
				Name: "name",
				Values: []raP.ArrayValue{
					{
						Type:  raP.ArrayValueTypeLong,
						Value: int32(1),
					},
					{
						Type:  raP.ArrayValueTypeLong,
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
			0x01, 0x00, 0x00, 0x00, // nested array 1 value 2
			0x03,                   // value 2 type
			0x02,                   // nested array 2 count
			0x02,                   // nested array 2 value 1 type
			0x03, 0x00, 0x00, 0x00, // nested array 2 value 1
			0x02,                   // nested array 2 value 2 type
			0x04, 0x00, 0x00, 0x00, // nested array 2 value 2
		}

		arr, err := raP.ReadArray(bytes.NewReader(input))
		require.NoError(t, err)
		t.Log(arr)
		assert.Equal(
			t,
			&raP.Array{
				Name: "name",
				Values: []raP.ArrayValue{
					{
						Type: raP.ArrayValueTypeArray,
						Value: []raP.ArrayValue{
							{
								Type:  raP.ArrayValueTypeLong,
								Value: int32(1),
							},
							{
								Type:  raP.ArrayValueTypeLong,
								Value: int32(2),
							},
						},
					},
					{
						Type: raP.ArrayValueTypeArray,
						Value: []raP.ArrayValue{
							{
								Type:  raP.ArrayValueTypeLong,
								Value: int32(3),
							},
							{
								Type:  raP.ArrayValueTypeLong,
								Value: int32(3),
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

		arr, err := raP.ReadArray(bytes.NewReader(input))
		require.NoError(t, err)
		t.Log(arr)
		assert.Equal(
			t,
			&raP.Array{
				Name: "name",
				Values: []raP.ArrayValue{
					{
						Type:  raP.ArrayValueTypeString,
						Value: "one",
					},
					{
						Type:  raP.ArrayValueTypeFloat,
						Value: float32(1.23),
					},
					{
						Type:  raP.ArrayValueTypeLong,
						Value: int32(12345),
					},
					{
						Type: raP.ArrayValueTypeArray,
						Value: []raP.ArrayValue{
							{
								Type:  raP.ArrayValueTypeLong,
								Value: int32(1)},
						},
					},
				},
			},
			arr,
		)
	})
}
