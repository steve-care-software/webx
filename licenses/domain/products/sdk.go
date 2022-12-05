package products

import (
	"github.com/steve-care-software/webx/blockchains/domain/cryptography/hash"
	"github.com/steve-care-software/webx/grammars/domain/grammars"
	"github.com/steve-care-software/webx/programs/domain/programs"
	"github.com/steve-care-software/webx/selectors/domain/selectors"
)

// Product represents a product
type Product interface {
	Hash() hash.Hash
	IsGrammar() bool
	Grammar() grammars.Grammar
	IsProgram() bool
	Program() programs.Program
	IsSelector() bool
	Selector() selectors.Selector
}
