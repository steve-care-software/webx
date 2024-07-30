package pointers

import "github.com/steve-care-software/webx/engine/databases/bytes/domain/states/containers/pointers/delimiters"

type pointer struct {
	delimiter delimiters.Delimiter
	isDeleted bool
}

func createPointer(
	delimiter delimiters.Delimiter,
) Pointer {
	return createPointerInternally(delimiter, false)
}

func createPointerWithDeleted(
	delimiter delimiters.Delimiter,
) Pointer {
	return createPointerInternally(delimiter, true)
}

func createPointerInternally(
	delimiter delimiters.Delimiter,
	isDeleted bool,
) Pointer {
	out := pointer{
		delimiter: delimiter,
		isDeleted: isDeleted,
	}

	return &out
}

// Delimiter returns the delimiter
func (obj *pointer) Delimiter() delimiters.Delimiter {
	return obj.delimiter
}

// IsDeleted returns true if deleted, false otherwise
func (obj *pointer) IsDeleted() bool {
	return obj.isDeleted
}
