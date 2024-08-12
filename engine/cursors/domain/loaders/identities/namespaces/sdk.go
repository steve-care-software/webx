package namespaces

import (
	"github.com/steve-care-software/webx/engine/cursors/domain/loaders/blockchains"
	"github.com/steve-care-software/webx/engine/cursors/domain/storages/versions"
)

// Adapter represents a namespace adapter
type Adapter interface {
	InstancesToBytes(ins Namespaces) ([]byte, error)
	BytesToInstances(data []byte) (Namespaces, error)
	InstanceToBytes(ins Namespaces) ([]byte, error)
	BytesToInstance(data []byte) (Namespace, error)
}

// Namespaces represents namespaces
type Namespaces interface {
	List() []Namespace
}

// NamespaceBuilder represents a namespace builder
type NamespaceBuilder interface {
	Create() NamespaceBuilder
	WithName(name string) NamespaceBuilder
	WithDescription(description string) NamespaceBuilder
	WithBlockchain(blockchain blockchains.Blockchain) NamespaceBuilder
	Now() (Namespace, error)
}

// Namespace represents a namespace namespace
type Namespace interface {
	Name() string
	Description() string
	Blockchain() blockchains.Blockchain
	HasVersions() bool
	Versions() versions.Version
}
