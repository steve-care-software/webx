package languages

import (
	"github.com/steve-care-software/syntax/domain/bytes/criterias"
	"github.com/steve-care-software/syntax/domain/bytes/grammars"
	"github.com/steve-care-software/syntax/domain/commands"
	"github.com/steve-care-software/syntax/domain/compilers"
)

// Application represents a language application
type Application interface {
	Grammar(input []byte) (grammars.Grammar, []byte, error)
	Criteria(input []byte) (criterias.Criteria, []byte, error)
	Compiler(input []byte) (compilers.Compiler, []byte, error)
	Commands(input []byte) (commands.Commands, []byte, error)
}
