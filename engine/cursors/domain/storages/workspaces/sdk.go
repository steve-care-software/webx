package workspaces

import (
	"github.com/steve-care-software/webx/engine/cursors/domain/storages/originals"
	"github.com/steve-care-software/webx/engine/cursors/domain/storages/workspaces/developments"
	"github.com/steve-care-software/webx/engine/cursors/domain/storages/workspaces/productions"
)

// Adapter represents a workspace adapter
type Adapter interface {
	InstancesToBytes(ins Workspaces) ([]byte, error)
	BytesToInstances(data []byte) (Workspaces, []byte, error)
	InstanceToBytes(ins Workspace) ([]byte, error)
	BytesToInstance(data []byte) (Workspace, []byte, error)
}

// Builder represents a workspaces builder
type Builder interface {
	Create() Builder
	WithList(list []Workspaces) Builder
	Now() (Workspaces, error)
}

// Workspaces represents workspaces
type Workspaces interface {
	List() []Workspace
}

// WorkspaceBuilder represents a workspace builder
type WorkspaceBuilder interface {
	Create() WorkspaceBuilder
	WithOriginal(original originals.Original) WorkspaceBuilder
	WithDevelopments(developments developments.Developments) WorkspaceBuilder
	WithProductions(productions productions.Productions) WorkspaceBuilder
	Now() (Workspace, error)
}

// Workspace represents a workspace
type Workspace interface {
	Original() originals.Original
	HasDevelopments() bool
	Developments() developments.Developments
	HasProductions() bool
	Productions() productions.Productions
}
