package programs

import (
	"github.com/steve-care-software/webx/engine/domain/nfts"
	"github.com/steve-care-software/webx/engine/domain/programs/grammars"
	"github.com/steve-care-software/webx/engine/domain/programs/instructions"
	instructions_tokens "github.com/steve-care-software/webx/engine/domain/programs/instructions/tokens"
	"github.com/steve-care-software/webx/engine/domain/programs/instructions/tokens/elements"
	instructions_elements "github.com/steve-care-software/webx/engine/domain/programs/instructions/tokens/elements"
	"github.com/steve-care-software/webx/engine/domain/programs/syscalls"
)

// NewParserAdapter creates a new parser adapter
func NewParserAdapter() ParserAdapter {
	grammarAdapter := grammars.NewParserAdapter()
	builder := NewBuilder()
	instructionsBuilder := instructions.NewBuilder()
	instructionBuilder := instructions.NewInstructionBuilder()
	tokensBuilder := instructions_tokens.NewBuilder()
	tokenBuilder := instructions_tokens.NewTokenBuilder()
	elementsBuilder := instructions_elements.NewBuilder()
	elementBuilder := instructions_elements.NewElementBuilder()
	return createParserAdapter(
		grammarAdapter,
		builder,
		instructionsBuilder,
		instructionBuilder,
		tokensBuilder,
		tokenBuilder,
		elementsBuilder,
		elementBuilder,
	)
}

// NewBuilder creates a new builder
func NewBuilder() Builder {
	return createBuilder()
}

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

// Builder represents the program builder
type Builder interface {
	Create() Builder
	WithGrammar(grammar grammars.Grammar) Builder
	WithRoot(root elements.Element) Builder
	WithInstructions(instructions instructions.Instructions) Builder
	WithSyscalls(syscalls syscalls.Syscalls) Builder
	Now() (Program, error)
}

// Program represents a program
type Program interface {
	Grammar() grammars.Grammar
	Root() elements.Element
	Instructions() instructions.Instructions
	HasSyscalls() bool
	Syscalls() syscalls.Syscalls
}
