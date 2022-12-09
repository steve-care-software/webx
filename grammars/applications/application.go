package applications

import (
	"bytes"
	"errors"
	"fmt"

	"github.com/steve-care-software/webx/blockchains/applications"
	"github.com/steve-care-software/webx/blockchains/domain/cryptography/hash"
	"github.com/steve-care-software/webx/domain/databases/references"
	contents_grammars "github.com/steve-care-software/webx/grammars/domain/contents/grammars"
	contents_channels "github.com/steve-care-software/webx/grammars/domain/contents/grammars/channels"
	contents_elements "github.com/steve-care-software/webx/grammars/domain/contents/grammars/elements"
	contents_everythings "github.com/steve-care-software/webx/grammars/domain/contents/grammars/everythings"
	contents_matches "github.com/steve-care-software/webx/grammars/domain/contents/grammars/matches"
	contents_suites "github.com/steve-care-software/webx/grammars/domain/contents/grammars/suites"
	contents_tokens "github.com/steve-care-software/webx/grammars/domain/contents/grammars/tokens"
	"github.com/steve-care-software/webx/grammars/domain/grammars"
	"github.com/steve-care-software/webx/grammars/domain/grammars/cardinalities"
	"github.com/steve-care-software/webx/grammars/domain/grammars/coverages"
	"github.com/steve-care-software/webx/grammars/domain/grammars/values"
	"github.com/steve-care-software/webx/grammars/domain/trees"
)

type application struct {
	blockchainApp                    applications.Application
	contentAdapter                   contents_grammars.Adapter
	contentBuilder                   contents_grammars.Builder
	contentTokenAdapter              contents_tokens.Adapter
	contentTokenBuilder              contents_tokens.Builder
	contentTokenLinesBuilder         contents_tokens.LinesBuilder
	contentTokenLineBuilder          contents_tokens.LineBuilder
	contentElementAdapter            contents_elements.Adapter
	contentElementBuilder            contents_elements.Builder
	contentElementCardinalityBuilder contents_elements.CardinalityBuilder
	contentEverythingAdapter         contents_everythings.Adapter
	contentEverythingBuilder         contents_everythings.Builder
	contentChannelAdapter            contents_channels.Adapter
	contentChannelBuilder            contents_channels.Builder
	contentMatchAdapter              contents_matches.Adapter
	contentSuiteAdapter              contents_suites.Adapter
	contentSuiteBuilder              contents_suites.Builder
	builder                          grammars.Builder
	grammarTokenBuilder              grammars.TokenBuilder
	grammarBlockBuilder              grammars.BlockBuilder
	grammarLineBuilder               grammars.LineBuilder
	grammarElementBuilder            grammars.ElementBuilder
	grammarInstanceBuilder           grammars.InstanceBuilder
	grammarExternalBuilder           grammars.ExternalBuilder
	grammarEverythingBuilder         grammars.EverythingBuilder
	grammarChannelsBuilder           grammars.ChannelsBuilder
	grammarChannelBuilder            grammars.ChannelBuilder
	grammarChannelConditionBuilder   grammars.ChannelConditionBuilder
	grammarSuitesBuilder             grammars.SuitesBuilder
	grammarSuiteBuilder              grammars.SuiteBuilder
	grammarValueBuilder              values.Builder
	grammarCardinalityBuilder        cardinalities.Builder
	treesBuilder                     trees.Builder
	treeBuilder                      trees.TreeBuilder
	treeBlockBuilder                 trees.BlockBuilder
	treeLineBuilder                  trees.LineBuilder
	treeElementsBuilder              trees.ElementsBuilder
	treeElementBuilder               trees.ElementBuilder
	treeContentsBuilder              trees.ContentsBuilder
	treeContentBuilder               trees.ContentBuilder
	treeValueBuilder                 trees.ValueBuilder
	coveragesBuilder                 coverages.Builder
	coverageBuilder                  coverages.CoverageBuilder
	coverageExecutionsBuilder        coverages.ExecutionsBuilder
	coverageExecutionBuilder         coverages.ExecutionBuilder
	coverageResultBuilder            coverages.ResultBuilder
	hashAdapter                      hash.Adapter
}

func createApplication(
	blockchainApp applications.Application,
	contentAdapter contents_grammars.Adapter,
	contentBuilder contents_grammars.Builder,
	contentTokenAdapter contents_tokens.Adapter,
	contentTokenBuilder contents_tokens.Builder,
	contentTokenLinesBuilder contents_tokens.LinesBuilder,
	contentTokenLineBuilder contents_tokens.LineBuilder,
	contentElementAdapter contents_elements.Adapter,
	contentElementBuilder contents_elements.Builder,
	contentElementCardinalityBuilder contents_elements.CardinalityBuilder,
	contentEverythingAdapter contents_everythings.Adapter,
	contentEverythingBuilder contents_everythings.Builder,
	contentChannelAdapter contents_channels.Adapter,
	contentChannelBuilder contents_channels.Builder,
	contentMatchAdapter contents_matches.Adapter,
	contentSuiteAdapter contents_suites.Adapter,
	contentSuiteBuilder contents_suites.Builder,
	builder grammars.Builder,
	grammarTokenBuilder grammars.TokenBuilder,
	grammarBlockBuilder grammars.BlockBuilder,
	grammarLineBuilder grammars.LineBuilder,
	grammarElementBuilder grammars.ElementBuilder,
	grammarInstanceBuilder grammars.InstanceBuilder,
	grammarExternalBuilder grammars.ExternalBuilder,
	grammarEverythingBuilder grammars.EverythingBuilder,
	grammarChannelsBuilder grammars.ChannelsBuilder,
	grammarChannelBuilder grammars.ChannelBuilder,
	grammarChannelConditionBuilder grammars.ChannelConditionBuilder,
	grammarSuitesBuilder grammars.SuitesBuilder,
	grammarSuiteBuilder grammars.SuiteBuilder,
	grammarValueBuilder values.Builder,
	grammarCardinalityBuilder cardinalities.Builder,
	treesBuilder trees.Builder,
	treeBuilder trees.TreeBuilder,
	treeBlockBuilder trees.BlockBuilder,
	treeLineBuilder trees.LineBuilder,
	treeElementsBuilder trees.ElementsBuilder,
	treeElementBuilder trees.ElementBuilder,
	treeContentsBuilder trees.ContentsBuilder,
	treeContentBuilder trees.ContentBuilder,
	treeValueBuilder trees.ValueBuilder,
	coveragesBuilder coverages.Builder,
	coverageBuilder coverages.CoverageBuilder,
	coverageExecutionsBuilder coverages.ExecutionsBuilder,
	coverageExecutionBuilder coverages.ExecutionBuilder,
	coverageResultBuilder coverages.ResultBuilder,
	hashAdapter hash.Adapter,
) Application {
	out := application{
		blockchainApp:                    blockchainApp,
		contentAdapter:                   contentAdapter,
		contentBuilder:                   contentBuilder,
		contentTokenAdapter:              contentTokenAdapter,
		contentTokenBuilder:              contentTokenBuilder,
		contentTokenLinesBuilder:         contentTokenLinesBuilder,
		contentTokenLineBuilder:          contentTokenLineBuilder,
		contentElementAdapter:            contentElementAdapter,
		contentElementBuilder:            contentElementBuilder,
		contentElementCardinalityBuilder: contentElementCardinalityBuilder,
		contentEverythingAdapter:         contentEverythingAdapter,
		contentEverythingBuilder:         contentEverythingBuilder,
		contentChannelAdapter:            contentChannelAdapter,
		contentChannelBuilder:            contentChannelBuilder,
		contentMatchAdapter:              contentMatchAdapter,
		contentSuiteAdapter:              contentSuiteAdapter,
		contentSuiteBuilder:              contentSuiteBuilder,
		builder:                          builder,
		grammarTokenBuilder:              grammarTokenBuilder,
		grammarBlockBuilder:              grammarBlockBuilder,
		grammarLineBuilder:               grammarLineBuilder,
		grammarElementBuilder:            grammarElementBuilder,
		grammarInstanceBuilder:           grammarInstanceBuilder,
		grammarExternalBuilder:           grammarExternalBuilder,
		grammarEverythingBuilder:         grammarEverythingBuilder,
		grammarChannelsBuilder:           grammarChannelsBuilder,
		grammarChannelBuilder:            grammarChannelBuilder,
		grammarChannelConditionBuilder:   grammarChannelConditionBuilder,
		grammarSuitesBuilder:             grammarSuitesBuilder,
		grammarSuiteBuilder:              grammarSuiteBuilder,
		grammarValueBuilder:              grammarValueBuilder,
		grammarCardinalityBuilder:        grammarCardinalityBuilder,
		treesBuilder:                     treesBuilder,
		treeBuilder:                      treeBuilder,
		treeBlockBuilder:                 treeBlockBuilder,
		treeLineBuilder:                  treeLineBuilder,
		treeElementsBuilder:              treeElementsBuilder,
		treeElementBuilder:               treeElementBuilder,
		treeContentsBuilder:              treeContentsBuilder,
		treeContentBuilder:               treeContentBuilder,
		treeValueBuilder:                 treeValueBuilder,
		coveragesBuilder:                 coveragesBuilder,
		coverageBuilder:                  coverageBuilder,
		coverageExecutionsBuilder:        coverageExecutionsBuilder,
		coverageExecutionBuilder:         coverageExecutionBuilder,
		coverageResultBuilder:            coverageResultBuilder,
		hashAdapter:                      hashAdapter,
	}

	return &out
}

// New creates a new database
func (app *application) New(name string) error {
	return app.blockchainApp.New(name)
}

// Retrieve retrieves a grammar by hash
func (app *application) Retrieve(context uint, hash hash.Hash) (grammars.Grammar, error) {
	content, err := app.blockchainApp.ReadByHash(context, hash)
	if err != nil {
		return nil, err
	}

	contentGrammarIns, err := app.contentAdapter.ToGrammar(content)
	if err != nil {
		return nil, err
	}

	root, err := app.retrieveToken(context, contentGrammarIns.Root())
	if err != nil {
		return nil, err
	}

	builder := app.builder.Create().WithRoot(root)
	if contentGrammarIns.HasChannels() {
		channels, err := app.retrieveChannels(context, contentGrammarIns.Channels())
		if err != nil {
			return nil, err
		}

		builder.WithChannels(channels)
	}

	return builder.Now()
}

func (app *application) retrieveToken(context uint, hash hash.Hash) (grammars.Token, error) {
	content, err := app.blockchainApp.ReadByHash(context, hash)
	if err != nil {
		return nil, err
	}

	contentTokenIns, err := app.contentTokenAdapter.ToToken(content)
	if err != nil {
		return nil, err
	}

	linesList := []grammars.Line{}
	contentLines := contentTokenIns.Lines().List()
	for _, oneLine := range contentLines {
		elementHashes := oneLine.Elements()
		elements, err := app.retrieveElements(context, elementHashes)
		if err != nil {
			return nil, err
		}

		line, err := app.grammarLineBuilder.Create().WithElements(elements).Now()
		if err != nil {
			return nil, err
		}

		linesList = append(linesList, line)
	}

	block, err := app.grammarBlockBuilder.Create().WithLines(linesList).Now()
	if err != nil {
		return nil, err
	}

	name := hash.String()
	builder := app.grammarTokenBuilder.Create().WithName(name).WithBlock(block)
	suites, err := app.retrieveSuitesByToken(context, hash)
	if err == nil {
		builder.WithSuites(suites)
	}

	return builder.Now()
}

func (app *application) retrieveSuitesByToken(context uint, tokenHash hash.Hash) (grammars.Suites, error) {
	pHash, err := app.hashAdapter.FromBytes([]byte(fmt.Sprintf(grammarMatchByTokenPattern, tokenHash.String())))
	if err != nil {
		return nil, err
	}

	content, err := app.blockchainApp.ReadByHash(context, *pHash)
	if err != nil {
		return nil, err
	}

	contentMatch, err := app.contentMatchAdapter.ToMatch(content)
	if err != nil {
		return nil, err
	}

	suiteHashes := contentMatch.Suites()
	return app.retrieveSuites(context, suiteHashes)
}

func (app *application) retrieveSuites(context uint, hashes []hash.Hash) (grammars.Suites, error) {
	contents, err := app.blockchainApp.ReadAllByHashes(context, hashes)
	if err != nil {
		return nil, err
	}

	list := []grammars.Suite{}
	for _, oneContent := range contents {
		suite, err := app.contentToSuite(context, oneContent)
		if err != nil {
			return nil, err
		}

		list = append(list, suite)

	}

	return app.grammarSuitesBuilder.Create().
		WithList(list).
		Now()
}

func (app *application) retrieveSuite(context uint, hash hash.Hash) (grammars.Suite, error) {
	content, err := app.blockchainApp.ReadByHash(context, hash)
	if err != nil {
		return nil, err
	}

	return app.contentToSuite(context, content)
}

func (app *application) contentToSuite(context uint, content []byte) (grammars.Suite, error) {
	contentSuite, err := app.contentSuiteAdapter.ToSuite(content)
	if err != nil {
		return nil, err
	}

	bytes := contentSuite.Content()
	isValid := contentSuite.IsValid()
	builder := app.grammarSuiteBuilder.Create()
	if isValid {
		builder.WithValid(bytes)
	}

	if !isValid {
		builder.WithInvalid(bytes)
	}

	return builder.Now()
}

func (app *application) retrieveElements(context uint, hashes []hash.Hash) ([]grammars.Element, error) {
	contents, err := app.blockchainApp.ReadAllByHashes(context, hashes)
	if err != nil {
		return nil, err
	}

	list := []grammars.Element{}
	for _, oneContent := range contents {
		element, err := app.contentToElement(context, oneContent)
		if err != nil {
			return nil, err
		}

		list = append(list, element)
	}

	return list, nil
}

func (app *application) retrieveElement(context uint, hash hash.Hash) (grammars.Element, error) {
	content, err := app.blockchainApp.ReadByHash(context, hash)
	if err != nil {
		return nil, err
	}

	return app.contentToElement(context, content)
}

func (app *application) contentToElement(context uint, content []byte) (grammars.Element, error) {
	elementContent, err := app.contentElementAdapter.ToElement(content)
	if err != nil {
		return nil, err
	}

	cardinalityContent := elementContent.Cardinality()
	min := cardinalityContent.Min()
	cardinalityBuilder := app.grammarCardinalityBuilder.Create().WithMin(min)
	if cardinalityContent.HasMax() {
		pMax := cardinalityContent.Max()
		cardinalityBuilder.WithMax(*pMax)
	}

	cardinality, err := cardinalityBuilder.Now()
	if err != nil {
		return nil, err
	}

	contentContent := elementContent.Content()
	builder := app.grammarElementBuilder.Create().WithCardinality(cardinality)
	if contentContent.IsValue() {
		pValue := contentContent.Value()
		pHash, err := app.hashAdapter.FromBytes([]byte{
			*pValue,
		})

		if err != nil {
			return nil, err
		}

		value, err := app.grammarValueBuilder.Create().WithName(pHash.String()).WithNumber(*pValue).Now()
		if err != nil {
			return nil, err
		}

		builder.WithValue(value)
	}

	if contentContent.IsExternal() {
		pExternalHash := contentContent.External()
		external, err := app.retrieveExternal(context, *pExternalHash)
		if err != nil {
			return nil, err
		}

		builder.WithExternal(external)
	}

	if contentContent.IsToken() {
		pTokenHash := contentContent.Token()
		token, err := app.retrieveToken(context, *pTokenHash)
		if err != nil {
			return nil, err
		}

		instance, err := app.grammarInstanceBuilder.Create().WithToken(token).Now()
		if err != nil {
			return nil, err
		}

		builder.WithInstance(instance)
	}

	if contentContent.IsEverything() {
		pEverythingHash := contentContent.Everything()
		everything, err := app.retrieveEverything(context, *pEverythingHash)
		if err != nil {
			return nil, err
		}

		instance, err := app.grammarInstanceBuilder.Create().WithEverything(everything).Now()
		if err != nil {
			return nil, err
		}

		builder.WithInstance(instance)
	}

	if contentContent.IsRecursive() {
		pRecursiveHash := contentContent.Recursive()
		builder.WithRecursive(pRecursiveHash.String())
	}

	return builder.Now()
}

func (app *application) retrieveExternal(context uint, hash hash.Hash) (grammars.External, error) {
	grammar, err := app.Retrieve(context, hash)
	if err != nil {
		return nil, err
	}

	name := hash.String()
	return app.grammarExternalBuilder.Create().
		WithName(name).
		WithGrammar(grammar).
		Now()
}

func (app *application) retrieveEverything(context uint, hash hash.Hash) (grammars.Everything, error) {
	content, err := app.blockchainApp.ReadByHash(context, hash)
	if err != nil {
		return nil, err
	}

	everythingContent, err := app.contentEverythingAdapter.ToEverything(content)
	if err != nil {
		return nil, err
	}

	exceptionHash := everythingContent.Exception()
	exception, err := app.retrieveToken(context, exceptionHash)
	if err != nil {
		return nil, err
	}

	name := hash.String()
	builder := app.grammarEverythingBuilder.Create().WithName(name).WithException(exception)
	if everythingContent.HasEscape() {
		pEscapeHash := everythingContent.Escape()
		escape, err := app.retrieveToken(context, *pEscapeHash)
		if err != nil {
			return nil, err
		}

		builder.WithEscape(escape)
	}

	return builder.Now()
}

func (app *application) retrieveChannels(context uint, hashes []hash.Hash) (grammars.Channels, error) {
	contents, err := app.blockchainApp.ReadAllByHashes(context, hashes)
	if err != nil {
		return nil, err
	}

	list := []grammars.Channel{}
	for _, oneContent := range contents {
		channel, err := app.contentToChannel(context, oneContent)
		if err != nil {
			return nil, err
		}

		list = append(list, channel)
	}

	return app.grammarChannelsBuilder.Create().
		WithList(list).
		Now()
}

func (app *application) retrieveChannel(context uint, hash hash.Hash) (grammars.Channel, error) {
	content, err := app.blockchainApp.ReadByHash(context, hash)
	if err != nil {
		return nil, err
	}

	return app.contentToChannel(context, content)
}

func (app *application) contentToChannel(context uint, content []byte) (grammars.Channel, error) {
	channelContent, err := app.contentChannelAdapter.ToChannel(content)
	if err != nil {
		return nil, err
	}

	token, err := app.retrieveToken(context, channelContent.Token())
	if err != nil {
		return nil, err
	}

	builder := app.grammarChannelBuilder.Create().WithToken(token)
	hasPrevious := channelContent.HasPrevious()
	hasNext := channelContent.HasNext()
	if hasPrevious || hasNext {
		conditionBuilder := app.grammarChannelConditionBuilder.Create()
		if hasPrevious {
			pPreviousHash := channelContent.Previous()
			previous, err := app.retrieveToken(context, *pPreviousHash)
			if err != nil {
				return nil, err
			}

			conditionBuilder.WithPrevious(previous)
		}

		if hasNext {
			pNextHash := channelContent.Previous()
			next, err := app.retrieveToken(context, *pNextHash)
			if err != nil {
				return nil, err
			}

			conditionBuilder.WithNext(next)
		}

		condition, err := conditionBuilder.Now()
		if err != nil {
			return nil, err
		}

		builder.WithCondition(condition)
	}

	return builder.Now()
}

// Scan scans all the tokens to find matches for our suites, when they do, insert the suite in the database
func (app *application) Scan(context uint, suites grammars.Suites) (grammars.Grammar, error) {
	return app.ScanWithChannels(context, suites, nil)
}

// ScanWithChannels executes a scan with channels
func (app *application) ScanWithChannels(context uint, suites grammars.Suites, channels grammars.Channels) (grammars.Grammar, error) {
	// retrieve the token content keys:
	contentKeys, err := app.blockchainApp.ContentKeys(context, references.KindGrammarToken)
	if err != nil {
		return nil, err
	}

	// for each contentKey:
	var selected grammars.Grammar
	list := contentKeys.List()
	for _, oneContentKey := range list {
		// retrieve the token:
		token, err := app.retrieveToken(context, oneContentKey.Hash())
		if err != nil {
			return nil, err
		}

		// build the grammar instance:
		builder := app.builder.Create().WithRoot(token)
		if channels != nil {
			builder.WithChannels(channels)
		}

		grammar, err := builder.Now()
		if err != nil {
			return nil, err
		}

		// execute the suite coverages:
		coverages, err := app.Coverages(grammar)
		if err != nil {
			return nil, err
		}

		if coverages.ContainsError() {
			continue
		}

		// the suite passed:
		if selected == nil {
			selected = grammar
			continue
		}

		// we always keep the valid grammar with the lowest amount of points:
		if selected.Points() > grammar.Points() {
			selected = grammar
		}
	}

	if selected == nil {
		return nil, errors.New("there is no token that matches the provided suites and channels and therefore no grammar could be returned")
	}

	// returns the select grammar:
	return selected, nil
}

// Insert inserts a grammar
func (app *application) Insert(context uint, grammar grammars.Grammar) error {
	contents, err := app.grammarToBytes(context, grammar)
	if err != nil {
		return err
	}

	if len(contents) > 0 {
		str := fmt.Sprintf("the grammar (hash: %s) cannot be inserted because it already exists", grammar.Hash().String())
		return errors.New(str)
	}

	return app.blockchainApp.WriteAll(context, contents)
}

func (app *application) grammarToBytes(context uint, grammar grammars.Grammar) ([][]byte, error) {
	output := [][]byte{}

	// if the grammar already exists:
	_, err := app.Retrieve(context, grammar.Hash())
	if err == nil {
		return output, nil
	}

	// root token:
	rootToken := grammar.Root()
	_, err = app.retrieveToken(context, rootToken.Hash())
	if err != nil {
		tokenBytes, err := app.tokenToBytes(context, rootToken, map[string]hash.Hash{})
		if err != nil {
			return nil, err
		}

		if len(tokenBytes) > 0 {
			output = append(output, tokenBytes...)
		}
	}

	// channels:
	if grammar.HasChannels() {
		channels := grammar.Channels()
		channelsBytes, err := app.channelsToBytes(context, channels)
		if err != nil {
			return nil, err
		}

		if len(channelsBytes) > 0 {
			output = append(output, channelsBytes...)
		}
	}

	grammarHash := grammar.Hash()
	root := grammar.Root().Hash()

	channels := []hash.Hash{}
	if grammar.HasChannels() {
		channelsList := grammar.Channels().List()
		for _, oneChannel := range channelsList {
			channels = append(channels, oneChannel.Hash())
		}
	}

	builder := app.contentBuilder.Create().WithHash(grammarHash).WithRoot(root)
	if len(channels) > 0 {
		builder.WithChannels(channels)
	}

	content, err := builder.Now()
	if err != nil {
		return nil, err
	}

	grammarBytes, err := app.contentAdapter.ToContent(content)
	if err != nil {
		return nil, err
	}

	return append(output, grammarBytes), nil
}

func (app *application) tokenToBytes(context uint, token grammars.Token, recursives map[string]hash.Hash) ([][]byte, error) {
	output := [][]byte{}

	name := token.Name()
	if _, ok := recursives[name]; !ok {
		recursives[name] = token.Hash()
	}

	// elements:
	contentLinesList := []contents_tokens.Line{}
	blockLinesList := token.Block().Lines()
	for _, oneLine := range blockLinesList {
		elementHashes := []hash.Hash{}
		elementsList := oneLine.Elements()
		for _, oneElement := range elementsList {
			elementHash := oneElement.Hash()
			_, err := app.retrieveElement(context, elementHash)
			if err != nil {
				elementBytes, err := app.elementToBytes(context, oneElement, recursives)
				if err != nil {
					return nil, err
				}

				if len(elementBytes) > 0 {
					output = append(output, elementBytes...)
				}
			}

			elementHashes = append(elementHashes, elementHash)
		}

		contentLine, err := app.contentTokenLineBuilder.Create().WithElements(elementHashes).Now()
		if err != nil {
			return nil, err
		}

		contentLinesList = append(contentLinesList, contentLine)
	}

	contentLines, err := app.contentTokenLinesBuilder.Create().WithList(contentLinesList).Now()
	if err != nil {
		return nil, err
	}

	hash := token.Hash()
	contentToken, err := app.contentTokenBuilder.Create().WithHash(hash).WithLines(contentLines).Now()
	if err != nil {
		return nil, err
	}

	tokenBytes, err := app.contentTokenAdapter.ToContent(contentToken)
	if err != nil {
		return nil, err
	}

	return append(output, tokenBytes), nil
}

func (app *application) suitesToBytes(context uint, suites grammars.Suites) ([][]byte, error) {
	output := [][]byte{}
	list := suites.List()
	for _, oneSuite := range list {
		content, err := app.suiteToBytes(context, oneSuite)
		if err != nil {
			return nil, err
		}

		if len(content) <= 0 {
			continue
		}

		output = append(output, content)
	}

	return output, nil
}

func (app *application) suiteToBytes(context uint, suite grammars.Suite) ([]byte, error) {
	hash := suite.Hash()
	_, err := app.retrieveSuite(context, hash)
	if err == nil {
		return []byte{}, nil
	}

	content := suite.Content()
	builder := app.contentSuiteBuilder.Create().WithHash(hash).WithContent(content)
	if suite.IsValid() {
		builder.IsValid()
	}

	contentSuite, err := builder.Now()
	if err != nil {
		return nil, err
	}

	return app.contentSuiteAdapter.ToContent(contentSuite)
}

func (app *application) elementToBytes(context uint, element grammars.Element, recursives map[string]hash.Hash) ([][]byte, error) {
	output := [][]byte{}

	cardinality := element.Cardinality()
	min := cardinality.Min()
	contentCardinalityBuilder := app.contentElementCardinalityBuilder.Create().WithMin(min)
	if cardinality.HasMax() {
		pMax := cardinality.Max()
		contentCardinalityBuilder.WithMax(*pMax)
	}

	contentCardinality, err := contentCardinalityBuilder.Now()
	if err != nil {
		return nil, err
	}

	hash := element.Hash()
	content := element.Content()
	builder := app.contentElementBuilder.Create().WithHash(hash).WithCardinality(contentCardinality)
	if content.IsValue() {
		number := content.Value().Number()
		builder.WithValue(number)
	}

	if content.IsExternal() {
		grammar := content.External().Grammar()
		grammarHash := grammar.Hash()
		_, err := app.Retrieve(context, grammarHash)
		if err != nil {
			grammarBytes, err := app.grammarToBytes(context, grammar)
			if err != nil {
				return nil, err
			}

			if len(grammarBytes) > 0 {
				output = append(output, grammarBytes...)
			}
		}

		builder.WithExternal(grammarHash)
	}

	if content.IsInstance() {
		instance := content.Instance()
		if instance.IsToken() {
			token := instance.Token()
			tokenHash := token.Hash()
			_, err := app.retrieveToken(context, tokenHash)
			if err != nil {
				tokenBytes, err := app.tokenToBytes(context, token, recursives)
				if err != nil {
					return nil, err
				}

				if len(tokenBytes) > 0 {
					output = append(output, tokenBytes...)
				}
			}

			builder.WithToken(tokenHash)
		}

		if instance.IsEverything() {
			everything := instance.Everything()
			everythingHash := everything.Hash()
			_, err := app.retrieveEverything(context, everythingHash)
			if err != nil {
				everythingBytes, err := app.everythingToBytes(context, everything, recursives)
				if err != nil {
					return nil, err
				}

				if len(everythingBytes) > 0 {
					output = append(output, everythingBytes...)
				}
			}

			builder.WithEverything(everythingHash)
		}
	}

	if content.IsRecursive() {
		name := content.Recursive()
		if hash, ok := recursives[name]; ok {
			builder.WithRecursive(hash)
		}

		str := fmt.Sprintf("the recursive token (name: %s) could notbe found in the recursive stack", name)
		return nil, errors.New(str)
	}

	contentElement, err := builder.Now()
	if err != nil {
		return nil, err
	}

	elementBytes, err := app.contentElementAdapter.ToContent(contentElement)
	if err != nil {
		return nil, err
	}

	return append(output, elementBytes), nil
}

func (app *application) everythingToBytes(context uint, everything grammars.Everything, recursives map[string]hash.Hash) ([][]byte, error) {
	output := [][]byte{}

	// eception:
	exception := everything.Exception()
	exceptionHash := exception.Hash()
	_, err := app.retrieveToken(context, exceptionHash)
	if err != nil {
		tokenBytes, err := app.tokenToBytes(context, exception, recursives)
		if err != nil {
			return nil, err
		}

		if len(tokenBytes) > 0 {
			output = append(output, tokenBytes...)
		}
	}

	// escape:
	builder := app.contentEverythingBuilder.Create().WithException(exceptionHash)
	if everything.HasEscape() {
		escape := everything.Escape()
		escapeHash := escape.Hash()
		_, err := app.retrieveToken(context, escapeHash)
		if err != nil {
			tokenBytes, err := app.tokenToBytes(context, escape, recursives)
			if err != nil {
				return nil, err
			}

			if len(tokenBytes) > 0 {
				output = append(output, tokenBytes...)
			}
		}

		builder.WithEscape(escapeHash)
	}

	contentEverything, err := builder.Now()
	if err != nil {
		return nil, err
	}

	contentBytes, err := app.contentEverythingAdapter.ToContent(contentEverything)
	if err != nil {
		return nil, err
	}

	return append(output, contentBytes), nil
}

func (app *application) channelsToBytes(context uint, channels grammars.Channels) ([][]byte, error) {
	output := [][]byte{}
	channelsList := channels.List()
	recursives := map[string]hash.Hash{}
	for _, oneChannel := range channelsList {
		bytes, err := app.channelToBytes(context, oneChannel, recursives)
		if err != nil {
			return nil, err
		}

		if len(bytes) <= 0 {
			continue
		}

		output = append(output, bytes...)
	}

	return output, nil
}

func (app *application) channelToBytes(context uint, channel grammars.Channel, recursives map[string]hash.Hash) ([][]byte, error) {
	output := [][]byte{}

	token := channel.Token()
	tokenHash := token.Hash()
	_, err := app.retrieveToken(context, tokenHash)
	if err != nil {
		tokenBytes, err := app.tokenToBytes(context, token, recursives)
		if err != nil {
			return nil, err
		}

		if len(tokenBytes) > 0 {
			output = append(output, tokenBytes...)
		}
	}

	hash := channel.Hash()
	builder := app.contentChannelBuilder.Create().WithHash(hash).WithToken(tokenHash)
	if channel.HasCondition() {
		condition := channel.Condition()
		if condition.HasPrevious() {
			previous := condition.Previous()
			previousHash := previous.Hash()
			_, err := app.retrieveToken(context, previousHash)
			if err != nil {
				tokenBytes, err := app.tokenToBytes(context, previous, recursives)
				if err != nil {
					return nil, err
				}

				if len(tokenBytes) > 0 {
					output = append(output, tokenBytes...)
				}
			}

			builder.WithPrevious(previousHash)
		}

		if condition.HasNext() {
			next := condition.Next()
			nextHash := next.Hash()
			_, err := app.retrieveToken(context, nextHash)
			if err != nil {
				tokenBytes, err := app.tokenToBytes(context, next, recursives)
				if err != nil {
					return nil, err
				}

				if len(tokenBytes) > 0 {
					output = append(output, tokenBytes...)
				}
			}

			builder.WithNext(nextHash)
		}
	}

	contentChannel, err := builder.Now()
	if err != nil {
		return nil, err
	}

	channelBytes, err := app.contentChannelAdapter.ToContent(contentChannel)
	if err != nil {
		return nil, err
	}

	return append(output, channelBytes), nil
}

// Execute executes grammar on data
func (app *application) Execute(grammar grammars.Grammar, values []byte) (trees.Tree, error) {
	return app.grammar(grammar, false, []byte{}, values)
}

// Coverages returns the coverages of a grammar
func (app *application) Coverages(grammar grammars.Grammar) (coverages.Coverages, error) {
	root := grammar.Root()
	channels := grammar.Channels()
	skip := map[string]bool{}
	rootCoverages, err := app.coveragesToken(root, channels, &skip)
	if err != nil {
		return nil, err
	}

	list := []coverages.Coverage{}
	if grammar.HasChannels() {
		channels := grammar.Channels().List()
		for _, oneChannel := range channels {
			token := oneChannel.Token()
			coverages, err := app.coveragesToken(token, nil, &skip)
			if err != nil {
				return nil, err
			}

			if coverages != nil {
				list = append(list, coverages.List()...)
			}
		}
	}

	if rootCoverages != nil {
		list = append(list, rootCoverages.List()...)
	}

	if len(list) <= 0 {
		return nil, nil
	}

	return app.coveragesBuilder.Create().WithList(list).Now()
}

// Covered returns the covered tokens
func (app *application) Covered(coverages coverages.Coverages) (map[string]map[uint]map[uint]string, error) {
	coveredElements := map[string]map[uint]map[uint]string{}
	err := app.findCoveraredElements(coverages, &coveredElements)
	if err != nil {
		return nil, err
	}

	return coveredElements, nil
}

// Uncovered returns the uncovered tokens
func (app *application) Uncovered(grammar grammars.Grammar) (map[string]map[uint]map[uint]string, error) {
	coverages, err := app.Coverages(grammar)
	if err != nil {
		return nil, err
	}

	coveredElements, err := app.Covered(coverages)
	if err != nil {
		return nil, err
	}

	allElements := map[string]map[uint]map[uint]string{}
	err = app.findElements(grammar, &allElements)
	if err != nil {
		return nil, err
	}

	uncoveredElements := map[string]map[uint]map[uint]string{}
	for tokenName, lines := range allElements {
		for lineIdx, elements := range lines {
			for elIdx, element := range elements {
				if _, ok := coveredElements[tokenName][lineIdx][elIdx]; !ok {
					if _, ok := uncoveredElements[tokenName]; !ok {
						uncoveredElements[tokenName] = map[uint]map[uint]string{}
					}

					if _, ok := uncoveredElements[tokenName][lineIdx]; !ok {
						uncoveredElements[tokenName][lineIdx] = map[uint]string{}
					}

					uncoveredElements[tokenName][lineIdx][elIdx] = element
				}
			}
		}

	}

	return uncoveredElements, nil
}

func (app *application) coveragesToken(token grammars.Token, channels grammars.Channels, pSkip *map[string]bool) (coverages.Coverages, error) {
	name := token.Name()
	skip := *pSkip
	if _, ok := skip[name]; ok {
		return nil, nil
	}

	skip[name] = true
	pSkip = &skip
	executionsList := []coverages.Execution{}
	if token.HasSuites() {
		suites := token.Suites().List()
		for _, oneSuite := range suites {
			execution, err := app.coverageTokenSuite(token, channels, oneSuite)
			if err != nil {
				return nil, err
			}

			if execution == nil {
				continue
			}

			executionsList = append(executionsList, execution)
		}
	}

	list := []coverages.Coverage{}
	lines := token.Block().Lines()
	for _, oneLine := range lines {
		elements := oneLine.Elements()
		for _, oneElement := range elements {
			content := oneElement.Content()
			if content.IsExternal() {
				grammar := content.External().Grammar()
				coverages, err := app.Coverages(grammar)
				if err != nil {
					return nil, err
				}

				if coverages != nil {
					list = append(list, coverages.List()...)
				}
			}

			if content.IsInstance() {
				instance := content.Instance()
				if instance.IsToken() {
					token := instance.Token()
					coverages, err := app.coveragesToken(token, channels, pSkip)
					if err != nil {
						return nil, err
					}

					if coverages != nil {
						list = append(list, coverages.List()...)
					}
				}

				if instance.IsEverything() {
					everything := instance.Everything()
					exception := everything.Exception()
					coverages, err := app.coveragesToken(exception, channels, pSkip)
					if err != nil {
						return nil, err
					}

					if coverages != nil {
						list = append(list, coverages.List()...)
					}

					if everything.HasEscape() {
						escape := everything.Escape()
						coverages, err := app.coveragesToken(escape, channels, pSkip)
						if err != nil {
							return nil, err
						}

						if coverages != nil {
							list = append(list, coverages.List()...)
						}
					}
				}
			}

		}
	}

	if len(executionsList) > 0 {
		executions, err := app.coverageExecutionsBuilder.Create().WithList(executionsList).Now()
		if err != nil {
			return nil, err
		}

		coverage, err := app.coverageBuilder.Create().WithToken(token).WithExecutions(executions).Now()
		if err != nil {
			return nil, err
		}

		list = append(list, coverage)
	}

	if len(list) <= 0 {
		return nil, nil
	}

	return app.coveragesBuilder.Create().WithList(list).Now()
}

func (app *application) coverageTokenSuite(token grammars.Token, channels grammars.Channels, suite grammars.Suite) (coverages.Execution, error) {
	input := suite.Content()
	tree, _, err := app.token(token, map[string]*stack{}, nil, channels, false, []byte{}, input)
	if err != nil {
		return nil, err
	}

	resultBuilder := app.coverageResultBuilder.Create()
	if tree != nil {
		resultBuilder.WithTree(tree)
	}

	if err != nil {
		resultBuilder.WithError(err.Error())
	}

	result, err := resultBuilder.Now()
	if err != nil {
		return nil, err
	}

	return app.coverageExecutionBuilder.Create().
		WithExpectation(suite).
		WithResult(result).
		Now()
}

func (app *application) findElements(grammar grammars.Grammar, pElements *map[string]map[uint]map[uint]string) error {
	elements := *pElements
	root := grammar.Root()
	err := app.findElementsFromToken(root, &elements)
	if err != nil {
		return err
	}

	if grammar.HasChannels() {
		channels := grammar.Channels().List()
		for _, oneChannel := range channels {
			token := oneChannel.Token()
			err := app.findElementsFromToken(token, &elements)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func (app *application) findElementsFromToken(token grammars.Token, pElements *map[string]map[uint]map[uint]string) error {
	elements := *pElements
	tokenName := token.Name()
	if _, ok := elements[tokenName]; !ok {
		elements[tokenName] = map[uint]map[uint]string{}
	}

	lines := token.Block().Lines()
	for idx, oneLine := range lines {
		castedIdx := uint(idx)
		if _, ok := elements[tokenName][castedIdx]; !ok {
			elements[tokenName][castedIdx] = map[uint]string{}
		}

		elementsList := oneLine.Elements()
		for elIdx, oneElement := range elementsList {
			castedElIdx := uint(elIdx)
			elements[tokenName][castedIdx][castedElIdx] = oneElement.Name()

			content := oneElement.Content()
			if content.IsExternal() {
				grammar := content.External().Grammar()
				err := app.findElements(grammar, &elements)
				if err != nil {
					return err
				}
			}

			if content.IsInstance() {
				instance := content.Instance()
				if instance.IsToken() {
					token := instance.Token()
					err := app.findElementsFromToken(token, &elements)
					if err != nil {
						return err
					}
				}

				if instance.IsEverything() {
					everything := instance.Everything()
					exception := everything.Exception()
					err := app.findElementsFromToken(exception, &elements)
					if err != nil {
						return err
					}

					if everything.HasEscape() {
						escape := everything.Escape()
						err := app.findElementsFromToken(escape, &elements)
						if err != nil {
							return err
						}
					}
				}
			}

		}
	}

	pElements = &elements
	return nil
}

func (app *application) findCoveraredElements(coverages coverages.Coverages, pCovered *map[string]map[uint]map[uint]string) error {
	list := coverages.List()
	for _, oneCoverage := range list {
		tokenName := oneCoverage.Token().Name()
		executionsList := oneCoverage.Executions().List()
		for _, oneExecution := range executionsList {
			result := oneExecution.Result()
			if !result.IsTree() {
				continue
			}

			block := result.Tree().Block()
			err := app.findCoveraredElementsFromBlock(tokenName, block, pCovered)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func (app *application) findCoveraredElementsFromBlock(tokenName string, block trees.Block, pCovered *map[string]map[uint]map[uint]string) error {
	if !block.HasSuccessful() {
		return nil
	}

	line := block.Successful()
	index := line.Index()
	elementsList := line.Elements().List()
	for elIdx, oneElement := range elementsList {
		if !oneElement.HasGrammar() {
			continue
		}

		elementName := oneElement.Grammar().Name()
		covered := *pCovered
		if _, ok := covered[tokenName]; !ok {
			covered[tokenName] = map[uint]map[uint]string{}
		}

		if _, ok := covered[tokenName][index]; !ok {
			covered[tokenName][index] = map[uint]string{}
		}

		castedElIdx := uint(elIdx)
		if _, ok := covered[tokenName][index][castedElIdx]; !ok {
			covered[tokenName][index][castedElIdx] = elementName
		}

		contents := oneElement.Contents().List()
		for _, oneContent := range contents {
			if oneContent.IsTree() {
				subBlock := oneContent.Tree().Block()
				err := app.findCoveraredElementsFromBlock(elementName, subBlock, &covered)
				if err != nil {
					return err
				}
			}
		}

		pCovered = &covered
	}

	return nil
}

func (app *application) grammar(grammar grammars.Grammar, isReverse bool, prevData []byte, currentData []byte) (trees.Tree, error) {
	root := grammar.Root()
	channels := grammar.Channels()
	tree, _, err := app.token(root, map[string]*stack{}, nil, channels, isReverse, prevData, currentData)
	if err != nil {
		return nil, err
	}

	return tree, nil
}

func (app *application) token(token grammars.Token, stackMap map[string]*stack, escape grammars.Token, channels grammars.Channels, isReverse bool, prevData []byte, currentData []byte) (trees.Tree, map[string]*stack, error) {
	tokenName := token.Name()
	if _, ok := stackMap[tokenName]; !ok {
		stackMap[tokenName] = &stack{
			token: token,
			lines: map[int][]byte{},
		}
	}

	tokenBlock := token.Block()
	block, remaining, retStackMap, err := app.block(token, stackMap, tokenBlock, escape, channels, isReverse, prevData, currentData)
	delete(stackMap, tokenName)
	if err != nil {
		return nil, nil, err
	}

	stackMap = retStackMap
	if block == nil {
		str := fmt.Sprintf("there was no line discovered in the token (name: %s) using the given data: %s", token.Name(), currentData)
		return nil, nil, errors.New(str)
	}

	builder := app.treeBuilder.Create().WithGrammar(token).WithBlock(block)
	if channels != nil {
		suffix, rem, err := app.channels(channels, prevData, remaining)
		if err == nil {
			builder.WithSuffix(suffix)
			remaining = rem
		}
	}

	if len(remaining) > 0 {
		builder.WithRemaining(remaining)
	}

	ins, err := builder.Now()
	if err != nil {
		return nil, nil, err
	}

	return ins, stackMap, nil
}

func (app *application) external(external grammars.External, isReverse bool, prevData []byte, currentData []byte) (trees.Tree, error) {
	grammar := external.Grammar()
	treeIns, err := app.grammar(grammar, isReverse, prevData, currentData)
	if err != nil {
		return nil, err
	}

	name := external.Name()
	root := grammar.Root()
	block := root.Block()
	if block == nil {
		return nil, nil
	}

	builder := app.grammarTokenBuilder.Create().WithName(name).WithBlock(block)
	if root.HasSuites() {
		suites := root.Suites()
		builder.WithSuites(suites)
	}

	grammarRoot, err := builder.Now()
	if err != nil {
		return nil, err
	}

	treeBlock := treeIns.Block()
	return app.treeBuilder.Create().WithGrammar(grammarRoot).WithBlock(treeBlock).Now()
}

func (app *application) block(token grammars.Token, stackMap map[string]*stack, block grammars.Block, escape grammars.Token, channels grammars.Channels, isReverse bool, prevData []byte, currentData []byte) (trees.Block, []byte, map[string]*stack, error) {
	tokenName := token.Name()
	list := []trees.Line{}
	lines := block.Lines()
	remaining := currentData
	currentStack := stackMap

	for idx, oneLine := range lines {
		// if we already went through this line, with the same data, in the stack, skip it to avoid infinite loops:
		if _, ok := currentStack[tokenName]; ok {
			if data, ok := currentStack[tokenName].lines[idx]; ok {
				if bytes.Compare(remaining, data) == 0 {
					continue
				}

			}
		}

		if _, ok := currentStack[tokenName]; !ok {
			currentStack[tokenName] = &stack{
				token: token,
				lines: map[int][]byte{},
			}
		}

		currentStack[tokenName].lines[idx] = remaining

		// if the line is in reverse:
		if isReverse {
			previousData := prevData
			contentsList := []trees.Content{}

			for {
				if len(remaining) <= 0 {
					break
				}

				if escape != nil {
					escapeTree, _, err := app.token(escape, stackMap, nil, channels, false, previousData, remaining)
					if err == nil {
						if escapeTree.Block().HasSuccessful() {
							if escapeTree.HasRemaining() {
								escapeRemaining := escapeTree.Remaining()
								treeLine, rem, _, err := app.line(tokenName, currentStack, oneLine, uint(idx), escape, channels, isReverse, remaining, escapeRemaining)
								if err == nil && treeLine.IsSuccessful() {
									amount := len(escapeRemaining) - len(rem)
									values := escapeRemaining[:amount]
									for _, oneValue := range values {
										value, err := app.treeValueBuilder.Create().WithContent(oneValue).Now()
										if err != nil {
											return nil, nil, nil, err
										}

										contentIns, err := app.treeContentBuilder.Create().WithValue(value).Now()
										if err != nil {
											return nil, nil, nil, err
										}

										contentsList = append(contentsList, contentIns)
									}

									previousData = escapeRemaining
									remaining = escapeRemaining[amount:]
								}
							}
						}
					}
				}

				_, _, _, err := app.line(tokenName, currentStack, oneLine, uint(idx), escape, channels, isReverse, previousData, remaining)
				if err == nil {
					break
				}

				value, err := app.treeValueBuilder.Create().WithContent(remaining[0]).Now()
				if err != nil {
					return nil, nil, nil, err
				}

				contentIns, err := app.treeContentBuilder.Create().WithValue(value).Now()
				if err != nil {
					return nil, nil, nil, err
				}

				contentsList = append(contentsList, contentIns)
				previousData = remaining
				remaining = remaining[1:]
			}

			contents, err := app.treeContentsBuilder.Create().WithList(contentsList).Now()
			if err != nil {
				return nil, nil, nil, err
			}

			elementIns, err := app.treeElementBuilder.Create().WithContents(contents).Now()
			if err != nil {
				return nil, nil, nil, err
			}

			elements, err := app.treeElementsBuilder.Create().WithList([]trees.Element{
				elementIns,
			}).Now()
			if err != nil {
				return nil, nil, nil, err
			}

			lineIns, err := app.treeLineBuilder.Create().
				WithIndex(uint(idx)).
				WithGrammar(oneLine).
				WithElements(elements).
				IsReverse().
				Now()

			if err != nil {
				return nil, nil, nil, err
			}

			list = append(list, lineIns)
			break
		}

		// the line is NOT in reverse:
		lineIns, rem, retStack, err := app.line(tokenName, currentStack, oneLine, uint(idx), escape, channels, isReverse, prevData, remaining)
		if err != nil {
			continue
		}

		// add the line to the list:
		list = append(list, lineIns)
		if lineIns.IsSuccessful() {
			remaining = rem
			currentStack = retStack
			break
		}
	}

	// if there is no line:
	if len(list) <= 0 {
		return nil, remaining, currentStack, nil
	}

	blockIns, err := app.treeBlockBuilder.Create().WithLines(list).Now()
	if err != nil {
		return nil, nil, nil, err
	}

	return blockIns, remaining, currentStack, nil
}

func (app *application) line(tokenName string, stackMap map[string]*stack, line grammars.Line, index uint, escape grammars.Token, channels grammars.Channels, isReverse bool, prevData []byte, currentData []byte) (trees.Line, []byte, map[string]*stack, error) {
	list := []trees.Element{}
	grElements := line.Elements()
	remaining := currentData
	previousData := prevData
	currentStack := stackMap
	for _, oneElement := range grElements {
		contentsList := []trees.Content{}
		cardinality := oneElement.Cardinality()
		pMax := cardinality.Max()
		for {

			if len(remaining) <= 0 {
				break
			}

			if cardinality.HasMax() {
				amount := uint(len(contentsList))
				if amount >= *pMax {
					break
				}
			}

			contentIns, rem, retStack, err := app.element(tokenName, oneElement, currentStack, escape, channels, isReverse, previousData, remaining)
			if err != nil {
				break
			}

			currentStack = retStack
			contentsList = append(contentsList, contentIns)
			previousData = remaining
			remaining = rem
		}

		contents, err := app.treeContentsBuilder.Create().WithList(contentsList).Now()
		if err != nil {
			return nil, nil, nil, err
		}

		elementIns, err := app.treeElementBuilder.Create().WithGrammar(oneElement).WithContents(contents).Now()
		if err != nil {
			return nil, nil, nil, err
		}

		list = append(list, elementIns)
	}

	builder := app.treeLineBuilder.Create().
		WithIndex(index).
		WithGrammar(line)

	if len(list) > 0 {
		elements, err := app.treeElementsBuilder.Create().WithList(list).Now()
		if err != nil {
			return nil, nil, nil, err
		}

		builder.WithElements(elements)
	}

	lineIns, err := builder.Now()
	if err != nil {
		return nil, nil, nil, err
	}

	return lineIns, remaining, currentStack, nil
}

func (app *application) element(tokenName string, element grammars.Element, stackMap map[string]*stack, escape grammars.Token, channels grammars.Channels, isReverse bool, prevData []byte, currentData []byte) (trees.Content, []byte, map[string]*stack, error) {
	if len(currentData) <= 0 {
		return nil, nil, nil, errors.New("no remaining data")
	}

	content := element.Content()
	value, tree, rem, retStack, err := app.elementContent(tokenName, content, stackMap, escape, channels, isReverse, prevData, currentData)
	if err != nil {
		return nil, nil, nil, err
	}

	if value == nil && tree == nil {
		return nil, nil, nil, errors.New("no value/tree found")
	}

	if tree != nil && !tree.Block().HasSuccessful() {
		return nil, nil, nil, errors.New("no successfull tree found")
	}

	contentBuilder := app.treeContentBuilder.Create()
	if value != nil {
		contentBuilder.WithValue(value)
	}

	if tree != nil {
		contentBuilder.WithTree(tree)
	}

	contentIns, err := contentBuilder.Now()
	if err != nil {
		return nil, nil, nil, err
	}

	return contentIns, rem, retStack, nil
}

func (app *application) elementContent(tokenName string, content grammars.ElementContent, stackMap map[string]*stack, escape grammars.Token, channels grammars.Channels, isReverse bool, prevData []byte, currentData []byte) (trees.Value, trees.Tree, []byte, map[string]*stack, error) {
	if content.IsExternal() {
		external := content.External()
		tree, err := app.external(external, isReverse, prevData, currentData)
		if err != nil {
			return nil, nil, nil, nil, err
		}

		if tree == nil {
			return nil, nil, nil, nil, nil
		}

		remaining := []byte{}
		if tree.HasRemaining() {
			remaining = tree.Remaining()
		}

		return nil, tree, remaining, stackMap, nil
	}

	if content.IsInstance() {
		instance := content.Instance()
		tree, retStack, err := app.instance(instance, stackMap, escape, channels, isReverse, prevData, currentData)
		if err != nil {
			return nil, nil, nil, nil, err
		}

		remaining := []byte{}
		if tree.HasRemaining() {
			remaining = tree.Remaining()
		}

		return nil, tree, remaining, retStack, nil
	}

	if content.IsRecursive() {
		recursive := content.Recursive()
		if stack, ok := stackMap[recursive]; ok {
			tree, retStack, err := app.token(stack.token, stackMap, escape, channels, isReverse, prevData, currentData)
			if err != nil {
				return nil, nil, nil, nil, err
			}

			remaining := []byte{}
			if tree.HasRemaining() {
				remaining = tree.Remaining()
			}

			return nil, tree, remaining, retStack, nil
		}

		str := fmt.Sprintf("the token (name: %s) was expected to be recursive, but it is not in the current stack", recursive)
		return nil, nil, nil, nil, errors.New(str)
	}

	if len(currentData) < 1 {
		return nil, nil, nil, nil, errors.New("there must be at least 1 value in the given data in order to have an element match, 0 provided")
	}

	remaining := currentData
	builder := app.treeValueBuilder.Create()
	if channels != nil {
		prefix, rem, err := app.channels(channels, prevData, remaining)
		if err == nil {
			builder.WithPrefix(prefix)
			remaining = rem
		}
	}

	if len(remaining) < 1 {
		return nil, nil, nil, nil, errors.New("there must be at least 1 value in the given data in order to have an element match, 0 provided")
	}

	number := content.Value().Number()
	if number == remaining[0] {
		ins, err := builder.WithContent(remaining[0]).Now()
		if err != nil {
			return nil, nil, nil, nil, err
		}

		return ins, nil, remaining[1:], stackMap, nil
	}

	return nil, nil, nil, nil, nil
}

func (app *application) instance(instance grammars.Instance, stackMap map[string]*stack, escape grammars.Token, channels grammars.Channels, isReverse bool, prevData []byte, currentData []byte) (trees.Tree, map[string]*stack, error) {
	if instance.IsToken() {
		token := instance.Token()
		return app.token(token, stackMap, escape, channels, isReverse, prevData, currentData)
	}

	everything := instance.Everything()
	return app.everything(everything, stackMap, isReverse, prevData, currentData)
}

func (app *application) everything(everything grammars.Everything, stackMap map[string]*stack, isReverse bool, prevData []byte, currentData []byte) (trees.Tree, map[string]*stack, error) {
	exception := everything.Exception()
	escape := everything.Escape()
	return app.token(exception, stackMap, escape, nil, !isReverse, prevData, currentData)
}

func (app *application) channels(channels grammars.Channels, prevData []byte, currentData []byte) (trees.Trees, []byte, error) {
	list := channels.List()
	treeList := []trees.Tree{}
	remaining := currentData
	previousData := prevData

	for {
		beginAmount := len(treeList)
		for _, oneChannel := range list {
			tree, err := app.channel(oneChannel, previousData, remaining)
			if err != nil {
				continue
			}

			if tree == nil {
				continue
			}

			prefixLength := len(tree.Bytes(true))
			rem := remaining[prefixLength:]
			if len(rem) == len(remaining) {
				continue
			}

			treeList = append(treeList, tree)
			previousData = remaining
			remaining = rem
		}

		if beginAmount == len(treeList) {
			break
		}
	}

	trees, err := app.treesBuilder.Create().WithList(treeList).Now()
	if err != nil {
		return nil, nil, err
	}

	return trees, remaining, nil
}

func (app *application) channel(channel grammars.Channel, prevData []byte, currentData []byte) (trees.Tree, error) {
	token := channel.Token()
	tree, _, err := app.token(token, map[string]*stack{}, nil, nil, false, prevData, currentData)
	if err != nil {
		return nil, err
	}

	if channel.HasCondition() {
		remaining := []byte{}
		if tree.HasRemaining() {
			remaining = tree.Remaining()
		}

		condition := channel.Condition()
		isAccepted, err := app.channelCondition(condition, prevData, remaining)
		if err != nil {
			return nil, err
		}

		if !isAccepted {
			return nil, nil
		}
	}

	return tree, nil
}

func (app *application) channelCondition(condition grammars.ChannelCondition, prevData []byte, nextData []byte) (bool, error) {
	isPrevMatch := true
	if condition.HasPrevious() {
		prevToken := condition.Previous()
		tree, _, err := app.token(prevToken, map[string]*stack{}, nil, nil, false, []byte{}, prevData)
		if err != nil {
			return false, err
		}

		isPrevMatch = tree != nil
	}

	isNextMatch := true
	if condition.HasNext() {
		nextToken := condition.Next()
		tree, _, err := app.token(nextToken, map[string]*stack{}, nil, nil, false, []byte{}, nextData)
		if err != nil {
			return false, err
		}

		isNextMatch = tree != nil
	}
	return isPrevMatch && isNextMatch, nil
}
