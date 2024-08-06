package developments

import "github.com/steve-care-software/webx/engine/cursors/domain/storages/workspaces/developments"

// Development represents a development
type Development interface {
	All() developments.Developments
	HasSingle() bool
	Single()
}
