package creates

import (
	"github.com/steve-care-software/webx/engine/cursors/domain/hash"
	"github.com/steve-care-software/webx/engine/cursors/domain/loaders/identities/namespaces/versions/workspaces/productions"
)

// Create represents a create
type Create interface {
	Name() string
	Development() productions.Production
	Whitelist() []hash.Hash
}
