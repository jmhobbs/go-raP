package entry

type EntryType uint8

const (
	EntryTypeClass         EntryType = 0
	EntryTypeAssignment    EntryType = 1
	EntryTypeArray         EntryType = 2
	EntryTypeExtern        EntryType = 3
	EntryTypeDelete        EntryType = 4
	EntryTypeArrayWithFlag EntryType = 5
)

type Entry interface {
	Type() EntryType
}
