package defaults

import (
	creates "github.com/steve-care-software/syntax/applications/engines/creates/grammars"
	"github.com/steve-care-software/syntax/domain/syntax/bytes/grammars"
	"github.com/steve-care-software/syntax/domain/syntax/bytes/grammars/cardinalities"
	grammar_values "github.com/steve-care-software/syntax/domain/syntax/bytes/grammars/values"
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

// Execute executes the create grammar application
func (app *grammar) Execute() (grammars.Grammar, error) {
	lowerCaseLetter, err := app.anyCharacterToken("lowerCaseLetter", "abcdefghijklmnopqrstuvwxyz")
	if err != nil {
		return nil, err
	}

	upperCaseLetter, err := app.anyCharacterToken("upperCaseLetter", "ABCDEFGHIJKLMNOPQRSTUVWXYZ")
	if err != nil {
		return nil, err
	}

	anyCaseLetter, err := app.anyCaseLetterToken(lowerCaseLetter, upperCaseLetter)
	if err != nil {
		return nil, err
	}

	name, err := app.nameToken(lowerCaseLetter, anyCaseLetter)
	if err != nil {
		return nil, err
	}

	moduleStr, err := app.allCharacterToken("moduleStr", "module")
	if err != nil {
		return nil, err
	}

	endOfInstruction, err := app.allCharacterToken("endOfIns", ";;")
	if err != nil {
		return nil, err
	}

	dollarSign, err := app.anyCharacterToken("dollarSign", "$")
	if err != nil {
		return nil, err
	}

	variableName, err := app.variableNameToken(dollarSign, name)
	if err != nil {
		return nil, err
	}

	moduleDeclaration, err := app.moduleDeclarationToken(moduleStr, name, endOfInstruction)
	if err != nil {
		return nil, err
	}

	applicationDeclaration, err := app.applicationDeclarationToken(name, variableName, endOfInstruction)
	if err != nil {
		return nil, err
	}

	inputDirection, err := app.allCharacterToken("inputDirection", "->")
	if err != nil {
		return nil, err
	}

	inputParameter, err := app.parameterToken("inputParameter", inputDirection, variableName, endOfInstruction)
	if err != nil {
		return nil, err
	}

	outputDirection, err := app.allCharacterToken("outputDirection", "<-")
	if err != nil {
		return nil, err
	}

	outputParameter, err := app.parameterToken("outputParameter", outputDirection, variableName, endOfInstruction)
	if err != nil {
		return nil, err
	}

	escapeChar, err := app.allCharacterToken("escapeChar", "\\")
	if err != nil {
		return nil, err
	}

	variableAssignee, err := app.assigneeToken(escapeChar, endOfInstruction)
	if err != nil {
		return nil, err
	}

	equalChar, err := app.allCharacterToken("equalChar", "=")
	if err != nil {
		return nil, err
	}

	variableAssignment, err := app.assignmentToken("variableAssignment", variableName, equalChar, variableAssignee, endOfInstruction)
	if err != nil {
		return nil, err
	}

	attachStr, err := app.allCharacterToken("attachStr", "attach")
	if err != nil {
		return nil, err
	}

	semiColon, err := app.allCharacterToken("semiColon", ":")
	if err != nil {
		return nil, err
	}

	attach, err := app.attachToken(attachStr, variableName, semiColon, endOfInstruction)
	if err != nil {
		return nil, err
	}

	executeStr, err := app.allCharacterToken("executeStr", "execute")
	if err != nil {
		return nil, err
	}

	execute, err := app.executeToken(executeStr, variableName, endOfInstruction)
	if err != nil {
		return nil, err
	}

	executeAssignment, err := app.assignmentToken("executeAssignment", variableName, equalChar, execute, endOfInstruction)
	if err != nil {
		return nil, err
	}

	tokens := []grammars.Token{
		moduleDeclaration,
		applicationDeclaration,
		inputParameter,
		outputParameter,
		variableAssignment,
		attach,
		execute,
		executeAssignment,
	}

	cardinality, err := app.cardinalityOneOccurence()
	if err != nil {
		return nil, err
	}

	elementsList := []grammars.Element{}
	for _, oneToken := range tokens {
		element, err := app.elementFromToken(oneToken, cardinality)
		if err != nil {
			return nil, err
		}

		elementsList = append(elementsList, element)
	}

	root, err := app.oneLinePerElement("root", elementsList)
	if err != nil {
		return nil, err
	}

	channels, err := app.channels()
	if err != nil {
		return nil, err
	}

	return app.builder.Create().WithRoot(root).WithChannels(channels).Now()
}

func (app *grammar) channels() (grammars.Channels, error) {
	space, err := app.allCharacterToken("space", " ")
	if err != nil {
		return nil, err
	}

	tab, err := app.allCharacterToken("tab", "\t")
	if err != nil {
		return nil, err
	}

	newLine, err := app.allCharacterToken("newLine", "\n")
	if err != nil {
		return nil, err
	}

	newRetCar, err := app.allCharacterToken("retCar", "\r")
	if err != nil {
		return nil, err
	}

	tokensList := []grammars.Token{
		space,
		tab,
		newLine,
		newRetCar,
	}

	channelsList := []grammars.Channel{}
	for _, oneToken := range tokensList {
		name := oneToken.Name()
		channel, err := app.channelBuilder.Create().WithName(name).WithToken(oneToken).Now()
		if err != nil {
			return nil, err
		}

		channelsList = append(channelsList, channel)
	}

	return app.channelsBuilder.Create().WithList(channelsList).Now()
}

func (app *grammar) executeToken(
	executeStr grammars.Token,
	variableName grammars.Token,
	endOfInstruction grammars.Token,
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

	endOfInsElement, err := app.elementFromToken(endOfInstruction, cardinality)
	if err != nil {
		return nil, err
	}

	return app.oneLineTokenFromElements("variableAssignment", []grammars.Element{
		executeElement,
		variableNameElement,
		endOfInsElement,
	})
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

	return app.oneLineTokenFromElements("variableAssignment", []grammars.Element{
		attachElement,
		variableNameElement,
		semiColonElement,
		variableNameElement,
		variableNameElement,
		endOfInsElement,
	})
}

func (app *grammar) assignmentToken(
	name string,
	variableName grammars.Token,
	equalChar grammars.Token,
	assignee grammars.Token,
	endOfInstruction grammars.Token,
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

	assigneeElement, err := app.elementFromToken(assignee, cardinality)
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
		assigneeElement,
		endOfInsElement,
	})
}

func (app *grammar) assigneeToken(
	escapeChar grammars.Token,
	endOfInstruction grammars.Token,
) (grammars.Token, error) {
	cardinality, err := app.cardinalityOneOccurence()
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

	return app.oneLineTokenFromElements("assignee", []grammars.Element{
		element,
	})
}

func (app *grammar) parameterToken(
	name string,
	direction grammars.Token,
	variableName grammars.Token,
	endOfInstruction grammars.Token,
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
	})
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

	return app.oneLineTokenFromElements("applicationDeclaration", []grammars.Element{
		moduleNameElement,
		variableNameElement,
		endOfInsElement,
	})
}

func (app *grammar) variableNameToken(
	dollarSign grammars.Token,
	moduleName grammars.Token,
) (grammars.Token, error) {
	cardinality, err := app.cardinalityOneOccurence()
	if err != nil {
		return nil, err
	}

	dollarSignElement, err := app.elementFromToken(dollarSign, cardinality)
	if err != nil {
		return nil, err
	}

	nameElement, err := app.elementFromToken(moduleName, cardinality)
	if err != nil {
		return nil, err
	}

	return app.oneLineTokenFromElements("variableName", []grammars.Element{
		dollarSignElement,
		nameElement,
	})
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

	return app.oneLineTokenFromElements("moduleDeclaration", []grammars.Element{
		moduleStrElement,
		moduleNameElement,
		endOfInsElement,
	})
}

func (app *grammar) nameToken(
	lowerCaseLetter grammars.Token,
	anyCaseLetter grammars.Token,
) (grammars.Token, error) {
	firstLetterCardinality, err := app.cardinalityOneOccurence()
	if err != nil {
		return nil, err
	}

	firstLetterElement, err := app.elementFromToken(lowerCaseLetter, firstLetterCardinality)
	if err != nil {
		return nil, err
	}

	lettersCardinality, err := app.cardinalityBuilder.Create().WithMin(1).Now()
	if err != nil {
		return nil, err
	}

	lettersElement, err := app.elementFromToken(anyCaseLetter, lettersCardinality)
	if err != nil {
		return nil, err
	}

	line, err := app.lineBuilder.Create().WithElements([]grammars.Element{
		firstLetterElement,
		lettersElement,
	}).Now()

	if err != nil {
		return nil, err
	}

	block, err := app.blockBuilder.Create().WithLines([]grammars.Line{
		line,
	}).Now()

	if err != nil {
		return nil, err
	}

	return app.tokenBuilder.Create().WithName("name").WithBlock(block).Now()
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

func (app *grammar) allCharacterToken(tokenName string, letters string) (grammars.Token, error) {
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

	return app.oneLineTokenFromElements(tokenName, elementsList)
}

func (app *grammar) anyCharacterToken(tokenName string, letters string) (grammars.Token, error) {
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

	return app.oneLinePerElement(tokenName, elementsList)
}

func (app *grammar) oneLinePerElement(name string, list []grammars.Element) (grammars.Token, error) {
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

	return app.tokenBuilder.Create().WithName(name).WithBlock(block).Now()
}

func (app *grammar) oneLineTokenFromElements(name string, list []grammars.Element) (grammars.Token, error) {
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

	return app.tokenBuilder.Create().WithName(name).WithBlock(block).Now()
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
