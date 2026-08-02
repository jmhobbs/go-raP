package raP

type EntryType uint8

const (
	EntryTypeClass         EntryType = 0
	EntryTypeAssignment    EntryType = 1
	EntryTypeArray         EntryType = 2
	EntryTypeExtern        EntryType = 3
	EntryTypeDelete        EntryType = 4
	EntryTypeArrayWithFlag EntryType = 5
)

type AssignmentType uint8

const (
	AssignmentTypeString AssignmentType = 0
	AssignmentTypeFloat  AssignmentType = 1
	AssignmentTypeLong   AssignmentType = 2
)

type ArrayValueType uint8

const (
	ArrayValueTypeString   ArrayValueType = 0
	ArrayValueTypeFloat    ArrayValueType = 1
	ArrayValueTypeLong     ArrayValueType = 2
	ArrayValueTypeArray    ArrayValueType = 3
	ArrayValueTypeVariable ArrayValueType = 4
)
