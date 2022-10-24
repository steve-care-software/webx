package defaults

import (
	compiler_applications "github.com/steve-care-software/webx/applications/compilers"
	"github.com/steve-care-software/webx/applications/creates"
	identity_applications "github.com/steve-care-software/webx/applications/identities"
	"github.com/steve-care-software/webx/domain/commands"
	"github.com/steve-care-software/webx/domain/compilers"
	"github.com/steve-care-software/webx/domain/criterias"
	"github.com/steve-care-software/webx/domain/cryptography/encryptions/keys"
	"github.com/steve-care-software/webx/domain/cryptography/signatures"
	"github.com/steve-care-software/webx/domain/grammars"
	"github.com/steve-care-software/webx/domain/grammars/cardinalities"
	"github.com/steve-care-software/webx/domain/grammars/values"
	grammar_values "github.com/steve-care-software/webx/domain/grammars/values"
	"github.com/steve-care-software/webx/domain/identities"
	"github.com/steve-care-software/webx/domain/identities/modifications"
	"github.com/steve-care-software/webx/domain/programs/modules"
	"github.com/steve-care-software/webx/infrastructure/disks"
)

// NewApplication creates a new create application
func NewApplication(
	bitrate int,
	basePath string,
	delimiter string,
	extension string,
) creates.Application {
	identityRepository := disks.NewIdentityRepository(basePath, delimiter, extension)
	identityService := disks.NewIdentityService(identityRepository, basePath, delimiter, extension)
	identityApplication := identity_applications.NewApplication(
		identityRepository,
		identityService,
	)

	grammarApp := createGrammar(
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

	commandApp := createCommand(
		commands.NewBuilder(),
		commands.NewAttachmentBuilder(),
		commands.NewVariableAssignmentBuilder(),
		commands.NewParameterDeclarationBuilder(),
		commands.NewApplicationDeclarationBuilder(),
		commands.NewValueBuilder(),
		criterias.NewBuilder(),
	)

	moduleApp := createModule(
		identityApplication,
		modules.NewBuilder(),
		modules.NewModuleBuilder(),
		signatures.NewPrivateKeyFactory(),
		keys.NewFactory(bitrate),
		identities.NewBuilder(),
		modifications.NewBuilder(),
		modifications.NewModificationBuilder(),
		criterias.NewBuilder(),
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
	)

	compilerApp := compiler_applications.NewApplication(
		creates.NewApplication(
			grammarApp,
			commandApp,
			moduleApp,
		))

	additionalModules, err := moduleApp.Execute()
	if err != nil {
		panic(err)
	}

	return creates.NewApplication(
		grammarApp,
		commandApp,
		createModuleWithCompiler(
			compilerApp,
			modules.NewBuilder(),
			modules.NewModuleBuilder(),
			compilers.NewBuilder(),
			compilers.NewElementsBuilder(),
			compilers.NewElementBuilder(),
			compilers.NewParametersBuilder(),
			compilers.NewParameterBuilder(),
			compilers.NewValueBuilder(),
			additionalModules,
		),
	)
}
