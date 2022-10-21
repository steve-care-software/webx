package instructions

import (
	"github.com/steve-care-software/webx/domain/commands"
	"github.com/steve-care-software/webx/domain/grammars"
	"github.com/steve-care-software/webx/domain/instructions"
)

// Application represents an instruction application
type Application interface {
	Execute(grammar grammars.Grammar, command commands.Command, script []byte) (instructions.Output, error)
}
