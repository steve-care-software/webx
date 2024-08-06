package namespaces

import "github.com/steve-care-software/webx/engine/cursors/domain/storages/headers/namespaces"

// Namespace represents a namespace
type Namespace interface {
	HasNamespaces() bool
	Namespaces() namespaces.Namespaces
}
