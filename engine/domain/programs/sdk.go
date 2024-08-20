package programs

import (
	"github.com/steve-care-software/webx/engine/domain/nfts"
	"github.com/steve-care-software/webx/engine/domain/programs/grammars"
	"github.com/steve-care-software/webx/engine/domain/programs/instructions"
	"github.com/steve-care-software/webx/engine/domain/programs/instructions/tokens/elements"
)

// ParserAdapter represents the program parser adapter
type ParserAdapter interface {
	// ToProgram takes the grammar and input and converts them to a program instance and the remaining data
	ToProgram(grammar grammars.Grammar, input []byte) (Program, []byte, error)

	// ToBytes takes a program and returns the bytes for the grammar and program
	ToBytes(program Program) ([]byte, []byte, error)
}

// NFTAdapter represents the program nft adapter
type NFTAdapter interface {
	// ToNFT converts a program instance to an NFT
	ToNFT(program Program) (nfts.NFT, error)

	// ToProgram converts an NFT to a program instance
	ToProgram(nft nfts.NFT) (Program, error)
}

// ComposeAdapter represents the program compose adapter
type ComposeAdapter interface {
	// ToBytes takes a program and returns its bytes
	ToBytes(program Program) ([]byte, error)
}

// Builder represents the program builder
type Builder interface {
	Create() Builder
	WithGrammar(grammar grammars.Grammar) Builder
	WithRoot(root elements.Element) Builder
	WithInstructions(instructions instructions.Instructions) Builder
	Now() (Program, error)
}

// Program represents a program
type Program interface {
	Grammar() grammars.Grammar
	Root() elements.Element
	Instructions() instructions.Instructions
}
