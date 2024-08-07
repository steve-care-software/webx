package pointers

import "github.com/steve-care-software/webx/engine/cursors/domain/storages/pointers"

type pointer struct {
	storage pointers.Pointer
	bytes   []byte
}

func createPointer(
	storage pointers.Pointer,
	bytes []byte,
) Pointer {
	out := pointer{
		storage: storage,
		bytes:   bytes,
	}

	return &out
}

// Storage returns the storage
func (obj *pointer) Storage() pointers.Pointer {
	return obj.storage
}

// Bytes returns the bytes
func (obj *pointer) Bytes() []byte {
	return obj.bytes
}
