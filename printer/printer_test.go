package printer_test

import (
	"bytes"
	"testing"

	"github.com/jmhobbs/go-raP"
	"github.com/jmhobbs/go-raP/printer"
	"github.com/stretchr/testify/assert"
)

func Test_PrintAssignment(t *testing.T) {
	t.Run("string", func(t *testing.T) {
		var buf bytes.Buffer

		printer.New().Assignment(
			0,
			&buf,
			&raP.Assignment{
				Name:    "example",
				Subtype: raP.AssignmentTypeString,
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
			&raP.Assignment{
				Name:    "example",
				Subtype: raP.AssignmentTypeFloat,
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
			&raP.Assignment{
				Name:    "example",
				Subtype: raP.AssignmentTypeLong,
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
			&raP.Assignment{
				Name:    "example",
				Subtype: raP.AssignmentTypeVariable,
				Value:   "anotherVariable",
			},
		)

		assert.Equal(t, `example = anotherVariable;`+"\n", buf.String())
	})
}
