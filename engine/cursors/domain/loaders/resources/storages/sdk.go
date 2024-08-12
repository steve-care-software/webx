package storages

import (
	"github.com/steve-care-software/webx/engine/cursors/domain/hash"
	"github.com/steve-care-software/webx/engine/cursors/domain/storages/delimiters"
)

// Adapter represents a storage adapter
type Adapter interface {
	InstancesToBytes(ins Storages) ([]byte, error)
	BytesToInstances(data []byte) (Storages, error)
	InstanceToBytes(ins Storage) ([]byte, error)
	BytesToInstance(data []byte) (Storage, error)
}

// Builder represents the builder
type Builder interface {
	Create() Builder
	WithList(list []Storage) Builder
	Now() (Storages, error)
}

// Storages represents storages
type Storages interface {
	List() []Storage
	FetchByName(name string) (Storage, error)
	NextIndex() uint64
}

// StorageBuilder represents the storage builder
type StorageBuilder interface {
	Create() StorageBuilder
	WithName(name string) StorageBuilder
	WithDelimiter(delimiter delimiters.Delimiter) StorageBuilder
	WithWhitelist(whitelist []hash.Hash) StorageBuilder
	IsDeleted() StorageBuilder
	Now() (Storage, error)
}

// Storage represents the storage resource
type Storage interface {
	Name() string
	Delimiter() delimiters.Delimiter
	IsDeleted() bool
	Whitelist() []hash.Hash
}
