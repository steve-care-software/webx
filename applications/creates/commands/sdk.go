package commands

import "github.com/steve-care-software/webx/domain/commands"

// Application represents the command application
type Application interface {
	Execute() (commands.Command, error)
}
