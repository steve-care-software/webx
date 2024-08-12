package creates

import (
	"github.com/steve-care-software/webx/engine/cursors/domain/hash"
	"github.com/steve-care-software/webx/engine/cursors/domain/loaders/identities/namespaces/versions/workspaces/branches"
)

// Builder represents a create branch builder
type Builder interface {
	Create() Builder
	WithName(name string) Builder
	WithBranch(branch branches.Branch) Builder
	WithWhitelist(whitelist []hash.Hash) Builder
	Now() (Create, error)
}

// Create represents a create branch
type Create interface {
	Name() string
	Branch() branches.Branch
	Whitelist() []hash.Hash
}
