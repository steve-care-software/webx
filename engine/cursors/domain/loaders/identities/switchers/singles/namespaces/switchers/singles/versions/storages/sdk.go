package storages

import "github.com/steve-care-software/webx/engine/cursors/domain/loaders/namespaces/singles/versions/singles/workspaces/singles/branches/singles/states/singles/pointers/storages"

// Builder represents a storages builder
type Builder interface {
	Create() Builder
	WithList(list []Storage) Builder
	Now() (Storages, error)
}

// Storages represents storages
type Storages interface {
	List() []Storage
	Names() []string
	Fetch(name string) (Storage, error)
}

// StorageBuilder represents a storage builder
type StorageBuilder interface {
	Create() StorageBuilder
	WithName(name string) StorageBuilder
	WithPointer(pointer storages.Storage) StorageBuilder
	Now() (Storage, error)
}

// Storage represents a namespace storage
type Storage interface {
	Name() string
	Pointer() storages.Storage
}
