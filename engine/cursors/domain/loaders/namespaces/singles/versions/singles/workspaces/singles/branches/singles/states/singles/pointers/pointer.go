package pointers

import "github.com/steve-care-software/webx/engine/cursors/domain/loaders/namespaces/singles/versions/singles/workspaces/singles/branches/singles/states/singles/pointers/storages"

type pointer struct {
	storage storages.Storage
	bytes   []byte
}

func createPointer(
	storage storages.Storage,
	bytes []byte,
) Pointer {
	out := pointer{
		storage: storage,
		bytes:   bytes,
	}

	return &out
}

// Storage returns the storage
func (obj *pointer) Storage() storages.Storage {
	return obj.storage
}

// Bytes returns the bytes
func (obj *pointer) Bytes() []byte {
	return obj.bytes
}
