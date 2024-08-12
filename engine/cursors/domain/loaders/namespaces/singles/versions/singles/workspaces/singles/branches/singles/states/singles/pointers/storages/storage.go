package storages

import "github.com/steve-care-software/webx/engine/cursors/domain/loaders/resources/storages/delimiters"

type storage struct {
	delimiter delimiters.Delimiter
	isDeleted bool
}

func createStorage(
	delimiter delimiters.Delimiter,
) Storage {
	return createStorageInternally(delimiter, false)
}

func createStorageWithDeleted(
	delimiter delimiters.Delimiter,
) Storage {
	return createStorageInternally(delimiter, true)
}

func createStorageInternally(
	delimiter delimiters.Delimiter,
	isDeleted bool,
) Storage {
	out := storage{
		delimiter: delimiter,
		isDeleted: isDeleted,
	}

	return &out
}

// Delimiter returns the delimiter
func (obj *storage) Delimiter() delimiters.Delimiter {
	return obj.delimiter
}

// IsDeleted returns true if deleted, false otherwise
func (obj *storage) IsDeleted() bool {
	return obj.isDeleted
}
