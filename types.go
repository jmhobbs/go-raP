package raP

func (c Class) Type() EntryType {
	return EntryTypeClass
}

type File struct {
	Entries []Entry
}

type Class struct {
	Name               string
	InheritedClassName string
	Entries            []Entry
}

type Assignment struct {
	Name    string
	Subtype AssignmentType
	Value   any
}

func (a Assignment) Type() EntryType {
	return EntryTypeAssignment
}

type Array struct {
	Name   string
	Values []ArrayValue
}

func (a Array) Type() EntryType {
	return EntryTypeArray
}

type ArrayValue struct {
	Type  ArrayValueType
	Value any
}
