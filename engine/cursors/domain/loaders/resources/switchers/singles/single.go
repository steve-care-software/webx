package singles

import "github.com/steve-care-software/webx/engine/cursors/domain/loaders/resources/storages"

type single struct {
	storage storages.Storage
	bytes   []byte
}

func createSingle(
	storage storages.Storage,
	bytes []byte,
) Single {
	out := single{
		storage: storage,
		bytes:   bytes,
	}

	return &out
}

// Storage returns the storage
func (obj *single) Storage() storages.Storage {
	return obj.storage
}

// Bytes returns the bytes
func (obj *single) Bytes() []byte {
	return obj.bytes
}
