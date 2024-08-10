package storages

import "github.com/steve-care-software/webx/engine/cursors/domain/loaders/namespaces/singles/versions/singles/workspaces/singles/branches/singles/states/singles/pointers/storages"

// NewBuilder creates a new builder
func NewBuilder() Builder {
	return createBuilder()
}

// NewStorageBuilder creates a new storage builder
func NewStorageBuilder() StorageBuilder {
	return createStorageBuilder()
}

// Adapter represents the storage adapter
type Adapter interface {
	InstancesToBytes(ins Storages) ([]byte, error)
	BytesToInstances(data []byte) (Storages, error)
	InstanceToBytes(ins Storage) ([]byte, error)
	BytesToInstance(data []byte) (Storage, error)
}

// Builder represents a builder
type Builder interface {
	Create() Builder
	WithList(list []Storage) Builder
	Now() (Storages, error)
}

// Storages represents storages
type Storages interface {
	List() []Storage
}

// StorageBuilder represents a storage builder
type StorageBuilder interface {
	Create() StorageBuilder
	WithName(name string) StorageBuilder
	WithPointer(pointer storages.Storage) StorageBuilder
	Now() (Storage, error)
}

// Storage represents an identity storage
type Storage interface {
	Name() string
	Pointer() storages.Storage
}
