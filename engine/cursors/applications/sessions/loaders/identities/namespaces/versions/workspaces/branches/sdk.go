package branches

import (
	"github.com/steve-care-software/webx/engine/cursors/domain/loaders/identities/namespaces/versions/workspaces/branches/creates"
	"github.com/steve-care-software/webx/engine/cursors/domain/loaders/resources"
)

// Application represents the branch application
type Application interface {
	Create(input resources.Resource, create creates.Create) (resources.Resource, error) // creates a new branch
}
