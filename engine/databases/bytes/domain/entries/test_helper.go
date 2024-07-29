package entries

import (
	"github.com/steve-care-software/webx/engine/databases/bytes/domain/states/containers/pointers"
)

// NewEntriesForTests creates new entries for tests
func NewEntriesForTests(list []Entry) Entries {
	ins, err := NewBuilder().Create().WithList(list).Now()
	if err != nil {
		panic(err)
	}

	return ins
}

// NewEntryForTests creates a new entry for tests
func NewEntryForTests(pointer pointers.Pointer, bytes []byte) Entry {
	ins, err := NewEntryBuilder().Create().WithPointer(pointer).WithBytes(bytes).Now()
	if err != nil {
		panic(err)
	}

	return ins
}
