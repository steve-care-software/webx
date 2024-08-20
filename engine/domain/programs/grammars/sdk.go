package grammars

import (
	"github.com/steve-care-software/webx/engine/domain/nfts"
	"github.com/steve-care-software/webx/engine/domain/programs/grammars/blocks"
	"github.com/steve-care-software/webx/engine/domain/programs/grammars/blocks/lines/tokens/elements"
	"github.com/steve-care-software/webx/engine/domain/programs/grammars/rules"
	"github.com/steve-care-software/webx/engine/domain/programs/grammars/syscalls"
)

// NewBuilder creates a new builder
func NewBuilder() Builder {
	return createBuilder()
}

// ParserAdapter represents the grammar parser adapter
type ParserAdapter interface {
	// ToGrammar takes the input and converts it to a grammar instance and the remaining data
	ToGrammar(input []byte) (Grammar, []byte, error)

	// ToBytes takes a grammar and returns the bytes
	ToBytes(grammar Grammar) ([]byte, error)
}

// NFTAdapter represents the grammar nft adapter
type NFTAdapter interface {
	// ToNFT converts a grammar instance to an NFT
	ToNFT(grammar Grammar) (nfts.NFT, error)

	// ToGrammar converts an NFT to a grammar instance
	ToGrammar(nft nfts.NFT) (Grammar, error)
}

// ComposeAdapter represents the grammar compose adapter
type ComposeAdapter interface {
	// ToBytes takes a grammar and a blockname and returns its bytes
	ToBytes(grammar Grammar, blockName string) ([]byte, error)
}

// Builder represents the grammar builder
type Builder interface {
	Create() Builder
	WithVersion(version uint) Builder
	WithRoot(root elements.Element) Builder
	WithRules(rules rules.Rules) Builder
	WithBlocks(blocks blocks.Blocks) Builder
	WithSyscalls(syscalls syscalls.Syscalls) Builder
	WithOmissions(omissions elements.Elements) Builder
	Now() (Grammar, error)
}

// Grammar represents a grammar
type Grammar interface {
	Version() uint
	Root() elements.Element
	Rules() rules.Rules
	Blocks() blocks.Blocks
	HasSyscalls() bool
	Syscalls() syscalls.Syscalls
	HasOmissions() bool
	Omissions() elements.Elements
}
