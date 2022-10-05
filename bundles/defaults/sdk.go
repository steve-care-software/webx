package defaults

import (
	creates_command "github.com/steve-care-software/syntax/applications/engines/creates/commands"
	creates "github.com/steve-care-software/syntax/applications/engines/creates/grammars"
	creates_module "github.com/steve-care-software/syntax/applications/engines/creates/modules"
	"github.com/steve-care-software/syntax/domain/syntax/bytes/grammars"
	"github.com/steve-care-software/syntax/domain/syntax/bytes/grammars/cardinalities"
	grammar_values "github.com/steve-care-software/syntax/domain/syntax/bytes/grammars/values"
	"github.com/steve-care-software/syntax/domain/syntax/commands"
)

// NewGrammarCreateApplication createsa new grammar create application
func NewGrammarCreateApplication() creates.Application {
	return createGrammar(
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
}

// NewCommandCreateApplication creates a new create command application
func NewCommandCreateApplication() creates_command.Application {
	return createCommand(
		commands.NewBuilder(),
		commands.NewExecutionBuilder(),
		commands.NewAttachmentBuilder(),
		commands.NewVariableAssignmentBuilder(),
		commands.NewParameterDeclarationBuilder(),
		commands.NewApplicationDeclarationBuilder(),
		commands.NewModuleDeclarationBuilder(),
	)
}

// NewModulesCreateApplication creates a new create module application
func NewModulesCreateApplication() creates_module.Application {
	return createModule()
}
