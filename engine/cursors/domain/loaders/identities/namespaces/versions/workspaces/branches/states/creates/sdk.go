package creates

import (
	"github.com/steve-care-software/webx/engine/cursors/domain/hash"
	"github.com/steve-care-software/webx/engine/cursors/domain/loaders/identities/namespaces/versions/workspaces/branches/states/lists"
)

// Builder represents a create state builder
type Builder interface {
	Create() Builder
	WithName(name string) Builder
	WithList(list lists.List) Builder
	WithWhitelist(whitelist []hash.Hash) Builder
	Now() (Create, error)
}

// Create represents a create state
type Create interface {
	Name() string
	List() lists.List
	Whitelist() []hash.Hash
}
