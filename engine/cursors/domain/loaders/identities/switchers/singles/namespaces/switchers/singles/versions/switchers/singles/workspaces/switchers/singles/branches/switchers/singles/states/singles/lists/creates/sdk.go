package creates

import (
	"github.com/steve-care-software/webx/engine/cursors/domain/hash"
	"github.com/steve-care-software/webx/engine/cursors/domain/loaders/identities/switchers/singles/namespaces/switchers/singles/versions/switchers/singles/workspaces/switchers/singles/branches/switchers/singles/states/singles/lists"
)

// Builder represents a create list builder
type Builder interface {
	Create() Builder
	WithList(list lists.List) Builder
	WithWhitelist(whitelist []hash.Hash) Builder
	Now() (Create, error)
}

// Create represents a create list
type Create interface {
	List() lists.List
	Whitelist() []hash.Hash
}
