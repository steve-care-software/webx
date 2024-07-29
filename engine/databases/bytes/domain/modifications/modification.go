package modifications

import (
	"github.com/steve-care-software/webx/engine/databases/bytes/domain/deletes"
	"github.com/steve-care-software/webx/engine/databases/bytes/domain/entries"
	"github.com/steve-care-software/webx/engine/databases/bytes/domain/states"
)

type modification struct {
	original   states.State
	insertions entries.Entries
	deletions  deletes.Deletes
}

func createModificationWithInsertions(
	original states.State,
	insertions entries.Entries,
) Modification {
	return createModificationInternally(original, insertions, nil)
}

func createModificationWithDeletions(
	original states.State,
	deletions deletes.Deletes,
) Modification {
	return createModificationInternally(original, nil, deletions)
}

func createModificationWithInsertionsAndDeletions(
	original states.State,
	insertions entries.Entries,
	deletions deletes.Deletes,
) Modification {
	return createModificationInternally(original, insertions, deletions)
}

func createModificationInternally(
	original states.State,
	insertions entries.Entries,
	deletions deletes.Deletes,
) Modification {
	out := modification{
		original:   original,
		insertions: insertions,
		deletions:  deletions,
	}

	return &out
}

// Original returns the original state
func (obj *modification) Original() states.State {
	return obj.original
}

// HasInsertions returns true if there is insertions, false otherwise
func (obj *modification) HasInsertions() bool {
	return obj.insertions != nil
}

// Insertions returns the insertion, if any
func (obj *modification) Insertions() entries.Entries {
	return obj.insertions
}

// HasDeletions returns true if there is deletions, false otherwise
func (obj *modification) HasDeletions() bool {
	return obj.deletions != nil
}

// Deletions returns the deletions, if any
func (obj *modification) Deletions() deletes.Deletes {
	return obj.deletions
}
