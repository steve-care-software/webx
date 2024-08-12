package developments

import (
	"github.com/steve-care-software/webx/engine/cursors/applications/sessions/loaders/identities/namespaces/versions/workspaces/branches"
	"github.com/steve-care-software/webx/engine/cursors/domain/loaders/identities/namespaces/versions/workspaces/developments/creates"
	"github.com/steve-care-software/webx/engine/cursors/domain/loaders/resources"
)

// Application represents a development application
type Application interface {
	Create(input resources.Resource, create creates.Create) (resources.Resource, error)   // creates a new branch
	List(input resources.Resource) ([]string, error)                                      // list all the developments
	Select(input resources.Resource, name string) (resources.Resource, error)             // selects the development by name
	Move(input resources.Resource, name string) (resources.Resource, error)               // moves the development branch in production
	Delete(input resources.Resource) (resources.Resource, error)                          // deletes the current development
	Recover(input resources.Resource) (resources.Resource, error)                         // recovers the current development
	Branch(input resources.Resource, create creates.Create) (branches.Application, error) // returns the core branch application
}
