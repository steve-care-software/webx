package cursors

import (
	"github.com/steve-care-software/webx/engine/bytes/domain/branches"
	"github.com/steve-care-software/webx/engine/bytes/domain/iterations"
	"github.com/steve-care-software/webx/engine/bytes/domain/namespaces"
	"github.com/steve-care-software/webx/engine/bytes/domain/states"
	"github.com/steve-care-software/webx/engine/bytes/domain/versions"
)

// Cursor represents a cursor
type Cursor interface {
	Index() Index
	HasNamespace() bool
	Namespace() Namespace
	HasVersion() bool
	Version() Version
	HasIteration() bool
	Iteration() Iteration
	HasWorkspace() bool
	Workspace() Workspace
	HasBranch() bool
	Branch() Branch
	HasState() bool
	State() State
}

// Index represents the index
type Index interface {
	Namespaces() uint64
	Versions() uint64
	Iterations() uint64
	Workspaces() uint64
	Branches() uint64
	States() uint64
}

// Namespace represents a namespace index
type Namespace interface {
	Length() uint64
	All() namespaces.Namespaces
	HasCurrent() bool
	Current() namespaces.Namespace
}

// Version represents a version index
type Version interface {
	Length() uint64
	All() versions.Versions
	HasCurrent() bool
	Current() versions.Version
}

// Iteration represents an iteration index
type Iteration interface {
	Length() uint64
	All() iterations.Iterations
	HasCurrent() bool
	Current() iterations.Iteration
}

// Workspace represents a workspace index
type Workspace interface {
	Length() uint64
	Branch()
	HasCurrent() bool
	Current()
}

// Branch represents a branch index
type Branch interface {
	Length() uint64
	All() branches.Branches
	HasCurrent() bool
	Current() branches.Branch
}

// State represents a state index
type State interface {
	Length() uint64
	All() states.States
	HasCurrent() bool
	Current() states.State
}
