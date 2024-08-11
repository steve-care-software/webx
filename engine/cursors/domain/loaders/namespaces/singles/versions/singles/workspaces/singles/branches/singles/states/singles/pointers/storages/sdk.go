package storages

import (
	"github.com/steve-care-software/webx/engine/cursors/domain/storages/delimiters"
)

// NewBuilder creates a new builder instance
func NewBuilder() Builder {
	return createBuilder()
}

// NewStorageBuilder creates a new storage builder
func NewStorageBuilder() StorageBuilder {
	return createStorageBuilder()
}

// Adapter represents a storages adapter
type Adapter interface {
	InstancesToBytes(ins Storages) ([]byte, error)
	BytesToInstances(data []byte) (Storages, error)
	InstanceToBytes(ins Storage) ([]byte, error)
	BytesToInstance(data []byte) (Storage, error)
}

// Builder represents the storages builder
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
	WithDelimiter(delimiter delimiters.Delimiter) StorageBuilder
	IsDeleted() StorageBuilder
	Now() (Storage, error)
}

// Storage represents a storage
type Storage interface {
	Delimiter() delimiters.Delimiter
	IsDeleted() bool
}
