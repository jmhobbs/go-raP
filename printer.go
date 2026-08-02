package raP

import (
	"errors"
	"fmt"
	"io"
	"strings"
)

func Printer(out io.Writer, root *File) error {
	return PrintEntries(0, out, root.Entries)
}

func indent(level int, out io.Writer) {
	for range level {
		fmt.Fprint(out, "  ")
	}
}

func PrintEntries(indentLevel int, out io.Writer, entries []Entry) error {
	for _, entry := range entries {
		switch entry.Type() {
		case EntryTypeClass:
			class, ok := entry.(*Class)
			if !ok {
				return errors.New("error: Class entry is not of type *Class")
			}
			PrintClass(indentLevel, out, class)
		case EntryTypeAssignment:
			assignment, ok := entry.(*Assignment)
			if !ok {
				return errors.New("error: Assignment entry is not of type *Assignment")
			}
			PrintAssignment(indentLevel, out, assignment)
		case EntryTypeArray:
			array, ok := entry.(*Array)
			if !ok {
				return errors.New("error: Array entry is not of type *Array")
			}
			PrintArray(indentLevel, out, array)
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

func PrintClass(indentLevel int, out io.Writer, class *Class) {
	indent(indentLevel, out)
	fmt.Fprintf(out, "class %s", class.Name)
	if class.InheritedClassName != "" {
		fmt.Fprintf(out, ": %s", class.InheritedClassName)
	}
	fmt.Fprintln(out, "")
	indent(indentLevel, out)
	fmt.Fprint(out, "{\n")
	PrintEntries(indentLevel+1, out, class.Entries)
	for range indentLevel {
		fmt.Fprint(out, "  ")
	}
	fmt.Fprint(out, "};\n")
}

func PrintAssignment(indentLevel int, out io.Writer, assignment *Assignment) {
	indent(indentLevel, out)
	fmt.Fprintf(out, "%s = ", assignment.Name)
	switch assignment.Subtype {
	case AssignmentTypeString:
		fmt.Fprintf(out, "%q", assignment.Value)
	case AssignmentTypeLong:
		fmt.Fprintf(out, "%d", assignment.Value)
	case AssignmentTypeFloat:
		fmt.Fprintf(out, "%f", assignment.Value)
	}
	fmt.Fprint(out, ";\n")
}

func PrintArray(indentLevel int, out io.Writer, array *Array) {
	indent(indentLevel, out)
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
		case ArrayValueTypeString:
			valuesAsStrings[i] = fmt.Sprintf("%s%q", valueIndent, v.Value)
		case ArrayValueTypeLong:
			valuesAsStrings[i] = fmt.Sprintf("%s%d", valueIndent, v.Value)
		case ArrayValueTypeFloat:
			valuesAsStrings[i] = fmt.Sprintf("%s%f", valueIndent, v.Value)
		}
	}
	fmt.Fprint(out, strings.Join(valuesAsStrings, ",\n")+"\n")
	indent(indentLevel, out)
	fmt.Fprint(out, "};\n")
}
