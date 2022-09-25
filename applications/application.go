package applications

import (
	"errors"
	"fmt"
	"strings"

	"github.com/steve-care-software/syntax/domain/bytes/criterias"
	"github.com/steve-care-software/syntax/domain/bytes/grammars"
	"github.com/steve-care-software/syntax/domain/bytes/trees"
	"github.com/steve-care-software/syntax/domain/compilers"
	"github.com/steve-care-software/syntax/domain/outputs"
	"github.com/steve-care-software/syntax/domain/programs"
)

type application struct {
	grammarTokenBuilder grammars.TokenBuilder
	treesBuilder        trees.Builder
	treeBuilder         trees.TreeBuilder
	treeBlockBuilder    trees.BlockBuilder
	treeLineBuilder     trees.LineBuilder
	treeElementsBuilder trees.ElementsBuilder
	treeElementBuilder  trees.ElementBuilder
	treeContentBuilder  trees.ContentBuilder
	treeValueBuilder    trees.ValueBuilder
}

func createApplication(
	grammarTokenBuilder grammars.TokenBuilder,
	treesBuilder trees.Builder,
	treeBuilder trees.TreeBuilder,
	treeBlockBuilder trees.BlockBuilder,
	treeLineBuilder trees.LineBuilder,
	treeElementsBuilder trees.ElementsBuilder,
	treeElementBuilder trees.ElementBuilder,
	treeContentBuilder trees.ContentBuilder,
	treeValueBuilder trees.ValueBuilder,
) Application {
	out := application{
		grammarTokenBuilder: grammarTokenBuilder,
		treesBuilder:        treesBuilder,
		treeBuilder:         treeBuilder,
		treeBlockBuilder:    treeBlockBuilder,
		treeLineBuilder:     treeLineBuilder,
		treeElementsBuilder: treeElementsBuilder,
		treeElementBuilder:  treeElementBuilder,
		treeContentBuilder:  treeContentBuilder,
		treeValueBuilder:    treeValueBuilder,
	}

	return &out
}

// Tokenize tokenize the values to a tree using the provided grammar
func (app *application) Tokenize(grammar grammars.Grammar, values []byte) (trees.Tree, []byte, error) {
	return app.grammar(grammar, false, []byte{}, values)
}

func (app *application) grammar(grammar grammars.Grammar, isInChannel bool, prevData []byte, currentData []byte) (trees.Tree, []byte, error) {
	root := grammar.Root()
	channels := grammar.Channels()
	return app.token(root, channels, isInChannel, prevData, currentData)
}

func (app *application) token(token grammars.Token, channels grammars.Channels, isInChannel bool, prevData []byte, currentData []byte) (trees.Tree, []byte, error) {
	tokenBlock := token.Block()
	block, remaining, err := app.block(tokenBlock, channels, isInChannel, prevData, currentData)
	if err != nil {
		return nil, nil, err
	}

	builder := app.treeBuilder.Create().WithGrammar(token).WithBlock(block)
	if !isInChannel {
		suffix, rem, err := app.channels(channels, prevData, remaining)
		if err != nil {
			return nil, nil, err
		}

		builder.WithSuffix(suffix)
		remaining = rem
	}

	ins, err := builder.Now()
	if err != nil {
		return nil, nil, err
	}

	return ins, remaining, nil
}

func (app *application) external(external grammars.External, isInChannel bool, prevData []byte, currentData []byte) (trees.Tree, []byte, error) {
	grammar := external.Grammar()
	treeIns, remaining, err := app.grammar(grammar, isInChannel, prevData, currentData)
	if err != nil {
		return nil, nil, err
	}

	name := external.Name()
	root := grammar.Root()
	block := root.Block()
	builder := app.grammarTokenBuilder.Create().WithName(name).WithBlock(block)
	if root.HasSuites() {
		suites := root.Suites()
		builder.WithSuites(suites)
	}

	grammarRoot, err := builder.Now()
	if err != nil {
		return nil, nil, err
	}

	treeBlock := treeIns.Block()
	ins, err := app.treeBuilder.Create().WithGrammar(grammarRoot).WithBlock(treeBlock).Now()
	if err != nil {
		return nil, nil, err
	}

	return ins, remaining, nil
}

func (app *application) block(block grammars.Block, channels grammars.Channels, isInChannel bool, prevData []byte, currentData []byte) (trees.Block, []byte, error) {
	list := []trees.Line{}
	lines := block.Lines()
	remaining := currentData
	for _, oneLine := range lines {
		lineIns, rem, err := app.line(oneLine, channels, isInChannel, prevData, remaining)
		if err != nil {
			return nil, nil, err
		}

		list = append(list, lineIns)
		if lineIns.IsSuccessful() {
			remaining = rem
			break
		}
	}

	blockIns, err := app.treeBlockBuilder.Create().WithLines(list).Now()
	if err != nil {
		return nil, nil, err
	}

	return blockIns, remaining, nil
}

func (app *application) line(line grammars.Line, channels grammars.Channels, isInChannel bool, prevData []byte, currentData []byte) (trees.Line, []byte, error) {
	list := []trees.Element{}
	grElements := line.Elements()
	remaining := currentData
	previousData := prevData
	for _, oneElement := range grElements {
		element, rem, err := app.element(oneElement, channels, isInChannel, previousData, remaining)
		if err != nil {
			break
		}

		list = append(list, element)
		previousData = remaining
		remaining = rem
	}

	elements, err := app.treeElementsBuilder.Create().WithList(list).Now()
	if err != nil {
		return nil, nil, err
	}

	lineIns, err := app.treeLineBuilder.Create().
		WithGrammar(line).
		WithElements(elements).
		Now()

	if err != nil {
		return nil, nil, err
	}

	return lineIns, remaining, nil
}

func (app *application) element(element grammars.Element, channels grammars.Channels, isInChannel bool, prevData []byte, currentData []byte) (trees.Element, []byte, error) {
	content := element.Content()
	cardinality := element.Cardinality()
	min := cardinality.Min()
	pMax := cardinality.Max()

	cpt := uint(0)
	remaining := currentData
	previousData := prevData
	var value trees.Value
	var tree trees.Tree
	for {
		if cpt >= *pMax {
			break
		}

		val, tr, rem, err := app.elementContent(content, channels, isInChannel, previousData, remaining)
		if err != nil {
			return nil, nil, err
		}

		if val == nil && tr == nil {
			break
		}

		if val != nil {
			value = val
		}

		if tr != nil {
			tree = tr
		}

		previousData = remaining
		remaining = rem
		cpt++
	}

	if cpt < min {
		str := fmt.Sprintf("the cardinality's minimum is %d, %d elements found", min, cpt)
		return nil, nil, errors.New(str)
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

	elementIns, err := app.treeElementBuilder.Create().
		WithGrammar(element).
		WithContent(contentIns).
		WithAmount(cpt).
		Now()

	if err != nil {
		return nil, nil, err
	}

	return elementIns, remaining, nil
}

func (app *application) elementContent(content grammars.ElementContent, channels grammars.Channels, isInChannel bool, prevData []byte, currentData []byte) (trees.Value, trees.Tree, []byte, error) {
	if content.IsExternal() {
		external := content.External()
		tree, remaining, err := app.external(external, isInChannel, prevData, currentData)
		if err != nil {
			return nil, nil, nil, err
		}

		return nil, tree, remaining, nil
	}

	if content.IsToken() {
		token := content.Token()
		tree, remaining, err := app.token(token, channels, isInChannel, prevData, currentData)
		if err != nil {
			return nil, nil, nil, err
		}

		return nil, tree, remaining, nil
	}

	if len(currentData) <= 1 {
		return nil, nil, nil, errors.New("there must be at least 1 value in the given data in order to have an element match, 0 provided")
	}

	remaining := currentData
	builder := app.treeValueBuilder.Create()
	if !isInChannel {
		prefix, rem, err := app.channels(channels, prevData, remaining)
		if err != nil {
			return nil, nil, nil, err
		}

		builder.WithPrefix(prefix)
		remaining = rem
	}

	value := content.Value()
	number := value.Number()
	if number == remaining[0] {
		ins, err := builder.WithContent(value).Now()
		if err != nil {
			return nil, nil, nil, err
		}

		return ins, nil, remaining[1:], nil
	}

	return nil, nil, nil, nil
}

func (app *application) channels(channels grammars.Channels, prevData []byte, currentData []byte) (trees.Trees, []byte, error) {
	list := channels.List()
	treeList := []trees.Tree{}
	remaining := currentData
	previousData := prevData

	for {
		beginAmount := len(treeList)
		for _, oneChannel := range list {
			tree, rem, err := app.channel(oneChannel, previousData, remaining)
			if err != nil {
				return nil, nil, err
			}

			if tree == nil {
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

func (app *application) channel(channel grammars.Channel, prevData []byte, currentData []byte) (trees.Tree, []byte, error) {
	token := channel.Token()
	tree, remaining, err := app.token(token, nil, true, prevData, currentData)
	if err != nil {
		return nil, nil, err
	}

	if channel.HasCondition() {
		condition := channel.Condition()
		isAccepted, err := app.channelCondition(condition, prevData, remaining)
		if err != nil {
			return nil, nil, err
		}

		if !isAccepted {
			return nil, nil, nil
		}
	}

	return tree, remaining, nil
}

func (app *application) channelCondition(condition grammars.ChannelCondition, prevData []byte, nextData []byte) (bool, error) {
	isPrevMatch := true
	if condition.HasPrevious() {
		prevToken := condition.Previous()
		tree, _, err := app.token(prevToken, nil, true, []byte{}, prevData)
		if err != nil {
			return false, err
		}

		isPrevMatch = tree != nil
	}

	isNextMatch := true
	if condition.HasNext() {
		nextToken := condition.Next()
		tree, _, err := app.token(nextToken, nil, true, []byte{}, nextData)
		if err != nil {
			return false, err
		}

		isNextMatch = tree != nil
	}
	return isPrevMatch && isNextMatch, nil
}

// Extract extracts data from a tree using the provided criteria
func (app *application) Extract(criteria criterias.Criteria, tree trees.Tree) ([]byte, error) {
	return app.extractWithPath([]string{}, criteria, tree)
}

func (app *application) extractWithPath(path []string, criteria criterias.Criteria, tree trees.Tree) ([]byte, error) {
	name := criteria.Name()
	index := criteria.Index()
	subTree, element, err := tree.Fetch(name, index)
	if err != nil {
		return nil, err
	}

	includeChannels := criteria.IncludeChannels()
	if subTree != nil {
		if criteria.HasChild() {
			child := criteria.Child()
			return app.extractWithPath(append(path, name), child, subTree)
		}

		return subTree.Bytes(includeChannels), nil
	}

	if criteria.HasChild() {
		str := fmt.Sprintf("the extraction did NOT succeed because it found an element (path: %s) but the criteria had a child", strings.Join(append(path, name), "/"))
		return nil, errors.New(str)
	}

	return element.Bytes(includeChannels), nil
}

// Combine combines the data of trees
func (app *application) Combine(trees []trees.Tree, includeChannels bool) ([]byte, error) {
	output := []byte{}
	for _, oneTree := range trees {
		output = append(output, oneTree.Bytes(includeChannels)...)
	}

	return output, nil
}

// Compile compiles a script into a program using the provided compiler
func (app *application) Compile(compiler compilers.Compiler, script []byte) (programs.Program, error) {
	return nil, nil
}

// Execute executes a program using the provided input parameters and returns its output
func (app *application) Execute(input map[string]interface{}, program programs.Program) (outputs.Output, error) {
	return nil, nil
}
