package commands

import "github.com/steve-care-software/syntax/domain/syntax/commands"

// Application represents the command application
type Application interface {
	Execute() (commands.Command, error)
}
