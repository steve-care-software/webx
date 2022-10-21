package grammars

import "github.com/steve-care-software/webx/domain/grammars"

// Application represents the grammar create application
type Application interface {
	Execute() (grammars.Grammar, error)
}
