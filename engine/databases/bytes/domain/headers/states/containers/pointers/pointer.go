package pointers

import "github.com/steve-care-software/webx/engine/databases/bytes/domain/retrievals"

type pointer struct {
	retrieval retrievals.Retrieval
	isDeleted bool
}

func createPointer(
	retrieval retrievals.Retrieval,
) Pointer {
	return createPointerInternally(retrieval, false)
}

func createPointerWithDeleted(
	retrieval retrievals.Retrieval,
) Pointer {
	return createPointerInternally(retrieval, true)
}

func createPointerInternally(
	retrieval retrievals.Retrieval,
	isDeleted bool,
) Pointer {
	out := pointer{
		retrieval: retrieval,
		isDeleted: isDeleted,
	}

	return &out
}

// Retrieval returns the retrieval
func (obj *pointer) Retrieval() retrievals.Retrieval {
	return obj.retrieval
}

// IsDeleted returns true if deleted, false otherwise
func (obj *pointer) IsDeleted() bool {
	return obj.isDeleted
}
