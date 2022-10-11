package grammars

import (
	"errors"
	"fmt"

	"github.com/steve-care-software/syntax/domain/syntax/bytes/grammars"
	"github.com/steve-care-software/syntax/domain/syntax/bytes/grammars/coverages"
	"github.com/steve-care-software/syntax/domain/syntax/bytes/trees"
)

type application struct {
	grammarTokenBuilder       grammars.TokenBuilder
	treesBuilder              trees.Builder
	treeBuilder               trees.TreeBuilder
	treeBlockBuilder          trees.BlockBuilder
	treeLineBuilder           trees.LineBuilder
	treeElementsBuilder       trees.ElementsBuilder
	treeElementBuilder        trees.ElementBuilder
	treeContentsBuilder       trees.ContentsBuilder
	treeContentBuilder        trees.ContentBuilder
	treeValueBuilder          trees.ValueBuilder
	coveragesBuilder          coverages.Builder
	coverageBuilder           coverages.CoverageBuilder
	coverageExecutionsBuilder coverages.ExecutionsBuilder
	coverageExecutionBuilder  coverages.ExecutionBuilder
	coverageResultBuilder     coverages.ResultBuilder
}

func createApplication(
	grammarTokenBuilder grammars.TokenBuilder,
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
) Application {
	out := application{
		grammarTokenBuilder:       grammarTokenBuilder,
		treesBuilder:              treesBuilder,
		treeBuilder:               treeBuilder,
		treeBlockBuilder:          treeBlockBuilder,
		treeLineBuilder:           treeLineBuilder,
		treeElementsBuilder:       treeElementsBuilder,
		treeElementBuilder:        treeElementBuilder,
		treeContentsBuilder:       treeContentsBuilder,
		treeContentBuilder:        treeContentBuilder,
		treeValueBuilder:          treeValueBuilder,
		coveragesBuilder:          coveragesBuilder,
		coverageBuilder:           coverageBuilder,
		coverageExecutionsBuilder: coverageExecutionsBuilder,
		coverageExecutionBuilder:  coverageExecutionBuilder,
		coverageResultBuilder:     coverageResultBuilder,
	}

	return &out
}

// Execute executes the application
func (app *application) Execute(grammar grammars.Grammar, values []byte) (trees.Tree, error) {
	return app.grammar(grammar, false, []byte{}, values)
}

// Coverages returns the coverages
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
	tree, err := app.token(token, nil, channels, false, []byte{}, input)
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

// Covered returns the covered report
func (app *application) Covered(coverages coverages.Coverages) (map[string]map[uint]map[uint]string, error) {
	coveredElements := map[string]map[uint]map[uint]string{}
	err := app.findCoveraredElements(coverages, &coveredElements)
	if err != nil {
		return nil, err
	}

	return coveredElements, nil
}

// Uncovered returns the uncovered report
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
	return app.token(root, nil, channels, isReverse, prevData, currentData)
}

func (app *application) token(token grammars.Token, escape grammars.Token, channels grammars.Channels, isReverse bool, prevData []byte, currentData []byte) (trees.Tree, error) {
	tokenName := token.Name()
	tokenBlock := token.Block()
	block, remaining, err := app.block(tokenName, tokenBlock, escape, channels, isReverse, prevData, currentData)
	if err != nil {
		return nil, err
	}

	if block == nil {
		str := fmt.Sprintf("there was no line discovered in the token (name: %s) using the given data: %v", token.Name(), currentData)
		return nil, errors.New(str)
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

	return builder.Now()
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

func (app *application) block(tokenName string, block grammars.Block, escape grammars.Token, channels grammars.Channels, isReverse bool, prevData []byte, currentData []byte) (trees.Block, []byte, error) {
	list := []trees.Line{}
	lines := block.Lines()
	remaining := currentData
	for idx, oneLine := range lines {
		if isReverse {
			previousData := prevData
			contentsList := []trees.Content{}
			for {
				if len(remaining) <= 0 {
					break
				}

				if escape != nil {
					fmt.Printf("\n finish escape!!! \n ")
					//escapeIns, rem, err:= app.element(oneElement, escape, channels, isReverse, previousData, remaining)
				}

				_, _, err := app.line(tokenName, oneLine, uint(idx), escape, channels, isReverse, previousData, remaining)
				if err == nil {
					break
				}

				value, err := app.treeValueBuilder.Create().WithContent(remaining[0]).Now()
				if err != nil {
					return nil, nil, err
				}

				contentIns, err := app.treeContentBuilder.Create().WithValue(value).Now()
				if err != nil {
					return nil, nil, err
				}

				contentsList = append(contentsList, contentIns)
				previousData = remaining
				remaining = remaining[1:]
			}

			contents, err := app.treeContentsBuilder.Create().WithList(contentsList).Now()
			if err != nil {
				return nil, nil, err
			}

			elementIns, err := app.treeElementBuilder.Create().WithContents(contents).Now()
			if err != nil {
				return nil, nil, err
			}

			elements, err := app.treeElementsBuilder.Create().WithList([]trees.Element{
				elementIns,
			}).Now()
			if err != nil {
				return nil, nil, err
			}

			lineIns, err := app.treeLineBuilder.Create().
				WithIndex(uint(idx)).
				WithGrammar(oneLine).
				WithElements(elements).
				IsReverse().
				Now()

			if err != nil {
				return nil, nil, err
			}

			list = append(list, lineIns)
			break
		}

		lineIns, rem, err := app.line(tokenName, oneLine, uint(idx), escape, channels, isReverse, prevData, remaining)
		if err != nil {
			continue
		}

		list = append(list, lineIns)
		if lineIns.IsSuccessful() {
			remaining = rem
			break
		}
	}

	if len(list) <= 0 {
		return nil, remaining, nil
	}

	blockIns, err := app.treeBlockBuilder.Create().WithLines(list).Now()
	if err != nil {
		return nil, nil, err
	}

	return blockIns, remaining, nil
}

func (app *application) line(tokenName string, line grammars.Line, index uint, escape grammars.Token, channels grammars.Channels, isReverse bool, prevData []byte, currentData []byte) (trees.Line, []byte, error) {
	list := []trees.Element{}
	grElements := line.Elements()
	remaining := currentData
	previousData := prevData
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

			contentIns, rem, err := app.element(oneElement, escape, channels, isReverse, previousData, remaining)
			if err != nil {
				break
			}

			contentsList = append(contentsList, contentIns)
			previousData = remaining
			remaining = rem
		}

		contents, err := app.treeContentsBuilder.Create().WithList(contentsList).Now()
		if err != nil {
			return nil, nil, err
		}

		elementIns, err := app.treeElementBuilder.Create().WithGrammar(oneElement).WithContents(contents).Now()
		if err != nil {
			return nil, nil, err
		}

		list = append(list, elementIns)
	}

	builder := app.treeLineBuilder.Create().
		WithIndex(index).
		WithGrammar(line)

	if len(list) > 0 {
		elements, err := app.treeElementsBuilder.Create().WithList(list).Now()
		if err != nil {
			return nil, nil, err
		}

		builder.WithElements(elements)
	}

	lineIns, err := builder.Now()
	if err != nil {
		return nil, nil, err
	}

	return lineIns, remaining, nil
}

func (app *application) element(element grammars.Element, escape grammars.Token, channels grammars.Channels, isReverse bool, prevData []byte, currentData []byte) (trees.Content, []byte, error) {
	if len(currentData) <= 0 {
		return nil, nil, errors.New("no remaining data")
	}

	content := element.Content()
	value, tree, rem, err := app.elementContent(content, escape, channels, isReverse, prevData, currentData)
	if err != nil {
		return nil, nil, err
	}

	if value == nil && tree == nil {
		return nil, nil, errors.New("no value/tree found")
	}

	if tree != nil && !tree.Block().HasSuccessful() {
		return nil, nil, errors.New("no successfull tree found")
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
		return nil, nil, err
	}

	return contentIns, rem, nil
}

func (app *application) elementContent(content grammars.ElementContent, escape grammars.Token, channels grammars.Channels, isReverse bool, prevData []byte, currentData []byte) (trees.Value, trees.Tree, []byte, error) {
	if content.IsExternal() {
		external := content.External()
		tree, err := app.external(external, isReverse, prevData, currentData)
		if err != nil {
			return nil, nil, nil, err
		}

		if tree == nil {
			return nil, nil, nil, nil
		}

		remaining := []byte{}
		if tree.HasRemaining() {
			remaining = tree.Remaining()
		}

		return nil, tree, remaining, nil
	}

	if content.IsInstance() {
		instance := content.Instance()
		tree, err := app.instance(instance, escape, channels, isReverse, prevData, currentData)
		if err != nil {
			return nil, nil, nil, err
		}

		remaining := []byte{}
		if tree.HasRemaining() {
			remaining = tree.Remaining()
		}

		return nil, tree, remaining, nil
	}

	if len(currentData) < 1 {
		return nil, nil, nil, errors.New("there must be at least 1 value in the given data in order to have an element match, 0 provided")
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

	number := content.Value().Number()
	if number == remaining[0] {
		ins, err := builder.WithContent(remaining[0]).Now()
		if err != nil {
			return nil, nil, nil, err
		}

		return ins, nil, remaining[1:], nil
	}

	return nil, nil, nil, nil
}

func (app *application) instance(instance grammars.Instance, escape grammars.Token, channels grammars.Channels, isReverse bool, prevData []byte, currentData []byte) (trees.Tree, error) {
	if instance.IsToken() {
		token := instance.Token()
		return app.token(token, escape, channels, isReverse, prevData, currentData)
	}

	everything := instance.Everything()
	return app.everything(everything, isReverse, prevData, currentData)
}

func (app *application) everything(everything grammars.Everything, isReverse bool, prevData []byte, currentData []byte) (trees.Tree, error) {
	exception := everything.Exception()
	escape := everything.Escape()
	return app.token(exception, escape, nil, !isReverse, prevData, currentData)
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
	tree, err := app.token(token, nil, nil, false, prevData, currentData)
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
		tree, err := app.token(prevToken, nil, nil, false, []byte{}, prevData)
		if err != nil {
			return false, err
		}

		isPrevMatch = tree != nil
	}

	isNextMatch := true
	if condition.HasNext() {
		nextToken := condition.Next()
		tree, err := app.token(nextToken, nil, nil, false, []byte{}, nextData)
		if err != nil {
			return false, err
		}

		isNextMatch = tree != nil
	}
	return isPrevMatch && isNextMatch, nil
}
