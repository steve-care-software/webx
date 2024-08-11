package applications

import (
	"github.com/steve-care-software/webx/engine/cursors/applications/encryptions"
	"github.com/steve-care-software/webx/engine/cursors/applications/sessions"
)

// Application represents an application
type Application interface {
	Encryption() encryptions.Application
	Begin(name string) (*uint, error)
	Session(context uint) (sessions.Application, error)
	Commit(context uint) error
	Cancel(context uint) error
}
