package vms

import (
	grammar_applications "github.com/steve-care-software/webx/grammars/applications"
	"github.com/steve-care-software/webx/grammars/domain/grammars"
	"github.com/steve-care-software/webx/grammars/domain/grammars/cardinalities"
	"github.com/steve-care-software/webx/grammars/domain/grammars/values"
	grammar_values "github.com/steve-care-software/webx/grammars/domain/grammars/values"
	interpreter_applications "github.com/steve-care-software/webx/interpreters/applications"
	"github.com/steve-care-software/webx/programs/domain/instructions"
	"github.com/steve-care-software/webx/programs/domain/instructions/applications"
	"github.com/steve-care-software/webx/programs/domain/instructions/attachments"
	instructions_modules "github.com/steve-care-software/webx/programs/domain/instructions/modules"
	"github.com/steve-care-software/webx/programs/domain/instructions/parameters"
	"github.com/steve-care-software/webx/programs/domain/programs/modules"
	"github.com/steve-care-software/webx/selectors/domain/selectors"
)

const (
	// ModuleCastToInt represents the castToInt module
	ModuleCastToInt = iota

	// ModuleCastToUint represents the castToUint module
	ModuleCastToUint

	// ModuleCastToBool represents the castToBool module
	ModuleCastToBool

	// ModuleCastToFloat32 represents the castToFloat32 module
	ModuleCastToFloat32

	// ModuleCastToFloat64 represents the castToFloat32 module
	ModuleCastToFloat64

	// ModuleContainerMapWithKeynames represents the containerMapWithKeynames module
	ModuleContainerMapWithKeynames

	// ModuleContainerList represents the containerList module
	ModuleContainerList

	// ModuleEngineGrammarValue represents the engineGrammarValue module
	ModuleEngineGrammarValue

	// ModuleEngineGrammarCardinality represents the engineGrammarCardinality module
	ModuleEngineGrammarCardinality

	// ModuleEngineGrammarElement represents the engineGrammarElement module
	ModuleEngineGrammarElement

	// ModuleEngineGrammarLine represents the engineGrammarLine module
	ModuleEngineGrammarLine

	// ModuleEngineGrammarBlock represents the engineGrammarBlock module
	ModuleEngineGrammarBlock

	// ModuleEngineGrammarSuite represents the engineGrammarSuite module
	ModuleEngineGrammarSuite

	// ModuleEngineGrammarSuites represents the engineGrammarSuites module
	ModuleEngineGrammarSuites

	// ModuleEngineGrammarToken represents the engineGrammarToken module
	ModuleEngineGrammarToken

	// ModuleEngineGrammarEverything represents the engineGrammarEverything module
	ModuleEngineGrammarEverything

	// ModuleEngineGrammarInstance represents the engineGrammarInstance module
	ModuleEngineGrammarInstance

	// ModuleEngineGrammarExternal represents the engineGrammarExternal module
	ModuleEngineGrammarExternal

	// ModuleEngineGrammarChannelCondition represents the engineGrammarChannelCondition module
	ModuleEngineGrammarChannelCondition

	// ModuleEngineGrammarChannel represents the engineGrammarChannel module
	ModuleEngineGrammarChannel

	// ModuleEngineGrammarChannels represents the engineGrammarChannels module
	ModuleEngineGrammarChannels

	// ModuleEngineGrammar represents the engineGrammar module
	ModuleEngineGrammar

	// ModuleEngineExecuteGrammar represents the engineExecuteGrammar module
	ModuleEngineExecuteGrammar
)

// NewApplication creates a new virtual machine application
func NewApplication() interpreter_applications.Application {
	interpreterAppBuilder := interpreter_applications.NewBuilder(func(name []byte) string {
		return string(name)
	})

	modules := newModules()
	grammar := newGrammar()
	selector := newSelector()
	interpreterApp, err := interpreterAppBuilder.Create().
		WithModules(modules).
		WithGrammar(grammar).
		WithSelector(selector).
		Now()

	if err != nil {
		panic(err)
	}

	return interpreterApp
}

func newModules() modules.Modules {
	// create the cast module funcs:
	castFnsMap := createModuleCast().Execute()

	// create the cast module funcs:
	containersFnsMap := createModuleContainers().Execute()

	// create the grammar module funcs:
	grammarApplication := grammar_applications.NewApplication()
	grammarBuilder := grammars.NewBuilder()
	grammarChannelsBuilder := grammars.NewChannelsBuilder()
	grammarChannelBuilder := grammars.NewChannelBuilder()
	grammarChannelConditionBuilder := grammars.NewChannelConditionBuilder()
	grammarExternalBuilder := grammars.NewExternalBuilder()
	grammarInstanceBuilder := grammars.NewInstanceBuilder()
	grammarEverythingBuilder := grammars.NewEverythingBuilder()
	grammarTokenBuilder := grammars.NewTokenBuilder()
	grammarSuitesBuilder := grammars.NewSuitesBuilder()
	grammarSuiteBuilder := grammars.NewSuiteBuilder()
	grammarBlockBuilder := grammars.NewBlockBuilder()
	grammarLineBuilder := grammars.NewLineBuilder()
	grammarElementBuilder := grammars.NewElementBuilder()
	grammarCardinalityBuilder := cardinalities.NewBuilder()
	grammarValueBuilder := values.NewBuilder()
	grammarFnsMap := createModuleEngineGrammar(
		grammarApplication,
		grammarBuilder,
		grammarChannelsBuilder,
		grammarChannelBuilder,
		grammarChannelConditionBuilder,
		grammarExternalBuilder,
		grammarInstanceBuilder,
		grammarEverythingBuilder,
		grammarTokenBuilder,
		grammarSuitesBuilder,
		grammarSuiteBuilder,
		grammarBlockBuilder,
		grammarLineBuilder,
		grammarElementBuilder,
		grammarCardinalityBuilder,
		grammarValueBuilder,
	).Execute()

	// create the module funcs list:
	moduleFuncs := map[uint]modules.ExecuteFn{}
	for idx, fn := range castFnsMap {
		moduleFuncs[idx] = fn
	}

	for idx, fn := range containersFnsMap {
		moduleFuncs[idx] = fn
	}

	for idx, fn := range grammarFnsMap {
		moduleFuncs[idx] = fn
	}

	// build the modules list:
	modulesList := []modules.Module{}
	moduleBuilder := modules.NewModuleBuilder()
	for idx, oneFunc := range moduleFuncs {
		ins, err := moduleBuilder.Create().WithIndex(uint(idx)).WithFunc(oneFunc).Now()
		if err != nil {
			panic(err)
		}

		modulesList = append(modulesList, ins)
	}

	modulesIns, err := modules.NewBuilder().Create().WithList(modulesList).Now()
	if err != nil {
		panic(err)
	}

	return modulesIns
}

func newGrammar() grammars.Grammar {
	builder := grammars.NewBuilder()
	channelsBuilder := grammars.NewChannelsBuilder()
	channelBuilder := grammars.NewChannelBuilder()
	instanceBuilder := grammars.NewInstanceBuilder()
	everythingBuilder := grammars.NewEverythingBuilder()
	tokensBuilder := grammars.NewTokensBuilder()
	tokenBuilder := grammars.NewTokenBuilder()
	suitesBuilder := grammars.NewSuitesBuilder()
	suiteBuilder := grammars.NewSuiteBuilder()
	blockBuilder := grammars.NewBlockBuilder()
	lineBuilder := grammars.NewLineBuilder()
	elementBuilder := grammars.NewElementBuilder()
	valueBuilder := grammar_values.NewBuilder()
	cardinalityBuilder := cardinalities.NewBuilder()
	grammarIns := createGrammar(
		builder,
		channelsBuilder,
		channelBuilder,
		instanceBuilder,
		everythingBuilder,
		tokensBuilder,
		tokenBuilder,
		suitesBuilder,
		suiteBuilder,
		blockBuilder,
		lineBuilder,
		elementBuilder,
		valueBuilder,
		cardinalityBuilder,
	)

	ins, err := grammarIns.Execute()
	if err != nil {
		panic(err)
	}

	return ins
}

func newSelector() selectors.Selector {
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
	instructionModuleBuilder := instructions_modules.NewBuilder()
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
