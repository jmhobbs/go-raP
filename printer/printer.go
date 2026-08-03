package printer

import (
	"errors"
	"fmt"
	"io"
	"strconv"
	"strings"

	"github.com/jmhobbs/go-raP"
	"github.com/jmhobbs/go-raP/entry"
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

func (p *Printer) indent(level int, out io.Writer) error {
	_, err := fmt.Fprint(out, strings.Repeat(" ", p.indentDepth*level))
	return err
}

func (p *Printer) File(out io.Writer, root *raP.File) error {
	return p.Entries(0, out, root.Entries)
}

func (p *Printer) Entries(indentLevel int, out io.Writer, entries []entry.Entry) error {
	for _, e := range entries {
		switch e.Type() {
		case entry.EntryTypeClass:
			class, ok := e.(*entry.Class)
			if !ok {
				return errors.New("error: Class entry is not of type *Class")
			}
			if err := p.Class(indentLevel, out, class); err != nil {
				return err
			}
		case entry.EntryTypeAssignment:
			assignment, ok := e.(*entry.Assignment)
			if !ok {
				return errors.New("error: Assignment entry is not of type *Assignment")
			}
			if err := p.Assignment(indentLevel, out, assignment); err != nil {
				return err
			}
		case entry.EntryTypeArray:
			array, ok := e.(*entry.Array)
			if !ok {
				return errors.New("error: Array entry is not of type *Array")
			}
			if err := p.Array(indentLevel, out, array); err != nil {
				return err
			}
		case entry.EntryTypeExtern:
			extern, ok := e.(*entry.Extern)
			if !ok {
				return errors.New("error: Extern entry is not of type *Extern")
			}
			if err := p.Extern(indentLevel, out, extern); err != nil {
				return err
			}
		case entry.EntryTypeDelete:
			del, ok := e.(*entry.Delete)
			if !ok {
				return errors.New("error: Delete entry is not of type *Delete")
			}
			if err := p.Delete(indentLevel, out, del); err != nil {
				return err
			}
		case entry.EntryTypeArrayWithFlag:
			array, ok := e.(*entry.ArrayWithFlag)
			if !ok {
				return errors.New("error: ArrayWithFlag entry is not of type *ArrayWithFlag")
			}
			if err := p.ArrayWithFlag(indentLevel, out, array); err != nil {
				return err
			}
		default:
			return fmt.Errorf("unknown entry type: %d", e.Type())
		}
	}

	return nil
}

func (p *Printer) Class(indentLevel int, out io.Writer, class *entry.Class) error {
	var err error
	if err = p.indent(indentLevel, out); err != nil {
		return err
	}
	if _, err = fmt.Fprintf(out, "class %s", class.Name); err != nil {
		return err
	}
	if class.InheritedClassName != "" {
		if _, err = fmt.Fprintf(out, ": %s", class.InheritedClassName); err != nil {
			return err
		}

	}
	if _, err = fmt.Fprintln(out, ""); err != nil {
		return err
	}

	if err = p.indent(indentLevel, out); err != nil {
		return err
	}

	if _, err = fmt.Fprint(out, "{\n"); err != nil {
		return err
	}

	if err = p.Entries(indentLevel+1, out, class.Entries); err != nil {
		return err
	}

	if _, err = fmt.Fprint(out, strings.Repeat("  ", indentLevel)); err != nil {
		return err
	}

	_, err = fmt.Fprint(out, "};\n")
	return err
}

func (p *Printer) Assignment(indentLevel int, out io.Writer, assignment *entry.Assignment) error {
	var err error
	if err = p.indent(indentLevel, out); err != nil {
		return err
	}
	if _, err = fmt.Fprintf(out, "%s = ", assignment.Name); err != nil {
		return err
	}
	switch assignment.Subtype {
	case entry.AssignmentTypeString:
		if _, err = fmt.Fprintf(out, "%q", assignment.Value); err != nil {
			return err
		}
	case entry.AssignmentTypeLong:
		if _, err = fmt.Fprintf(out, "%d", assignment.Value); err != nil {
			return err
		}
	case entry.AssignmentTypeFloat:
		if _, err = fmt.Fprintf(out, "%f", assignment.Value); err != nil {
			return err
		}
	case entry.AssignmentTypeVariable:
		if _, err = fmt.Fprintf(out, "%s", assignment.Value); err != nil {
			return err
		}
	}
	_, err = fmt.Fprint(out, ";\n")
	return err
}

func (p *Printer) Array(indentLevel int, out io.Writer, array *entry.Array) error {
	var err error
	if err = p.indent(indentLevel, out); err != nil {
		return err
	}
	// special case for empty arrays
	if len(array.Values) == 0 {
		_, err = fmt.Fprintf(out, "%s[] = {};\n", array.Name)
		return err
	}
	if _, err = fmt.Fprintf(out, "%s[] = {\n", array.Name); err != nil {
		return err
	}

	if err = p.arrayValues(indentLevel, out, array.Values); err != nil {
		return err
	}

	if err = p.indent(indentLevel, out); err != nil {
		return err
	}
	_, err = fmt.Fprint(out, "};\n")
	return err
}

func (p *Printer) ArrayWithFlag(indentLevel int, out io.Writer, array *entry.ArrayWithFlag) error {
	var err error
	if err = p.indent(indentLevel, out); err != nil {
		return err
	}

	var modifier string
	if array.Flag&int32(0x01) == int32(0x01) {
		modifier = "+"
	}
	// special case for empty arrays
	if len(array.Values) == 0 {
		_, err = fmt.Fprintf(out, "%s[] %s= {};\n", array.Name, modifier)
		return err
	}
	if _, err = fmt.Fprintf(out, "%s[] %s= {\n", array.Name, modifier); err != nil {
		return err
	}

	if err = p.arrayValues(indentLevel, out, array.Values); err != nil {
		return err
	}

	if err = p.indent(indentLevel, out); err != nil {
		return err
	}
	_, err = fmt.Fprint(out, "};\n")
	return err
}

func (p *Printer) arrayValues(indentLevel int, out io.Writer, values []entry.ArrayValue) error {
	valuesAsStrings := make([]string, len(values))
	valueIndent := strings.Repeat("  ", indentLevel+1)
	for i, v := range values {
		switch v.Type {
		case entry.ArrayValueTypeString:
			valuesAsStrings[i] = fmt.Sprintf("%s%q", valueIndent, v.Value)
		case entry.ArrayValueTypeLong:
			valuesAsStrings[i] = fmt.Sprintf("%s%d", valueIndent, v.Value)
		case entry.ArrayValueTypeFloat:
			valuesAsStrings[i] = valueIndent + strconv.FormatFloat(float64(v.Value.(float32)), 'f', -1, 32)
		case entry.ArrayValueTypeArray:
			// TODO
			fallthrough
		case entry.ArrayValueTypeVariable:
			return errors.New("not implemented")
		}
	}
	if _, err := fmt.Fprint(out, strings.Join(valuesAsStrings, ",\n")); err != nil {
		return err
	}
	_, err := fmt.Fprintln(out, "")
	return err
}

func (p *Printer) Extern(indentLevel int, out io.Writer, extern *entry.Extern) error {
	var err error
	if err = p.indent(indentLevel, out); err != nil {
		return err
	}
	_, err = fmt.Fprintf(out, "class %s;\n", string(*extern))
	return err
}

func (p *Printer) Delete(indentLevel int, out io.Writer, del *entry.Delete) error {
	var err error
	if err = p.indent(indentLevel, out); err != nil {
		return err
	}
	// TODO: I have never actually seen this, not sure what the real syntax is
	_, err = fmt.Fprintf(out, "delete %s;\n", string(*del))
	return err
}
