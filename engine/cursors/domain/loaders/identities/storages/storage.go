package storages

import "github.com/steve-care-software/webx/engine/cursors/domain/loaders/namespaces/singles/versions/singles/workspaces/singles/branches/singles/states/singles/pointers/storages"

type storage struct {
	name    string
	pointer storages.Storage
}

func createStorage(
	name string,
	pointer storages.Storage,
) Storage {
	out := storage{
		name:    name,
		pointer: pointer,
	}

	return &out
}

// Name returns the name
func (obj *storage) Name() string {
	return obj.name
}

// Pointer returns the pointer
func (obj *storage) Pointer() storages.Storage {
	return obj.pointer
}
