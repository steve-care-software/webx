package creates

import (
	"github.com/steve-care-software/webx/engine/cursors/domain/hash"
	"github.com/steve-care-software/webx/engine/cursors/domain/loaders/identities/namespaces/versions/workspaces/developments"
)

// Create represents a create
type Create interface {
	Name() string
	Development() developments.Development
	Whitelist() []hash.Hash
}
