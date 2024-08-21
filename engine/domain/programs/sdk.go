package programs

import (
	"github.com/steve-care-software/webx/engine/domain/nfts"
	"github.com/steve-care-software/webx/engine/domain/programs/grammars"
	"github.com/steve-care-software/webx/engine/domain/programs/instructions"
	instructions_tokens "github.com/steve-care-software/webx/engine/domain/programs/instructions/tokens"
	"github.com/steve-care-software/webx/engine/domain/programs/instructions/tokens/elements"
	instructions_elements "github.com/steve-care-software/webx/engine/domain/programs/instructions/tokens/elements"
	instructions_syscalls "github.com/steve-care-software/webx/engine/domain/programs/instructions/tokens/elements/syscalls"
	instructions_syscalls_values_parameters "github.com/steve-care-software/webx/engine/domain/programs/instructions/tokens/elements/syscalls/parameters"
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
	syscallsBuilder := instructions_syscalls.NewBuilder()
	syscallBuilder := instructions_syscalls.NewSyscallBuilder()
	parametersBuilder := instructions_syscalls_values_parameters.NewBuilder()
	parameterBuilder := instructions_syscalls_values_parameters.NewParameterBuilder()
	return createParserAdapter(
		grammarAdapter,
		builder,
		instructionsBuilder,
		instructionBuilder,
		tokensBuilder,
		tokenBuilder,
		elementsBuilder,
		elementBuilder,
		syscallsBuilder,
		syscallBuilder,
		parametersBuilder,
		parameterBuilder,
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
}

// NFTAdapter represents the program nft adapter
type NFTAdapter interface {
	// ToNFT converts a program instance to an NFT
	ToNFT(program Program) (nfts.NFT, error)

	// ToProgram converts an NFT to a program instance
	ToProgram(nft nfts.NFT) (Program, error)
}

// ComposeAdapter represents the grammar compose adapter
type ComposeAdapter interface {
	// ToBytes takes a program and an element name and returns its bytes
	ToBytes(program Program, elementName string) ([]byte, error)
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
