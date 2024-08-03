package entries

import (
	"github.com/steve-care-software/webx/engine/bytes/domain/states/pointers/delimiters"
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
func NewEntryForTests(delimiter delimiters.Delimiter, bytes []byte) Entry {
	ins, err := NewEntryBuilder().Create().WithDelimiter(delimiter).WithBytes(bytes).Now()
	if err != nil {
		panic(err)
	}

	return ins
}
