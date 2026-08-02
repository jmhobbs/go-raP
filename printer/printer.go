package printer

import (
	"errors"
	"fmt"
	"io"
	"strings"

	"github.com/jmhobbs/go-raP"
)

type Printer struct {
	indentDepth int
}

type PrinterOption func(*Printer)

func WithIndentDepth(depth int) PrinterOption {
	return func(p *Printer) {
		p.indentDepth = depth
	}
}

func New(opts ...PrinterOption) *Printer {
	p := &Printer{
		indentDepth: 2,
	}
	for _, opt := range opts {
		opt(p)
	}
	return p
}

func (p *Printer) indent(level int, out io.Writer) {
	fmt.Fprint(out, strings.Repeat(" ", p.indentDepth*level))
}

func (p *Printer) File(out io.Writer, root *raP.File) error {
	return p.Entries(0, out, root.Entries)
}

func (p *Printer) Entries(indentLevel int, out io.Writer, entries []raP.Entry) error {
	for _, entry := range entries {
		switch entry.Type() {
		case raP.EntryTypeClass:
			class, ok := entry.(*raP.Class)
			if !ok {
				return errors.New("error: Class entry is not of type *Class")
			}
			p.Class(indentLevel, out, class)
		case raP.EntryTypeAssignment:
			assignment, ok := entry.(*raP.Assignment)
			if !ok {
				return errors.New("error: Assignment entry is not of type *Assignment")
			}
			p.Assignment(indentLevel, out, assignment)
		case raP.EntryTypeArray:
			array, ok := entry.(*raP.Array)
			if !ok {
				return errors.New("error: Array entry is not of type *Array")
			}
			p.Array(indentLevel, out, array)
		default:
			return fmt.Errorf("unknown entry type: %d", entry.Type())
			// todo:
			// extern
			// delete
			// array with flag
		}
	}

	return nil
}

func (p *Printer) Class(indentLevel int, out io.Writer, class *raP.Class) {
	p.indent(indentLevel, out)
	fmt.Fprintf(out, "class %s", class.Name)
	if class.InheritedClassName != "" {
		fmt.Fprintf(out, ": %s", class.InheritedClassName)
	}
	fmt.Fprintln(out, "")
	p.indent(indentLevel, out)
	fmt.Fprint(out, "{\n")
	p.Entries(indentLevel+1, out, class.Entries)
	for range indentLevel {
		fmt.Fprint(out, "  ")
	}
	fmt.Fprint(out, "};\n")
}

func (p *Printer) Assignment(indentLevel int, out io.Writer, assignment *raP.Assignment) {
	p.indent(indentLevel, out)
	fmt.Fprintf(out, "%s = ", assignment.Name)
	switch assignment.Subtype {
	case raP.AssignmentTypeString:
		fmt.Fprintf(out, "%q", assignment.Value)
	case raP.AssignmentTypeLong:
		fmt.Fprintf(out, "%d", assignment.Value)
	case raP.AssignmentTypeFloat:
		fmt.Fprintf(out, "%f", assignment.Value)
	}
	fmt.Fprint(out, ";\n")
}

func (p *Printer) Array(indentLevel int, out io.Writer, array *raP.Array) {
	p.indent(indentLevel, out)
	// special case for empty arrays
	if len(array.Values) == 0 {
		fmt.Fprintf(out, "%s[] = {};\n", array.Name)
		return
	}
	fmt.Fprintf(out, "%s[] = {\n", array.Name)
	valuesAsStrings := make([]string, len(array.Values))
	valueIndent := strings.Repeat("  ", indentLevel+1)
	for i, v := range array.Values {
		switch v.Type {
		case raP.ArrayValueTypeString:
			valuesAsStrings[i] = fmt.Sprintf("%s%q", valueIndent, v.Value)
		case raP.ArrayValueTypeLong:
			valuesAsStrings[i] = fmt.Sprintf("%s%d", valueIndent, v.Value)
		case raP.ArrayValueTypeFloat:
			valuesAsStrings[i] = fmt.Sprintf("%s%f", valueIndent, v.Value)
		}
	}
	fmt.Fprint(out, strings.Join(valuesAsStrings, ",\n")+"\n")
	p.indent(indentLevel, out)
	fmt.Fprint(out, "};\n")
}
