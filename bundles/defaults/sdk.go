package defaults

import (
	"github.com/steve-care-software/syntax/applications/engines/creates"
	"github.com/steve-care-software/syntax/domain/syntax/bytes/grammars"
	"github.com/steve-care-software/syntax/domain/syntax/bytes/grammars/cardinalities"
	grammar_values "github.com/steve-care-software/syntax/domain/syntax/bytes/grammars/values"
	"github.com/steve-care-software/syntax/domain/syntax/commands"
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
		),
		createModule(),
	)
}
