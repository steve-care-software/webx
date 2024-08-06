package singles

import (
	"github.com/steve-care-software/webx/engine/cursors/domain/cursors/namespaces/versions/workspaces/developments"
)

// Builder represents the single builder
type Builder interface {
	Create() Builder
	WithDevelopment(development developments.Development) Builder
	WithWorkspace() Builder
	WithMaster(master Single) Builder
	Now() (Single, error)
}

// Single represents the single
type Single interface {
	Development() developments.Development
	HasBranch() bool
	Branch()
}
