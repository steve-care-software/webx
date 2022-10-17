package compilers

import (
	"github.com/steve-care-software/syntax/domain/syntax/compilers"
	"github.com/steve-care-software/syntax/domain/syntax/programs"
)

// Application represents the compiler application
type Application interface {
	Execute(compiler compilers.Compiler, script []byte) (programs.Program, []byte, []byte, error)
}
