package programs

import (
	"github.com/steve-care-software/webx/engine/domain/nfts"
	"github.com/steve-care-software/webx/engine/domain/programs/grammars"
	"github.com/steve-care-software/webx/engine/domain/programs/grammars/rules"
	"github.com/steve-care-software/webx/engine/domain/programs/instructions"
)

// NewParserAdapter creates a new parser adapter
func NewParserAdapter() ParserAdapter {
	grammarAdapter := grammars.NewParserAdapter()
	builder := NewBuilder()
	instructionsBuilder := instructions.NewBuilder()
	instructionBuilder := instructions.NewInstructionBuilder()
	tokensBuilder := instructions.NewTokensBuilder()
	tokenBuilder := instructions.NewTokenBuilder()
	elementsBuilder := instructions.NewElementsBuilder()
	elementBuilder := instructions.NewElementBuilder()
	ruleBuilder := rules.NewRuleBuilder()
	syscallBuilder := instructions.NewSyscallBuilder()
	parametersBuilder := instructions.NewParametersBuilder()
	parameterBuilder := instructions.NewParameterBuilder()
	valueBuilder := instructions.NewValueBuilder()
	referenceBuilder := instructions.NewReferenceBuilder()
	return createParserAdapter(
		grammarAdapter,
		builder,
		instructionsBuilder,
		instructionBuilder,
		tokensBuilder,
		tokenBuilder,
		elementsBuilder,
		elementBuilder,
		ruleBuilder,
		syscallBuilder,
		parametersBuilder,
		parameterBuilder,
		valueBuilder,
		referenceBuilder,
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

	// ToProgramWithRoot creates a program but changes the root block of the grammar
	ToProgramWithRoot(grammar grammars.Grammar, rootBlockName string, input []byte) (Program, []byte, error)
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
	WithRoot(root instructions.Element) Builder
	Now() (Program, error)
}

// Program represents a program
type Program interface {
	Grammar() grammars.Grammar
	Root() instructions.Element
}
