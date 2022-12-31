package instances

import (
	"github.com/steve-care-software/webx/grammars/domain/grammars"
	"github.com/steve-care-software/webx/grammars/domain/grammars/cardinalities"
	grammar_values "github.com/steve-care-software/webx/grammars/domain/grammars/values"
	"github.com/steve-care-software/webx/programs/domain/instructions"
	"github.com/steve-care-software/webx/programs/domain/instructions/applications"
	"github.com/steve-care-software/webx/programs/domain/instructions/attachments"
	"github.com/steve-care-software/webx/programs/domain/instructions/modules"
	"github.com/steve-care-software/webx/programs/domain/instructions/parameters"
	"github.com/steve-care-software/webx/selectors/domain/selectors"
)

// NewGrammar creates a new gramar instance
func NewGrammar() grammars.Grammar {
	grammarIns := createGrammar(
		grammars.NewBuilder(),
		grammars.NewChannelsBuilder(),
		grammars.NewChannelBuilder(),
		grammars.NewInstanceBuilder(),
		grammars.NewEverythingBuilder(),
		grammars.NewTokensBuilder(),
		grammars.NewTokenBuilder(),
		grammars.NewSuitesBuilder(),
		grammars.NewSuiteBuilder(),
		grammars.NewBlockBuilder(),
		grammars.NewLineBuilder(),
		grammars.NewElementBuilder(),
		grammar_values.NewBuilder(),
		cardinalities.NewBuilder(),
	)

	ins, err := grammarIns.Execute()
	if err != nil {
		panic(err)
	}

	return ins
}

// NewSelector creates a new selector instance
func NewSelector() selectors.Selector {
	builder := selectors.NewBuilder()
	selectorFnBuilder := selectors.NewSelectorFnBuilder()
	tokenBuilder := selectors.NewTokenBuilder()
	elementBuilder := selectors.NewElementBuilder()
	insideBuilder := selectors.NewInsideBuilder()
	fetchersBuilder := selectors.NewFetchersBuilder()
	fetcherBuilder := selectors.NewFetcherBuilder()
	contentFnBuilder := selectors.NewContentFnBuilder()
	instructionsBuilder := instructions.NewBuilder()
	instructionBuilder := instructions.NewInstructionBuilder()
	instructionApplicationBuilder := applications.NewBuilder()
	instructionParameterBuilder := parameters.NewBuilder()
	instructionAttachmentBuilder := attachments.NewBuilder()
	instructionAttachmentVariableBuilder := attachments.NewVariableBuilder()
	instructionAssignmentBuilder := instructions.NewAssignmentBuilder()
	instructionValueBuilder := instructions.NewValueBuilder()
	instructionModuleBuilder := modules.NewBuilder()
	selectorIns := createSelector(
		builder,
		selectorFnBuilder,
		tokenBuilder,
		elementBuilder,
		insideBuilder,
		fetchersBuilder,
		fetcherBuilder,
		contentFnBuilder,
		instructionsBuilder,
		instructionBuilder,
		instructionApplicationBuilder,
		instructionParameterBuilder,
		instructionAttachmentBuilder,
		instructionAttachmentVariableBuilder,
		instructionAssignmentBuilder,
		instructionValueBuilder,
		instructionModuleBuilder,
	)

	ins, err := selectorIns.Execute()
	if err != nil {
		panic(err)
	}

	return ins
}
