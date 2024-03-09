package compilers

import "github.com/steve-care-software/datastencil/domain/libraries"

// Application represents the compiler application
type Application interface {
	Compile(input []byte) (libraries.Library, error)
	Decompile(library libraries.Library) ([]byte, error)
}
