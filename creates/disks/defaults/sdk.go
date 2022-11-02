package defaults

import (
	"github.com/steve-care-software/webx/applications"
	compiler_applications "github.com/steve-care-software/webx/applications/compilers"
	"github.com/steve-care-software/webx/applications/creates"
	grammar_applications "github.com/steve-care-software/webx/applications/grammars"
	identity_applications "github.com/steve-care-software/webx/applications/identities"
	"github.com/steve-care-software/webx/domain/compilers"
	"github.com/steve-care-software/webx/domain/criterias"
	"github.com/steve-care-software/webx/domain/cryptography/encryptions/keys"
	"github.com/steve-care-software/webx/domain/cryptography/hash"
	"github.com/steve-care-software/webx/domain/cryptography/signatures"
	"github.com/steve-care-software/webx/domain/grammars"
	"github.com/steve-care-software/webx/domain/grammars/cardinalities"
	"github.com/steve-care-software/webx/domain/grammars/values"
	grammar_values "github.com/steve-care-software/webx/domain/grammars/values"
	"github.com/steve-care-software/webx/domain/identities"
	"github.com/steve-care-software/webx/domain/identities/modifications"
	"github.com/steve-care-software/webx/domain/instructions"
	instruction_applications "github.com/steve-care-software/webx/domain/instructions/applications"
	"github.com/steve-care-software/webx/domain/instructions/attachments"
	"github.com/steve-care-software/webx/domain/instructions/parameters"
	"github.com/steve-care-software/webx/domain/programs/modules"
	"github.com/steve-care-software/webx/domain/selectors"
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

	moduleApp := createModule(
		identityApplication,
		grammar_applications.NewApplication(),
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
		hash.NewAdapter(),
	)

	selectorApp := createSelector(
		selectors.NewBuilder(),
		selectors.NewSelectorFnBuilder(),
		selectors.NewTokenBuilder(),
		selectors.NewElementBuilder(),
		selectors.NewInsideBuilder(),
		selectors.NewFetchersBuilder(),
		selectors.NewFetcherBuilder(),
		selectors.NewContentFnBuilder(),
		instructions.NewBuilder(),
		instructions.NewInstructionBuilder(),
		instruction_applications.NewBuilder(),
		parameters.NewBuilder(),
		attachments.NewBuilder(),
		attachments.NewVariableBuilder(),
		instructions.NewAssignmentBuilder(),
		instructions.NewValueBuilder(),
	)

	compilerApp := compiler_applications.NewApplication(
		creates.NewApplication(
			grammarApp,
			selectorApp,
			moduleApp,
		))

	initialAdditionalModules, err := moduleApp.Execute()
	if err != nil {
		panic(err)
	}

	moduleWithCompilerApp := createModuleWithCompiler(
		compilerApp,
		modules.NewBuilder(),
		modules.NewModuleBuilder(),
		compilers.NewBuilder(),
		compilers.NewElementsBuilder(),
		compilers.NewElementBuilder(),
		compilers.NewExecutionBuilder(),
		compilers.NewParametersBuilder(),
		compilers.NewParameterBuilder(),
		compilers.NewValueBuilder(),
		initialAdditionalModules,
	)

	createApp := creates.NewApplication(
		grammarApp,
		selectorApp,
		moduleWithCompilerApp,
	)

	additionalModulesWithCompiler, err := moduleWithCompilerApp.Execute()
	if err != nil {
		panic(err)
	}

	engineApp := applications.NewApplication(createApp)
	moduleWithInterpreterApp := createModuleWithInterpreter(
		engineApp,
		modules.NewBuilder(),
		modules.NewModuleBuilder(),
		additionalModulesWithCompiler,
	)

	return creates.NewApplication(
		grammarApp,
		selectorApp,
		moduleWithInterpreterApp,
	)
}
