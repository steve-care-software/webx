package states

import "github.com/steve-care-software/webx/engine/cursors/domain/loaders/namespaces/singles/versions/singles/workspaces/singles/branches/singles/states"

// Application represents the state loader application
type Application interface {
	Set(state states.State, name string) (states.State, error)
}
