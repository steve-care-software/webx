package states

import (
	"github.com/steve-care-software/webx/engine/cursors/applications/sessions/loaders/identities/namespaces/versions/workspaces/branches/states/entities"
	"github.com/steve-care-software/webx/engine/cursors/applications/sessions/loaders/identities/namespaces/versions/workspaces/branches/states/lists"
	"github.com/steve-care-software/webx/engine/cursors/domain/loaders/identities/namespaces/versions/workspaces/branches/states/creates"
	"github.com/steve-care-software/webx/engine/cursors/domain/loaders/resources"
)

const (
	// FlagResource represents the list flag
	FlagResource (uint8) = iota

	// FlagList represents the resource flag
	FlagList

	// FlagEntity represents the tntity flag
	FlagEntity
)

const (
	// KindDeleted represents the deleted kind
	KindDeleted (uint8) = iota

	// KindActive represents the active kind
	KindActive
)

// Application represents the state loader application
type Application interface {
	Create(input resources.Resource, create creates.Create) (resources.Resource, error) // creates a new state
	List(input resources.Resource, kind uint8) ([]string, error)                        // list all the states of the branch
	Select(input resources.Resource, name string) (resources.Resource, error)           // selects the state by name
	Set(input resources.Resource, flag uint8) (resources.Resource, error)               // sets the pointer to the resource, list or entity
	Flag() (*uint, error)                                                               // returns the current selected flag
	Kind() (*uint, error)                                                               // returns the current selected kind
	Delete(input resources.Resource) (resources.Resource, error)                        // deletes the current state
	Recover(input resources.Resource) (resources.Resource, error)                       // recovers the current state
	Entities(input resources.Resource) (entities.Application, error)                    // returns the entity application with the current state's entities
	Lists(input resources.Resource) (lists.Application, error)                          // returns the list application with the current state's lists
}
