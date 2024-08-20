package programs

import "github.com/steve-care-software/webx/engine/domain/grammars"

// Program represents a program
type Program interface {
	Grammar() grammars.Grammar
	Version() uint
}
