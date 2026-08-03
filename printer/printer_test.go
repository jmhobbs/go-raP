package printer_test

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/jmhobbs/go-raP/entry"
	"github.com/jmhobbs/go-raP/printer"
)

func Test_PrintAssignment(t *testing.T) {
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

		assert.Equal(t, `example = 1.230000;`+"\n", buf.String())
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
