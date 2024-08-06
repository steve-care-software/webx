package namespaces

import (
	"github.com/steve-care-software/webx/engine/cursors/domain/storages/delimiters"
	"github.com/steve-care-software/webx/engine/cursors/domain/storages/originals"
)

// Namespaces represents namespaces
type Namespaces interface {
	List() []Namespace
}

// Namespace represents a namespace
type Namespace interface {
	Original() originals.Original
	HasVersions() bool
	Versions() delimiters.Delimiter
}
