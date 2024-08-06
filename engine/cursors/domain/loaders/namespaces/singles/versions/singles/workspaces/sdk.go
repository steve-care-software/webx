package workspaces

import (
	"github.com/steve-care-software/webx/engine/cursors/domain/loaders/namespaces/singles/versions/singles/workspaces/singles"
	"github.com/steve-care-software/webx/engine/cursors/domain/storages/workspaces"
)

// Workspace represents a workspace
type Workspace interface {
	All() workspaces.Workspaces
	HasCurrent() bool
	Current() singles.Single
}
