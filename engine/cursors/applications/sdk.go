package applications

import "github.com/steve-care-software/webx/engine/cursors/applications/sessions"

// Application represents an application
type Application interface {
	Begin(name string) (*uint, error)
	Session(context uint) (sessions.Application, error)
	Commit(context uint) error
	Cancel(context uint) error
}
