package modifications

import (
	"github.com/steve-care-software/webx/engine/databases/bytes/domain/deletes"
	"github.com/steve-care-software/webx/engine/databases/bytes/domain/entries"
	"github.com/steve-care-software/webx/engine/databases/bytes/domain/states"
)

// NewModificationWithInsertionsForTests creates a new modification with insertions for tests
func NewModificationWithInsertionsForTests(original states.State, insertions entries.Entries) Modification {
	ins, err := NewBuilder().Create().WithOriginal(original).WithInsertions(insertions).Now()
	if err != nil {
		panic(err)
	}

	return ins
}

// NewModificationWithDeletionsForTests creates a new modification with deletions for tests
func NewModificationWithDeletionsForTests(original states.State, deletions deletes.Deletes) Modification {
	ins, err := NewBuilder().Create().WithOriginal(original).WithDeletions(deletions).Now()
	if err != nil {
		panic(err)
	}

	return ins
}

// NewModificationWithInsertionsAndDeletionsForTEsts creates a new modification with insertions for tests
func NewModificationWithInsertionsAndDeletionsForTEsts(original states.State, insertions entries.Entries, deletions deletes.Deletes) Modification {
	ins, err := NewBuilder().Create().WithOriginal(original).WithInsertions(insertions).WithDeletions(deletions).Now()
	if err != nil {
		panic(err)
	}

	return ins
}
