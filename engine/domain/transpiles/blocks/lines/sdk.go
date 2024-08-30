package lines

import (
	"github.com/steve-care-software/webx/engine/domain/programs/grammars/blocks/suites"
	"github.com/steve-care-software/webx/engine/domain/transpiles/blocks/lines/tokens"
)

// Lines represents lines
type Lines interface {
	List() []Line
}

// Line represents a line
type Line interface {
	Tokens() tokens.Tokens
	HasSuites() bool
	Suites() suites.Suites
}
