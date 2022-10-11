package grammars

import "github.com/steve-care-software/syntax/domain/syntax/grammars"

// Application represents the grammar create application
type Application interface {
	Execute() (grammars.Grammar, error)
}
