package defaults

import (
	creates "github.com/steve-care-software/webx/applications/creates/grammars"
	"github.com/steve-care-software/webx/domain/grammars"
	"github.com/steve-care-software/webx/domain/grammars/cardinalities"
	grammar_values "github.com/steve-care-software/webx/domain/grammars/values"
)

type grammar struct {
	builder            grammars.Builder
	channelsBuilder    grammars.ChannelsBuilder
	channelBuilder     grammars.ChannelBuilder
	instanceBuilder    grammars.InstanceBuilder
	everythingBuilder  grammars.EverythingBuilder
	tokensBuilder      grammars.TokensBuilder
	tokenBuilder       grammars.TokenBuilder
	suitesBuilder      grammars.SuitesBuilder
	suiteBuilder       grammars.SuiteBuilder
	blockBuilder       grammars.BlockBuilder
	lineBuilder        grammars.LineBuilder
	elementBuilder     grammars.ElementBuilder
	valueBuilder       grammar_values.Builder
	cardinalityBuilder cardinalities.Builder
}

func createGrammar(
	builder grammars.Builder,
	channelsBuilder grammars.ChannelsBuilder,
	channelBuilder grammars.ChannelBuilder,
	instanceBuilder grammars.InstanceBuilder,
	everythingBuilder grammars.EverythingBuilder,
	tokensBuilder grammars.TokensBuilder,
	tokenBuilder grammars.TokenBuilder,
	suitesBuilder grammars.SuitesBuilder,
	suiteBuilder grammars.SuiteBuilder,
	blockBuilder grammars.BlockBuilder,
	lineBuilder grammars.LineBuilder,
	elementBuilder grammars.ElementBuilder,
	valueBuilder grammar_values.Builder,
	cardinalityBuilder cardinalities.Builder,
) creates.Application {
	out := grammar{
		builder:            builder,
		channelsBuilder:    channelsBuilder,
		channelBuilder:     channelBuilder,
		instanceBuilder:    instanceBuilder,
		everythingBuilder:  everythingBuilder,
		tokensBuilder:      tokensBuilder,
		tokenBuilder:       tokenBuilder,
		suitesBuilder:      suitesBuilder,
		suiteBuilder:       suiteBuilder,
		blockBuilder:       blockBuilder,
		lineBuilder:        lineBuilder,
		elementBuilder:     elementBuilder,
		valueBuilder:       valueBuilder,
		cardinalityBuilder: cardinalityBuilder,
	}

	return &out
}

// LowerCaseLetter returns the lowerCase letter token
func (app *grammar) LowerCaseLetter() grammars.Token {
	token, err := app.anyCharacterToken("lowerCaseLetter", "abcdefghijklmnopqrstuvwxyz")
	if err != nil {
		panic(err)
	}

	return token
}

// UpperCaseLetter returns the upperCase letter token
func (app *grammar) UpperCaseLetter() grammars.Token {
	token, err := app.anyCharacterToken("upperCaseLetter", "ABCDEFGHIJKLMNOPQRSTUVWXYZ")
	if err != nil {
		panic(err)
	}

	return token
}

// AnyCaseLetter returns the anyCase letter token
func (app *grammar) AnyCaseLetter() grammars.Token {
	token, err := app.anyCaseLetterToken(
		app.LowerCaseLetter(),
		app.UpperCaseLetter(),
	)

	if err != nil {
		panic(err)
	}

	return token
}

// AnyNumber returns the anyNumber token
func (app *grammar) AnyNumber() grammars.Token {
	token, err := app.anyCharacterToken("anyNumber", "0123456789")
	if err != nil {
		panic(err)
	}

	return token
}

// Name returns the name token
func (app *grammar) Name() grammars.Token {
	token, err := app.nameToken(
		app.LowerCaseLetter(),
		app.AnyCaseLetter(),
		app.AnyNumber(),
	)

	if err != nil {
		panic(err)
	}

	return token
}

// ModuleString returns the module string token
func (app *grammar) ModuleString() grammars.Token {
	token, err := app.allCharacterToken("moduleStr", "module")
	if err != nil {
		panic(err)
	}

	return token
}

// EndOfLine returns the end of line token
func (app *grammar) EndOfLine() grammars.Token {
	token, err := app.allCharacterToken("endOfIns", ";;")
	if err != nil {
		panic(err)
	}

	return token
}

// CommercialAChar returns the commercial a char (@) token
func (app *grammar) CommercialAChar() grammars.Token {
	commercialASuites, err := app.suites([][]byte{
		[]byte("@"),
	}, nil)

	if err != nil {
		panic(err)
	}

	token, err := app.singleCharacterToken("singleCommercialA", "commercialA", []byte("@")[0], commercialASuites)
	if err != nil {
		panic(err)
	}

	return token
}

// ModuleName returns the moduleName token
func (app *grammar) ModuleName() grammars.Token {
	token, err := app.moduleNameToken(
		app.CommercialAChar(),
		app.Name(),
	)

	if err != nil {
		panic(err)
	}

	return token
}

// ModuleDeclaration returns the module declaration token
func (app *grammar) ModuleDeclaration() grammars.Token {
	token, err := app.moduleDeclarationToken(
		app.ModuleString(),
		app.ModuleName(),
		app.EndOfLine(),
	)

	if err != nil {
		panic(err)
	}

	return token
}

// DollarSignChar returns the dollar sign char ($) token
func (app *grammar) DollarSignChar() grammars.Token {
	dollarSignSuites, err := app.suites([][]byte{
		[]byte("$"),
	}, nil)

	if err != nil {
		panic(err)
	}

	token, err := app.singleCharacterToken("singleDollarSign", "dollarSign", []byte("$")[0], dollarSignSuites)
	if err != nil {
		panic(err)
	}

	return token
}

// VariableName returns the variable name token
func (app *grammar) VariableName() grammars.Token {
	token, err := app.variableNameToken(
		app.DollarSignChar(),
		app.Name(),
	)

	if err != nil {
		panic(err)
	}

	return token
}

// ApplicationDeclaration returns the application declaration token
func (app *grammar) ApplicationDeclaration() grammars.Token {
	token, err := app.applicationDeclarationToken(
		app.ModuleName(),
		app.VariableName(),
		app.EndOfLine(),
	)

	if err != nil {
		panic(err)
	}

	return token
}

// InputDirection returns the input direction token
func (app *grammar) InputDirection() grammars.Token {
	token, err := app.allCharacterToken("inputDirection", "->")
	if err != nil {
		panic(err)
	}

	return token
}

// OutputDirection returns the output direction token
func (app *grammar) OutputDirection() grammars.Token {
	token, err := app.allCharacterToken("outputDirection", "<-")
	if err != nil {
		panic(err)
	}

	return token
}

// InputParameter returns the input parameter token
func (app *grammar) InputParameter() grammars.Token {
	inputParameterSuites, err := app.suites([][]byte{
		[]byte("-> $myInput;;"),
	}, nil)

	if err != nil {
		panic(err)
	}

	token, err := app.parameterToken(
		"inputParameter",
		app.InputDirection(),
		app.VariableName(),
		app.EndOfLine(),
		inputParameterSuites,
	)

	if err != nil {
		panic(err)
	}

	return token
}

// OutputParameter returns the output parameter token
func (app *grammar) OutputParameter() grammars.Token {
	outputParameterSuites, err := app.suites([][]byte{
		[]byte("<- $myInput;;"),
	}, nil)

	if err != nil {
		panic(err)
	}

	token, err := app.parameterToken(
		"outputParameter",
		app.OutputDirection(),
		app.VariableName(),
		app.EndOfLine(),
		outputParameterSuites,
	)

	if err != nil {
		panic(err)
	}

	return token
}

// EqualChar returns the equal char (=) token
func (app *grammar) EqualChar() grammars.Token {
	equalSuites, err := app.suites([][]byte{
		[]byte("="),
	}, nil)

	if err != nil {
		panic(err)
	}

	token, err := app.singleCharacterToken("singleEqualChar", "equalChar", []byte("=")[0], equalSuites)
	if err != nil {
		panic(err)
	}

	return token
}

// AttachString returns the attach string token
func (app *grammar) AttachString() grammars.Token {
	token, err := app.allCharacterToken("attachStr", "attach")
	if err != nil {
		panic(err)
	}

	return token
}

// SemiColonChar returns the semiColon char (:) token
func (app *grammar) SemiColonChar() grammars.Token {
	semiColonSuites, err := app.suites([][]byte{
		[]byte(":"),
	}, nil)

	if err != nil {
		panic(err)
	}

	token, err := app.singleCharacterToken("singleSemiColon", "semiColon", []byte(":")[0], semiColonSuites)
	if err != nil {
		panic(err)
	}

	return token
}

// AttachInstruction returns the attach instruction token
func (app *grammar) Attach() grammars.Token {
	token, err := app.attachToken(
		app.AttachString(),
		app.VariableName(),
		app.SemiColonChar(),
		app.EndOfLine(),
	)

	if err != nil {
		panic(err)
	}

	return token
}

// Assignment returns the assignment instruction token
func (app *grammar) Assignment() grammars.Token {
	token, err := app.valueToken(
		app.VariableName(),
		app.EqualChar(),
		app.EndOfLine(),
	)

	if err != nil {
		panic(err)
	}

	return token
}

// Instruction returns the instruction token
func (app *grammar) Instruction() grammars.Token {
	tokens := []grammars.Token{
		app.ModuleDeclaration(),
		app.ApplicationDeclaration(),
		app.InputParameter(),
		app.OutputParameter(),
		app.Attach(),
		app.Assignment(),
	}

	cardinality, err := app.cardinalityOneOccurence()
	if err != nil {
		panic(err)
	}

	elementsList := []grammars.Element{}
	for _, oneToken := range tokens {
		element, err := app.elementFromToken(oneToken, cardinality)
		if err != nil {
			panic(err)
		}

		elementsList = append(elementsList, element)
	}

	instructionSuites, err := app.suites([][]byte{
		[]byte("module @myModule;;"),
		[]byte("@myModule $myApplication;;"),
		[]byte("-> $myInput;;"),
		[]byte("<- $myOutput;;"),
		[]byte("$myVariable = ANY VALUE EXCEPT \\;; NON-ESCAPED SEMI-COLON;;"),
		[]byte("attach $myDataVariable:$data $myAppVariable;;"),
		[]byte("$myOutput = execute $myAppVariable;;"),
		[]byte("$myValue = $myOtherVariable;;"),
	}, nil)

	if err != nil {
		panic(err)
	}

	token, err := app.oneLinePerElement("instruction", elementsList, instructionSuites)
	if err != nil {
		panic(err)
	}

	return token
}

// Root returns the root token
func (app *grammar) Root() grammars.Token {
	instruction := app.Instruction()
	multipleCardinality, err := app.cardinalityBuilder.Create().WithMin(1).Now()
	if err != nil {
		panic(err)
	}

	instance, err := app.instanceBuilder.Create().WithToken(instruction).Now()
	if err != nil {
		panic(err)
	}

	rootElement, err := app.elementBuilder.Create().WithCardinality(multipleCardinality).WithInstance(instance).Now()
	if err != nil {
		panic(err)
	}

	line, err := app.lineBuilder.Create().WithElements([]grammars.Element{
		rootElement,
	}).Now()

	if err != nil {
		panic(err)
	}

	block, err := app.blockBuilder.Create().WithLines([]grammars.Line{
		line,
	}).Now()
	if err != nil {
		panic(err)
	}

	rootSuites, err := app.suites([][]byte{
		[]byte(`
			-> $script;;
			<- $output;;

			$createGrammarValueCode = {
				-> $name;;
				-> $numberStr;;
				<- $value;;

				module @castToInt;;
				@castToInt $castToIntApp;;
				attach $numberStr:$value $castToIntApp;;
				$number = execute $castToIntApp;;

				module @newGrammarValue;;
				@newGrammarValue $valueApp;;
				attach $number:$number $valueApp;;
				attach $name:$name $valueApp;;
				$value = execute $valueApp;;
			};;


			module @containerMapWithStringKeynames;;
			@containerMapWithStringKeynames $paramsApp;;
			$nameStr = dollarSign;;
			$valueStr = 36;
			attach $nameStr:$name $paramsApp;;
			attach $valueStr:$number $valueApp;;
			$params = execute $paramsApp;;

			module @parseThenInterpret;;
			@parseThenInterpret $interpreterApp;;
			attach $params:$params $interpreterApp;;
			attach $createGrammarValueCode:$script $interpreterApp;;
			$output = execute $interpreterApp;;
		`),
	}, nil)

	if err != nil {
		panic(err)
	}

	tokenBuilder := app.tokenBuilder.Create().WithName("root").WithBlock(block)
	if rootSuites != nil {
		tokenBuilder.WithSuites(rootSuites)
	}

	token, err := tokenBuilder.Now()
	if err != nil {
		panic(err)
	}

	return token
}

// Execute executes the create grammar application
func (app *grammar) Execute() (grammars.Grammar, error) {
	rootToken := app.Root()
	channels := app.Channels()
	return app.builder.Create().
		WithRoot(rootToken).
		WithChannels(channels).
		Now()
}

func (app *grammar) suite(isValid bool, data []byte) (grammars.Suite, error) {
	builder := app.suiteBuilder.Create()
	if isValid {
		return builder.WithValid(data).Now()
	}

	return builder.WithInvalid(data).Now()
}

func (app *grammar) suites(valid [][]byte, invalid [][]byte) (grammars.Suites, error) {
	suitesList := []grammars.Suite{}
	if valid != nil {
		for _, oneValid := range valid {
			suite, err := app.suite(true, oneValid)
			if err != nil {
				return nil, err
			}

			suitesList = append(suitesList, suite)
		}
	}

	if invalid != nil {
		for _, oneInvalid := range invalid {
			suite, err := app.suite(false, oneInvalid)
			if err != nil {
				return nil, err
			}

			suitesList = append(suitesList, suite)
		}
	}

	return app.suitesBuilder.Create().WithList(suitesList).Now()
}

func (app *grammar) valueToken(
	variableName grammars.Token,
	equalChar grammars.Token,
	endOfInstruction grammars.Token,
) (grammars.Token, error) {
	cardinality, err := app.cardinalityOneOccurence()
	if err != nil {
		return nil, err
	}

	constantToken, err := app.valueConstantToken(variableName, equalChar, endOfInstruction)
	if err != nil {
		return nil, err
	}

	constantElement, err := app.elementFromToken(constantToken, cardinality)
	if err != nil {
		return nil, err
	}

	variableToken, err := app.valueVariableToken(variableName, equalChar, endOfInstruction)
	if err != nil {
		return nil, err
	}

	variableElement, err := app.elementFromToken(variableToken, cardinality)
	if err != nil {
		return nil, err
	}

	instructionsToken, err := app.valueInstructionsToken(variableName, equalChar, endOfInstruction)
	if err != nil {
		return nil, err
	}

	instructionElement, err := app.elementFromToken(instructionsToken, cardinality)
	if err != nil {
		return nil, err
	}

	executionToken, err := app.valueExecutionToken(variableName, equalChar, endOfInstruction)
	if err != nil {
		return nil, err
	}

	executionElement, err := app.elementFromToken(executionToken, cardinality)
	if err != nil {
		return nil, err
	}

	valueSuites, err := app.suites([][]byte{
		[]byte("$myVariable = ANY VALUE EXCEPT \\;; NON-ESCAPED SEMI-COLON;;"),
		[]byte("$myOutput = execute $myAppVariable;;"),
		[]byte("$myValue = $myOtherVariable;;"),
	}, nil)

	if err != nil {
		return nil, err
	}

	return app.oneLinePerElement("assignment", []grammars.Element{
		instructionElement,
		executionElement,
		variableElement,
		constantElement,
	}, valueSuites)
}

func (app *grammar) valueConstantToken(
	variableName grammars.Token,
	equalChar grammars.Token,
	endOfInstruction grammars.Token,
) (grammars.Token, error) {
	escapeChar, err := app.allCharacterToken("escapeChar", "\\")
	if err != nil {
		return nil, err
	}

	assignmentValue, err := app.assignmentValueToken(escapeChar, endOfInstruction)
	if err != nil {
		return nil, err
	}

	variableAssignmentSuites, err := app.suites([][]byte{
		[]byte("$myVariable = ANY VALUE EXCEPT NON-ESCAPED SEMI-COLON;;"),
	}, nil)

	if err != nil {
		return nil, err
	}

	return app.assignmentToken("valueConstant", variableName, equalChar, assignmentValue, endOfInstruction, variableAssignmentSuites)
}

func (app *grammar) valueVariableToken(
	variableName grammars.Token,
	equalChar grammars.Token,
	endOfInstruction grammars.Token,
) (grammars.Token, error) {
	variableAssignmentSuites, err := app.suites([][]byte{
		[]byte("$myVariable = $other;;"),
	}, nil)

	if err != nil {
		return nil, err
	}

	return app.assignmentToken("valueVariable", variableName, equalChar, variableName, endOfInstruction, variableAssignmentSuites)

}

func (app *grammar) valueInstructionsToken(
	variableName grammars.Token,
	equalChar grammars.Token,
	endOfInstruction grammars.Token,
) (grammars.Token, error) {
	assignmentCode, err := app.assignmentCodeToken("root")
	if err != nil {
		return nil, err
	}

	return app.assignmentToken("codeAssignment", variableName, equalChar, assignmentCode, endOfInstruction, nil)
}

func (app *grammar) valueExecutionToken(
	variableName grammars.Token,
	equalChar grammars.Token,
	endOfInstruction grammars.Token,
) (grammars.Token, error) {
	executeStr, err := app.allCharacterToken("executeStr", "execute")
	if err != nil {
		return nil, err
	}

	execute, err := app.executeToken(executeStr, variableName)
	if err != nil {
		return nil, err
	}

	executeAssignmentSuites, err := app.suites([][]byte{
		[]byte("$myOutput = execute $myAppVariable;;"),
	}, nil)

	if err != nil {
		return nil, err
	}

	return app.assignmentToken("valueExecution", variableName, equalChar, execute, endOfInstruction, executeAssignmentSuites)
}

func (app *grammar) singleLineComment(
	endOfCommentIns grammars.Token,
) (grammars.Token, error) {
	cardinality, err := app.cardinalityOneOccurence()
	if err != nil {
		return nil, err
	}

	slashSlashStr, err := app.allCharacterToken("slashSlash", "//")
	if err != nil {
		return nil, err
	}

	slashSlashIns, err := app.instanceBuilder.Create().WithToken(slashSlashStr).Now()
	if err != nil {
		return nil, err
	}

	slashSlashElement, err := app.elementBuilder.Create().WithCardinality(cardinality).WithInstance(slashSlashIns).Now()
	if err != nil {
		return nil, err
	}

	everything, err := app.everythingBuilder.Create().WithName("anythingExceptEndOfCommentInstruction").WithException(endOfCommentIns).Now()
	if err != nil {
		return nil, err
	}

	instance, err := app.instanceBuilder.Create().WithEverything(everything).Now()
	if err != nil {
		return nil, err
	}

	element, err := app.elementBuilder.Create().WithCardinality(cardinality).WithInstance(instance).Now()
	if err != nil {
		return nil, err
	}

	suites, err := app.suites([][]byte{
		[]byte("// this is a comment"),
	}, nil)

	if err != nil {
		return nil, err
	}

	return app.oneLineTokenFromElements("singleLineComment", []grammars.Element{
		slashSlashElement,
		element,
	}, suites)
}

func (app *grammar) assignmentCodeToken(
	recursiveTokenName string,
) (grammars.Token, error) {
	cardinality, err := app.cardinalityOneOccurence()
	if err != nil {
		return nil, err
	}

	openBracketSuites, err := app.suites([][]byte{
		[]byte("{"),
	}, nil)

	if err != nil {
		return nil, err
	}

	openBracket, err := app.singleCharacterToken("singleOpenBracket", "openBracket", []byte("{")[0], openBracketSuites)
	if err != nil {
		return nil, err
	}

	openBracketElement, err := app.elementFromToken(openBracket, cardinality)
	if err != nil {
		return nil, err
	}

	closeBracketSuites, err := app.suites([][]byte{
		[]byte("}"),
	}, nil)

	if err != nil {
		return nil, err
	}

	closeBracket, err := app.singleCharacterToken("singleCloseBracket", "closeBracket", []byte("}")[0], closeBracketSuites)
	if err != nil {
		return nil, err
	}

	closeBracketElement, err := app.elementFromToken(closeBracket, cardinality)
	if err != nil {
		return nil, err
	}

	recursiveTokenElement, err := app.elementBuilder.Create().WithCardinality(cardinality).WithRecursive(recursiveTokenName).Now()
	if err != nil {
		return nil, err
	}

	return app.oneLineTokenFromElements("assignmentCode", []grammars.Element{
		openBracketElement,
		recursiveTokenElement,
		closeBracketElement,
	}, nil)
}

func (app *grammar) executeToken(
	executeStr grammars.Token,
	variableName grammars.Token,
) (grammars.Token, error) {
	cardinality, err := app.cardinalityOneOccurence()
	if err != nil {
		return nil, err
	}

	executeElement, err := app.elementFromToken(executeStr, cardinality)
	if err != nil {
		return nil, err
	}

	variableNameElement, err := app.elementFromToken(variableName, cardinality)
	if err != nil {
		return nil, err
	}

	suites, err := app.suites([][]byte{
		[]byte("execute $myVariable"),
	}, nil)

	if err != nil {
		return nil, err
	}

	return app.oneLineTokenFromElements("execute", []grammars.Element{
		executeElement,
		variableNameElement,
	}, suites)
}

func (app *grammar) attachToken(
	attachStr grammars.Token,
	variableName grammars.Token,
	semiColon grammars.Token,
	endOfInstruction grammars.Token,
) (grammars.Token, error) {
	cardinality, err := app.cardinalityOneOccurence()
	if err != nil {
		return nil, err
	}

	attachElement, err := app.elementFromToken(attachStr, cardinality)
	if err != nil {
		return nil, err
	}

	variableNameElement, err := app.elementFromToken(variableName, cardinality)
	if err != nil {
		return nil, err
	}

	semiColonElement, err := app.elementFromToken(semiColon, cardinality)
	if err != nil {
		return nil, err
	}

	endOfInsElement, err := app.elementFromToken(endOfInstruction, cardinality)
	if err != nil {
		return nil, err
	}

	suites, err := app.suites([][]byte{
		[]byte("attach $current:$inside $myApplication;;"),
	}, nil)

	if err != nil {
		return nil, err
	}

	return app.oneLineTokenFromElements("attachment", []grammars.Element{
		attachElement,
		variableNameElement,
		semiColonElement,
		variableNameElement,
		variableNameElement,
		endOfInsElement,
	}, suites)
}

func (app *grammar) assignmentToken(
	name string,
	variableName grammars.Token,
	equalChar grammars.Token,
	assignmentValue grammars.Token,
	endOfInstruction grammars.Token,
	suites grammars.Suites,
) (grammars.Token, error) {
	cardinality, err := app.cardinalityOneOccurence()
	if err != nil {
		return nil, err
	}

	variableNameElement, err := app.elementFromToken(variableName, cardinality)
	if err != nil {
		return nil, err
	}

	equalElement, err := app.elementFromToken(equalChar, cardinality)
	if err != nil {
		return nil, err
	}

	assignmentValueElement, err := app.elementFromToken(assignmentValue, cardinality)
	if err != nil {
		return nil, err
	}

	endOfInsElement, err := app.elementFromToken(endOfInstruction, cardinality)
	if err != nil {
		return nil, err
	}

	return app.oneLineTokenFromElements(name, []grammars.Element{
		variableNameElement,
		equalElement,
		assignmentValueElement,
		endOfInsElement,
	}, suites)
}

func (app *grammar) assignmentValueToken(
	escapeChar grammars.Token,
	endOfInstruction grammars.Token,
) (grammars.Token, error) {
	cardinality, err := app.cardinalityBuilder.Create().
		WithMin(1).
		Now()

	if err != nil {
		return nil, err
	}

	everything, err := app.everythingBuilder.Create().WithName("everythingExceptEndOfInstruction").WithException(endOfInstruction).WithEscape(escapeChar).Now()
	if err != nil {
		return nil, err
	}

	instance, err := app.instanceBuilder.Create().WithEverything(everything).Now()
	if err != nil {
		return nil, err
	}

	element, err := app.elementBuilder.Create().WithCardinality(cardinality).WithInstance(instance).Now()
	if err != nil {
		return nil, err
	}

	suites, err := app.suites([][]byte{
		[]byte("ANY VALUE EXCEPT NON-ESCAPED SEMI-COLON"),
	}, nil)

	if err != nil {
		return nil, err
	}

	return app.oneLineTokenFromElements("assignmentValue", []grammars.Element{
		element,
	}, suites)
}

func (app *grammar) parameterToken(
	name string,
	direction grammars.Token,
	variableName grammars.Token,
	endOfInstruction grammars.Token,
	suites grammars.Suites,
) (grammars.Token, error) {
	cardinality, err := app.cardinalityOneOccurence()
	if err != nil {
		return nil, err
	}

	directionElement, err := app.elementFromToken(direction, cardinality)
	if err != nil {
		return nil, err
	}

	variableNameElement, err := app.elementFromToken(variableName, cardinality)
	if err != nil {
		return nil, err
	}

	endOfInsElement, err := app.elementFromToken(endOfInstruction, cardinality)
	if err != nil {
		return nil, err
	}

	return app.oneLineTokenFromElements(name, []grammars.Element{
		directionElement,
		variableNameElement,
		endOfInsElement,
	}, suites)
}

func (app *grammar) applicationDeclarationToken(
	moduleName grammars.Token,
	variableName grammars.Token,
	endOfInstruction grammars.Token,
) (grammars.Token, error) {
	cardinality, err := app.cardinalityOneOccurence()
	if err != nil {
		return nil, err
	}

	moduleNameElement, err := app.elementFromToken(moduleName, cardinality)
	if err != nil {
		return nil, err
	}

	variableNameElement, err := app.elementFromToken(variableName, cardinality)
	if err != nil {
		return nil, err
	}

	endOfInsElement, err := app.elementFromToken(endOfInstruction, cardinality)
	if err != nil {
		return nil, err
	}

	suites, err := app.suites([][]byte{
		[]byte("@myModule $myApplication;;"),
	}, nil)

	if err != nil {
		return nil, err
	}

	return app.oneLineTokenFromElements("applicationDeclaration", []grammars.Element{
		moduleNameElement,
		variableNameElement,
		endOfInsElement,
	}, suites)
}

func (app *grammar) moduleNameToken(
	commercialA grammars.Token,
	name grammars.Token,
) (grammars.Token, error) {
	cardinality, err := app.cardinalityOneOccurence()
	if err != nil {
		return nil, err
	}

	commercialAElement, err := app.elementFromToken(commercialA, cardinality)
	if err != nil {
		return nil, err
	}

	nameElement, err := app.elementFromToken(name, cardinality)
	if err != nil {
		return nil, err
	}

	suites, err := app.suites([][]byte{
		[]byte("@myModule"),
	}, nil)

	if err != nil {
		return nil, err
	}

	return app.oneLineTokenFromElements("moduleName", []grammars.Element{
		commercialAElement,
		nameElement,
	}, suites)
}

func (app *grammar) variableNameToken(
	dollarSign grammars.Token,
	name grammars.Token,
) (grammars.Token, error) {
	cardinality, err := app.cardinalityOneOccurence()
	if err != nil {
		return nil, err
	}

	dollarSignElement, err := app.elementFromToken(dollarSign, cardinality)
	if err != nil {
		return nil, err
	}

	nameElement, err := app.elementFromToken(name, cardinality)
	if err != nil {
		return nil, err
	}

	suites, err := app.suites([][]byte{
		[]byte("$myName"),
	}, nil)

	if err != nil {
		return nil, err
	}

	return app.oneLineTokenFromElements("variableName", []grammars.Element{
		dollarSignElement,
		nameElement,
	}, suites)
}

func (app *grammar) moduleDeclarationToken(
	moduleStr grammars.Token,
	moduleName grammars.Token,
	endOfInstruction grammars.Token,
) (grammars.Token, error) {
	cardinality, err := app.cardinalityOneOccurence()
	if err != nil {
		return nil, err
	}

	moduleStrElement, err := app.elementFromToken(moduleStr, cardinality)
	if err != nil {
		return nil, err
	}

	moduleNameElement, err := app.elementFromToken(moduleName, cardinality)
	if err != nil {
		return nil, err
	}

	endOfInsElement, err := app.elementFromToken(endOfInstruction, cardinality)
	if err != nil {
		return nil, err
	}

	suites, err := app.suites([][]byte{
		[]byte("module @myModule;;"),
	}, nil)

	if err != nil {
		return nil, err
	}

	return app.oneLineTokenFromElements("moduleDeclaration", []grammars.Element{
		moduleStrElement,
		moduleNameElement,
		endOfInsElement,
	}, suites)
}

func (app *grammar) nameToken(
	lowerCaseLetter grammars.Token,
	anyCaseLetter grammars.Token,
	anyNumber grammars.Token,
) (grammars.Token, error) {
	oneCardinality, err := app.cardinalityOneOccurence()
	if err != nil {
		return nil, err
	}

	firstLetterElement, err := app.elementFromToken(lowerCaseLetter, oneCardinality)
	if err != nil {
		return nil, err
	}

	onePlusCardinality, err := app.cardinalityBuilder.Create().WithMin(1).Now()
	if err != nil {
		return nil, err
	}

	lettersElement, err := app.elementFromToken(anyCaseLetter, onePlusCardinality)
	if err != nil {
		return nil, err
	}

	numberElement, err := app.elementFromToken(anyNumber, onePlusCardinality)
	if err != nil {
		return nil, err
	}

	firstLine, err := app.lineBuilder.Create().WithElements([]grammars.Element{
		firstLetterElement,
		lettersElement,
	}).Now()

	if err != nil {
		return nil, err
	}

	secondLine, err := app.lineBuilder.Create().WithElements([]grammars.Element{
		numberElement,
	}).Now()

	if err != nil {
		return nil, err
	}

	block, err := app.blockBuilder.Create().WithLines([]grammars.Line{
		firstLine,
		secondLine,
	}).Now()

	if err != nil {
		return nil, err
	}

	suites, err := app.suites([][]byte{
		[]byte("myName"),
		[]byte("0"),
	}, nil)

	if err != nil {
		return nil, err
	}

	return app.tokenBuilder.Create().WithName("name").WithBlock(block).WithSuites(suites).Now()
}

func (app *grammar) anyCaseLetterToken(
	lowerCaseLetter grammars.Token,
	upperCaseLetter grammars.Token,
) (grammars.Token, error) {
	cardinality, err := app.cardinalityOneOccurence()
	if err != nil {
		return nil, err
	}

	lowerCaseLine, err := app.lineFromToken(lowerCaseLetter, cardinality)
	if err != nil {
		return nil, err
	}

	upperCaseLine, err := app.lineFromToken(upperCaseLetter, cardinality)
	if err != nil {
		return nil, err
	}

	block, err := app.blockBuilder.Create().WithLines([]grammars.Line{
		lowerCaseLine,
		upperCaseLine,
	}).Now()

	if err != nil {
		return nil, err
	}

	return app.tokenBuilder.Create().WithName("anyCaseLetter").WithBlock(block).Now()
}

func (app *grammar) singleCharacterToken(tokenName string, valueName string, valueNumber byte, suites grammars.Suites) (grammars.Token, error) {
	cardinality, err := app.cardinalityOneOccurence()
	if err != nil {
		return nil, err
	}

	value, err := app.valueBuilder.Create().WithName(valueName).WithNumber(valueNumber).Now()
	if err != nil {
		return nil, err
	}

	element, err := app.elementBuilder.Create().WithCardinality(cardinality).WithValue(value).Now()
	if err != nil {
		return nil, err
	}

	return app.oneLineTokenFromElements(tokenName, []grammars.Element{
		element,
	}, suites)
}

func (app *grammar) allCharacterToken(tokenName string, letters string) (grammars.Token, error) {
	suites, err := app.suites([][]byte{
		[]byte(letters),
	}, nil)
	if err != nil {
		return nil, err
	}

	cardinality, err := app.cardinalityOneOccurence()
	if err != nil {
		return nil, err
	}

	elementsList := []grammars.Element{}
	for _, oneLetter := range letters {
		name := string(oneLetter)
		number := byte(oneLetter)
		value, err := app.valueBuilder.Create().WithName(name).WithNumber(number).Now()
		if err != nil {
			return nil, err
		}

		element, err := app.elementBuilder.Create().WithCardinality(cardinality).WithValue(value).Now()
		if err != nil {
			return nil, err
		}

		elementsList = append(elementsList, element)
	}

	return app.oneLineTokenFromElements(tokenName, elementsList, suites)
}

func (app *grammar) anyCharacterToken(tokenName string, letters string) (grammars.Token, error) {
	valid := [][]byte{}
	for _, oneLetter := range letters {
		valid = append(valid, []byte(string(oneLetter)))
	}

	suites, err := app.suites(valid, nil)
	if err != nil {
		return nil, err
	}

	cardinality, err := app.cardinalityOneOccurence()
	if err != nil {
		return nil, err
	}

	elementsList := []grammars.Element{}
	for _, oneLetter := range letters {
		name := string(oneLetter)
		number := byte(oneLetter)
		value, err := app.valueBuilder.Create().WithName(name).WithNumber(number).Now()
		if err != nil {
			return nil, err
		}

		element, err := app.elementBuilder.Create().WithCardinality(cardinality).WithValue(value).Now()
		if err != nil {
			return nil, err
		}

		elementsList = append(elementsList, element)
	}

	return app.oneLinePerElement(tokenName, elementsList, suites)
}

func (app *grammar) oneLinePerElement(name string, list []grammars.Element, suites grammars.Suites) (grammars.Token, error) {
	lines := []grammars.Line{}
	for _, oneElement := range list {
		line, err := app.lineBuilder.Create().WithElements([]grammars.Element{
			oneElement,
		}).Now()

		if err != nil {
			return nil, err
		}

		lines = append(lines, line)
	}

	block, err := app.blockBuilder.Create().WithLines(lines).Now()
	if err != nil {
		return nil, err
	}

	builder := app.tokenBuilder.Create().WithName(name).WithBlock(block)
	if suites != nil {
		builder.WithSuites(suites)
	}

	return builder.Now()
}

func (app *grammar) oneLineTokenFromElements(name string, list []grammars.Element, suites grammars.Suites) (grammars.Token, error) {
	line, err := app.lineBuilder.Create().WithElements(list).Now()
	if err != nil {
		return nil, err
	}

	block, err := app.blockBuilder.Create().WithLines([]grammars.Line{
		line,
	}).Now()

	if err != nil {
		return nil, err
	}

	builder := app.tokenBuilder.Create().WithName(name).WithBlock(block)
	if suites != nil {
		builder.WithSuites(suites)
	}

	return builder.Now()
}

func (app *grammar) lineFromToken(token grammars.Token, cardinality cardinalities.Cardinality) (grammars.Line, error) {
	element, err := app.elementFromToken(token, cardinality)
	if err != nil {
		return nil, err
	}

	return app.lineBuilder.Create().WithElements([]grammars.Element{
		element,
	}).Now()
}

func (app *grammar) elementFromToken(token grammars.Token, cardinality cardinalities.Cardinality) (grammars.Element, error) {
	instance, err := app.instanceBuilder.Create().WithToken(token).Now()
	if err != nil {
		return nil, err
	}

	return app.elementBuilder.Create().WithCardinality(cardinality).WithInstance(instance).Now()
}

func (app *grammar) cardinalityOneOccurence() (cardinalities.Cardinality, error) {
	return app.cardinalityBuilder.Create().
		WithMin(1).
		WithMax(1).
		Now()
}

// Channels returns the channels
func (app *grammar) Channels() grammars.Channels {
	spaceSuites, err := app.suites([][]byte{
		[]byte(" "),
	}, nil)

	if err != nil {
		panic(err)
	}

	space, err := app.singleCharacterToken("singleSpace", "space", []byte(" ")[0], spaceSuites)
	if err != nil {
		panic(err)
	}

	tabSuites, err := app.suites([][]byte{
		[]byte("\t"),
	}, nil)

	if err != nil {
		panic(err)
	}

	tab, err := app.singleCharacterToken("singleTab", "tab", []byte("\t")[0], tabSuites)
	if err != nil {
		panic(err)
	}

	newLineSuites, err := app.suites([][]byte{
		[]byte("\n"),
	}, nil)

	if err != nil {
		panic(err)
	}

	newLine, err := app.singleCharacterToken("singleNewLine", "newLine", []byte("\n")[0], newLineSuites)
	if err != nil {
		panic(err)
	}

	retCarSuites, err := app.suites([][]byte{
		[]byte("\r"),
	}, nil)

	if err != nil {
		panic(err)
	}

	newRetCar, err := app.singleCharacterToken("singleRetCar", "retCar", []byte("\r")[0], retCarSuites)
	if err != nil {
		panic(err)
	}

	singleLineComment, err := app.singleLineComment(newLine)
	if err != nil {
		panic(err)
	}

	tokensList := []grammars.Token{
		space,
		tab,
		newLine,
		newRetCar,
		singleLineComment,
	}

	channelsList := []grammars.Channel{}
	for _, oneToken := range tokensList {
		name := oneToken.Name()
		channel, err := app.channelBuilder.Create().WithName(name).WithToken(oneToken).Now()
		if err != nil {
			panic(err)
		}

		channelsList = append(channelsList, channel)
	}

	ins, err := app.channelsBuilder.Create().WithList(channelsList).Now()
	if err != nil {
		panic(err)
	}

	return ins
}
