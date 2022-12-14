package applications

import (
	"github.com/steve-care-software/webx/blockchains/applications"
	"github.com/steve-care-software/webx/blockchains/domain/cryptography/hash"
	"github.com/steve-care-software/webx/grammars/domain/grammars"
	"github.com/steve-care-software/webx/grammars/domain/grammars/coverages"
	"github.com/steve-care-software/webx/grammars/domain/trees"
)

const (
	// KindGrammar represents the grammar kind
	KindGrammar = iota

	// KindToken represents the token kind
	KindToken

	// KindSuite represents the suite kind
	KindSuite

	// KindElement represents the element kind
	KindElement

	// KindEverything represents the everything kind
	KindEverything

	// KindChannel represents the channel kind
	KindChannel
)

// Builder represents the application builder
type Builder interface {
	Create() Builder
	WithBlockchain(blockchain applications.Application) Builder
	Now() (Application, error)
}

// Application represents a grammar application
type Application interface {
	Database
	Software
}

// Software represents the grammar software application
type Software interface {
	Execute(grammar grammars.Grammar, values []byte) (trees.Tree, error)
	Coverages(grammar grammars.Grammar) (coverages.Coverages, error)
	Covered(coverages coverages.Coverages) (map[string]map[uint]map[uint]string, error)
	Uncovered(grammar grammars.Grammar) (map[string]map[uint]map[uint]string, error)
}

// Database represents the grammar database application
type Database interface {
	Retrieve(context uint, hash hash.Hash) (grammars.Grammar, error)
	Scan(context uint, suites grammars.Suites) (grammars.Grammar, error)
	ScanWithChannels(context uint, suites grammars.Suites, channels grammars.Channels) (grammars.Grammar, error)
	Insert(context uint, grammar grammars.Grammar) error
}
