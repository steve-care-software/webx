package grammars

import (
	"errors"
	"fmt"

	database_applications "github.com/steve-care-software/webx/databases/applications"
	"github.com/steve-care-software/webx/databases/domain/cryptography/hash"
	grammar_applications "github.com/steve-care-software/webx/grammars/applications"
	"github.com/steve-care-software/webx/grammars/domain/grammars"
	"github.com/steve-care-software/webx/grammars/domain/grammars/cardinalities"
	"github.com/steve-care-software/webx/grammars/domain/grammars/values"
	contents_grammars "github.com/steve-care-software/webx/roots/domain/grammars"
	contents_channels "github.com/steve-care-software/webx/roots/domain/grammars/channels"
	contents_elements "github.com/steve-care-software/webx/roots/domain/grammars/elements"
	contents_everythings "github.com/steve-care-software/webx/roots/domain/grammars/everythings"
	contents_tokens "github.com/steve-care-software/webx/roots/domain/grammars/tokens"
)

type application struct {
	grammarApp                       grammar_applications.Application
	databaseApp                      database_applications.Application
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
	grammarValueBuilder              values.Builder
	grammarCardinalityBuilder        cardinalities.Builder
	hashAdapter                      hash.Adapter
}

func createApplication(
	grammarApp grammar_applications.Application,
	databaseApp database_applications.Application,
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
	grammarValueBuilder values.Builder,
	grammarCardinalityBuilder cardinalities.Builder,
	hashAdapter hash.Adapter,
) Application {
	out := application{
		grammarApp:                       grammarApp,
		databaseApp:                      databaseApp,
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
		grammarValueBuilder:              grammarValueBuilder,
		grammarCardinalityBuilder:        grammarCardinalityBuilder,
		hashAdapter:                      hashAdapter,
	}

	return &out
}

// Retrieve retrieves a grammar by hash
func (app *application) Retrieve(context uint, hash hash.Hash) (grammars.Grammar, error) {
	content, err := app.databaseApp.ReadByHash(context, hash)
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
	content, err := app.databaseApp.ReadByHash(context, hash)
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
	return app.grammarTokenBuilder.Create().
		WithName(name).
		WithBlock(block).
		Now()
}

func (app *application) retrieveElements(context uint, hashes []hash.Hash) ([]grammars.Element, error) {
	contents, err := app.databaseApp.ReadAllByHashes(context, hashes)
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
	content, err := app.databaseApp.ReadByHash(context, hash)
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
	content, err := app.databaseApp.ReadByHash(context, hash)
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
	contents, err := app.databaseApp.ReadAllByHashes(context, hashes)
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
	content, err := app.databaseApp.ReadByHash(context, hash)
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
	contentKeys, err := app.databaseApp.ContentKeysByKind(context, KindGrammar)
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
		coverages, err := app.grammarApp.Coverages(grammar)
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
	_, err := app.insertGrammar(context, grammar, map[string]hash.Hash{})
	if err != nil {
		return err
	}

	return nil
}

func (app *application) insertGrammar(context uint, grammar grammars.Grammar, recursives map[string]hash.Hash) (map[string]hash.Hash, error) {
	// if the grammar already exists:
	grammarHash := grammar.Hash()
	_, err := app.Retrieve(context, grammarHash)
	if err == nil {
		return recursives, nil
	}

	// root token:
	rootToken := grammar.Root()
	_, err = app.retrieveToken(context, rootToken.Hash())
	if err != nil {
		rec, err := app.insertToken(context, rootToken, recursives)
		if err != nil {
			return nil, err
		}

		recursives = rec
	}

	// channels:
	if grammar.HasChannels() {
		channels := grammar.Channels()
		rec, err := app.insertChannels(context, channels, recursives)
		if err != nil {
			return nil, err
		}

		recursives = rec
	}

	root := grammar.Root().Hash()
	channelsHashes := []hash.Hash{}
	if grammar.HasChannels() {
		channelsList := grammar.Channels().List()
		for _, oneChannel := range channelsList {
			channelsHashes = append(channelsHashes, oneChannel.Hash())
		}
	}

	builder := app.contentBuilder.Create().WithHash(grammarHash).WithRoot(root)
	if len(channelsHashes) > 0 {
		builder.WithChannels(channelsHashes)
	}

	content, err := builder.Now()
	if err != nil {
		return nil, err
	}

	grammarBytes, err := app.contentAdapter.ToContent(content)
	if err != nil {
		return nil, err
	}

	err = app.databaseApp.Write(context, grammarHash, grammarBytes, KindGrammar)
	if err != nil {
		return nil, err
	}

	return recursives, nil
}

func (app *application) insertToken(context uint, token grammars.Token, recursives map[string]hash.Hash) (map[string]hash.Hash, error) {
	// if the token already exists:
	tokenHash := token.Hash()
	_, err := app.retrieveToken(context, tokenHash)
	if err == nil {
		return recursives, nil
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
				rec, err := app.insertElement(context, oneElement, recursives)
				if err != nil {
					return nil, err
				}

				recursives = rec
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

	err = app.databaseApp.Write(context, tokenHash, tokenBytes, KindToken)
	if err != nil {
		return nil, err
	}

	return recursives, nil
}

func (app *application) insertElement(context uint, element grammars.Element, recursives map[string]hash.Hash) (map[string]hash.Hash, error) {
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

	elementHash := element.Hash()
	content := element.Content()
	builder := app.contentElementBuilder.Create().WithHash(elementHash).WithCardinality(contentCardinality)
	if content.IsValue() {
		number := content.Value().Number()
		builder.WithValue(number)
	}

	if content.IsExternal() {
		grammar := content.External().Grammar()
		grammarHash := grammar.Hash()
		_, err := app.Retrieve(context, grammarHash)
		if err != nil {
			rec, err := app.insertGrammar(context, grammar, recursives)
			if err != nil {
				return nil, err
			}

			recursives = rec
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
				rec, err := app.insertToken(context, token, recursives)
				if err != nil {
					return nil, err
				}

				recursives = rec
			}

			builder.WithToken(tokenHash)
		}

		if instance.IsEverything() {
			everything := instance.Everything()
			everythingHash := everything.Hash()
			_, err := app.retrieveEverything(context, everythingHash)
			if err != nil {
				rec, err := app.insertEverything(context, everything, recursives)
				if err != nil {
					return nil, err
				}

				recursives = rec
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

	err = app.databaseApp.Write(context, elementHash, elementBytes, KindElement)
	if err != nil {
		return nil, err
	}

	return recursives, nil
}

func (app *application) insertEverything(context uint, everything grammars.Everything, recursives map[string]hash.Hash) (map[string]hash.Hash, error) {
	// eception:
	exception := everything.Exception()
	exceptionHash := exception.Hash()
	_, err := app.retrieveToken(context, exceptionHash)
	if err != nil {
		rec, err := app.insertToken(context, exception, recursives)
		if err != nil {
			return nil, err
		}

		recursives = rec
	}

	// escape:
	everythingHash := everything.Hash()
	builder := app.contentEverythingBuilder.Create().WithHash(everythingHash).WithException(exceptionHash)
	if everything.HasEscape() {
		escape := everything.Escape()
		escapeHash := escape.Hash()
		_, err := app.retrieveToken(context, escapeHash)
		if err != nil {
			rec, err := app.insertToken(context, escape, recursives)
			if err != nil {
				return nil, err
			}

			recursives = rec
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

	err = app.databaseApp.Write(context, everythingHash, contentBytes, KindEverything)
	if err != nil {
		return nil, err
	}

	return recursives, nil
}

func (app *application) insertChannels(context uint, channels grammars.Channels, recursives map[string]hash.Hash) (map[string]hash.Hash, error) {
	channelsList := channels.List()
	for _, oneChannel := range channelsList {
		rec, err := app.insertChannel(context, oneChannel, recursives)
		if err != nil {
			return nil, err
		}

		recursives = rec
	}

	return recursives, nil
}

func (app *application) insertChannel(context uint, channel grammars.Channel, recursives map[string]hash.Hash) (map[string]hash.Hash, error) {
	token := channel.Token()
	tokenHash := token.Hash()
	_, err := app.retrieveToken(context, tokenHash)
	if err != nil {
		rec, err := app.insertToken(context, token, recursives)
		if err != nil {
			return nil, err
		}

		recursives = rec
	}

	channelHash := channel.Hash()
	builder := app.contentChannelBuilder.Create().WithHash(channelHash).WithToken(tokenHash)
	if channel.HasCondition() {
		condition := channel.Condition()
		if condition.HasPrevious() {
			previous := condition.Previous()
			previousHash := previous.Hash()
			_, err := app.retrieveToken(context, previousHash)
			if err != nil {
				rec, err := app.insertToken(context, previous, recursives)
				if err != nil {
					return nil, err
				}

				recursives = rec
			}

			builder.WithPrevious(previousHash)
		}

		if condition.HasNext() {
			next := condition.Next()
			nextHash := next.Hash()
			_, err := app.retrieveToken(context, nextHash)
			if err != nil {
				rec, err := app.insertToken(context, next, recursives)
				if err != nil {
					return nil, err
				}

				recursives = rec
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

	err = app.databaseApp.Write(context, channelHash, channelBytes, KindChannel)
	if err != nil {
		return nil, err
	}

	return recursives, nil
}
