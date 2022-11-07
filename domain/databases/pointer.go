package databases

import "time"

type pointer struct {
	beginsOn   SizeInBytes
	length     SizeInBytes
	createdOn  time.Time
	references Pointers
}

func createPointer(
	beginsOn SizeInBytes,
	length SizeInBytes,
	createdOn time.Time,
) Pointer {
	return createPointerInternally(beginsOn, length, createdOn, nil)
}

func createPointerWithReferences(
	beginsOn SizeInBytes,
	length SizeInBytes,
	createdOn time.Time,
	references Pointers,
) Pointer {
	return createPointerInternally(beginsOn, length, createdOn, references)
}

func createPointerInternally(
	beginsOn SizeInBytes,
	length SizeInBytes,
	createdOn time.Time,
	references Pointers,
) Pointer {
	out := pointer{
		beginsOn:   beginsOn,
		length:     length,
		createdOn:  createdOn,
		references: references,
	}

	return &out
}

// BeginsOn returns the beginsOn
func (obj *pointer) BeginsOn() SizeInBytes {
	return obj.beginsOn
}

// Length returns the length
func (obj *pointer) Length() SizeInBytes {
	return obj.length
}

// CreatedOn returns the creation time
func (obj *pointer) CreatedOn() time.Time {
	return obj.createdOn
}

// HasReferences returns true if there is references, false otherwise
func (obj *pointer) HasReferences() bool {
	return obj.references != nil
}

// References returns the references, if any
func (obj *pointer) References() Pointers {
	return obj.references
}
