package defaults

import (
	"github.com/steve-care-software/syntax/applications/engines/creates"
	"github.com/steve-care-software/syntax/domain/syntax/commands"
	"github.com/steve-care-software/syntax/domain/syntax/criterias"
	"github.com/steve-care-software/syntax/domain/syntax/grammars"
	"github.com/steve-care-software/syntax/domain/syntax/grammars/cardinalities"
	"github.com/steve-care-software/syntax/domain/syntax/grammars/values"
	grammar_values "github.com/steve-care-software/syntax/domain/syntax/grammars/values"
	"github.com/steve-care-software/syntax/domain/syntax/programs/applications/modules"
)

// NewApplication creates a new create application
func NewApplication() creates.Application {
	return creates.NewApplication(
		createGrammar(
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
		),
		createCommand(
			commands.NewBuilder(),
			commands.NewExecutionBuilder(),
			commands.NewAttachmentBuilder(),
			commands.NewVariableAssignmentBuilder(),
			commands.NewParameterDeclarationBuilder(),
			commands.NewApplicationDeclarationBuilder(),
			commands.NewModuleDeclarationBuilder(),
			criterias.NewBuilder(),
		),
		createModule(
			modules.NewBuilder(),
			modules.NewModuleBuilder(),
			grammars.NewBuilder(),
			grammars.NewChannelsBuilder(),
			grammars.NewChannelBuilder(),
			grammars.NewChannelConditionBuilder(),
			grammars.NewExternalBuilder(),
			grammars.NewInstanceBuilder(),
			grammars.NewEverythingBuilder(),
			grammars.NewTokenBuilder(),
			grammars.NewSuitesBuilder(),
			grammars.NewSuiteBuilder(),
			grammars.NewBlockBuilder(),
			grammars.NewLineBuilder(),
			grammars.NewElementBuilder(),
			cardinalities.NewBuilder(),
			values.NewBuilder(),
		),
	)
}
