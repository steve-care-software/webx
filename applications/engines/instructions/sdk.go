package instructions

import (
	"github.com/steve-care-software/syntax/domain/syntax/commands"
	"github.com/steve-care-software/syntax/domain/syntax/grammars"
	"github.com/steve-care-software/syntax/domain/syntax/instructions"
)

// Application represents an instruction application
type Application interface {
	Execute(grammar grammars.Grammar, command commands.Command, script []byte) (instructions.Output, error)
}
