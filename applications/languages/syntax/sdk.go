package syntax

import (
	"github.com/steve-care-software/syntax/domain/syntax/bytes/criterias"
	"github.com/steve-care-software/syntax/domain/syntax/bytes/grammars"
	"github.com/steve-care-software/syntax/domain/syntax/commands"
	"github.com/steve-care-software/syntax/domain/syntax/compilers"
)

// Application represents a language application
type Application interface {
	Grammar(input []byte) (grammars.Grammar, []byte, error)
	Criteria(input []byte) (criterias.Criteria, []byte, error)
	Compiler(input []byte) (compilers.Compiler, []byte, error)
	Commands(input []byte) (commands.Commands, []byte, error)
}
