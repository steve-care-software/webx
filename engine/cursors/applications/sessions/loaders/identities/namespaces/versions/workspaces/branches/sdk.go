package branches

import (
	"github.com/steve-care-software/webx/engine/cursors/applications/sessions/loaders/identities/namespaces/versions/workspaces/branches/states"
	"github.com/steve-care-software/webx/engine/cursors/domain/loaders/identities/namespaces/versions/workspaces/branches/creates"
	"github.com/steve-care-software/webx/engine/cursors/domain/loaders/resources"
)

// Application represents the branch application
type Application interface {
	Create(input resources.Resource, create creates.Create) (resources.Resource, error) // creates a new branch
	List(input resources.Resource) ([]string, error)                                    // list all the sub branches of the current branch
	Select(input resources.Resource, name string) (resources.Resource, error)           // selects the branch by name
	Merge(input resources.Resource) (resources.Resource, error)                         // merges the current branch into its parent and set the parent as the current branch
	Clone(input resources.Resource, create creates.Create) (resources.Resource, error)  // clones the current branch into a children branch
	Delete(input resources.Resource) (resources.Resource, error)                        // deletes the current branch
	Recover(input resources.Resource) (resources.Resource, error)                       // recovers the current branch
	States(input resources.Resource) (states.Application, error)                        // returns the state application of the current branch
}
