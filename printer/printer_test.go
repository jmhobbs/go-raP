package printer_test

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/jmhobbs/go-raP/entry"
	"github.com/jmhobbs/go-raP/printer"
)

func Test_PrinterOptions(t *testing.T) {
	t.Run("indent depth", func(t *testing.T) {
		var buf bytes.Buffer
		err := printer.New(printer.WithIndentDepth(3)).Assignment(
			1,
			&buf,
			&entry.Assignment{
				Name:    "example",
				Subtype: entry.AssignmentTypeString,
				Value:   "Hello, world!",
			},
		)
		require.NoError(t, err)

		assert.Equal(t, `   example = "Hello, world!";`+"\n", buf.String())
	})

	t.Run("indent rune", func(t *testing.T) {
		var buf bytes.Buffer
		err := printer.New(printer.WithIndentRune('-')).Assignment(
			1,
			&buf,
			&entry.Assignment{
				Name:    "example",
				Subtype: entry.AssignmentTypeString,
				Value:   "Hello, world!",
			},
		)
		require.NoError(t, err)

		assert.Equal(t, `--example = "Hello, world!";`+"\n", buf.String())
	})
}

func Test_Assignment(t *testing.T) {
	t.Run("string", func(t *testing.T) {
		var buf bytes.Buffer

		printer.New().Assignment(
			0,
			&buf,
			&entry.Assignment{
				Name:    "example",
				Subtype: entry.AssignmentTypeString,
				Value:   "Hello, world!",
			},
		)

		assert.Equal(t, `example = "Hello, world!";`+"\n", buf.String())
	})

	t.Run("float", func(t *testing.T) {
		var buf bytes.Buffer

		printer.New().Assignment(
			0,
			&buf,
			&entry.Assignment{
				Name:    "example",
				Subtype: entry.AssignmentTypeFloat,
				Value:   float32(1.23),
			},
		)

		assert.Equal(t, `example = 1.23;`+"\n", buf.String())
	})

	t.Run("long", func(t *testing.T) {
		var buf bytes.Buffer

		printer.New().Assignment(
			0,
			&buf,
			&entry.Assignment{
				Name:    "example",
				Subtype: entry.AssignmentTypeLong,
				Value:   int32(23512),
			},
		)

		assert.Equal(t, `example = 23512;`+"\n", buf.String())
	})

	t.Run("variable", func(t *testing.T) {
		var buf bytes.Buffer

		printer.New().Assignment(
			0,
			&buf,
			&entry.Assignment{
				Name:    "example",
				Subtype: entry.AssignmentTypeVariable,
				Value:   "anotherVariable",
			},
		)

		assert.Equal(t, `example = anotherVariable;`+"\n", buf.String())
	})
}

func Test_Array(t *testing.T) {
	var buf bytes.Buffer

	printer.New().Array(
		0,
		&buf,
		&entry.Array{
			Name: "example",
			Values: []entry.ArrayValue{
				{
					Type:  entry.ArrayValueTypeString,
					Value: "Hello",
				},
				{
					Type:  entry.ArrayValueTypeFloat,
					Value: float32(1.23),
				},
				{
					Type:  entry.ArrayValueTypeLong,
					Value: 54321,
				},
			},
		},
	)

	assert.Equal(t, `example[] = {
  "Hello",
  1.23,
  54321
};
`, buf.String())
}

func Test_Extern(t *testing.T) {
	var buf bytes.Buffer

	extern := entry.Extern("example")

	printer.New().Extern(
		0,
		&buf,
		&extern,
	)

	assert.Equal(t, "class example;\n", buf.String())
}

func Test_Delete(t *testing.T) {
	var buf bytes.Buffer

	del := entry.Delete("example")

	printer.New().Delete(
		0,
		&buf,
		&del,
	)

	assert.Equal(t, "delete example;\n", buf.String())
}
