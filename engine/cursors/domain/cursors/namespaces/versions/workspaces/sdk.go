package workspaces

import (
	"github.com/steve-care-software/webx/engine/cursors/domain/cursors/namespaces/versions/workspaces/developments"
	"github.com/steve-care-software/webx/engine/cursors/domain/cursors/namespaces/versions/workspaces/productions"
)

// Workspace represents a workspace
type Workspace interface {
	Name() string
	IsDevelopment() bool
	Development() developments.Development
	IsProduction() bool
	Production() productions.Production
}
