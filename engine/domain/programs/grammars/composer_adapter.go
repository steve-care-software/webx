package grammars

import (
	"errors"
	"fmt"

	"github.com/steve-care-software/webx/engine/domain/programs/grammars/blocks"
	"github.com/steve-care-software/webx/engine/domain/programs/grammars/blocks/lines"
	"github.com/steve-care-software/webx/engine/domain/programs/grammars/blocks/lines/processors"
	"github.com/steve-care-software/webx/engine/domain/programs/grammars/blocks/lines/tokens"
	"github.com/steve-care-software/webx/engine/domain/programs/grammars/blocks/lines/tokens/elements"
)

type composeAdapter struct {
	funcsMap map[string]CoreFn
}

func createComposeAdapter(
	funcsMap map[string]CoreFn,
) ComposeAdapter {
	out := composeAdapter{
		funcsMap: funcsMap,
	}

	return &out
}

// ToBytes takes a grammar and a blockname and returns its bytes
func (app *composeAdapter) ToBytes(grammar Grammar, blockName string) ([]byte, error) {
	block, err := grammar.Blocks().Fetch(blockName)
	if err != nil {
		return nil, err
	}

	return app.writeBlock(grammar, block)
}

func (app *composeAdapter) writeBlock(grammar Grammar, block blocks.Block) ([]byte, error) {
	if !block.HasLine() {
		str := fmt.Sprintf("the block (name: %s) cannot be written because it contains lines instead of a line", block.Name())
		return nil, errors.New(str)
	}

	line := block.Line()
	return app.writeLine(grammar, line)
}

func (app *composeAdapter) writeLine(grammar Grammar, line lines.Line) ([]byte, error) {
	tokens := line.Tokens()
	retBytes, tokensMap, err := app.writeTokens(grammar, tokens)
	if err != nil {
		return nil, err
	}

	output := retBytes
	if line.HasProcessor() {
		processor := line.Processor()
		retBytes, err := app.writeProcessor(grammar, tokensMap, processor)
		if err != nil {
			return nil, err
		}

		output = retBytes
	}

	return output, nil
}

func (app *composeAdapter) writeProcessor(
	grammar Grammar,
	tokensMap map[string][][]byte,
	processor processors.Processor,
) ([]byte, error) {
	if processor.IsExecution() {
		execution := processor.Execution()
		fnName := execution.FuncName()
		params := map[string][]byte{}
		if execution.HasParameters() {
			parametersList := execution.Parameters().List()
			for _, oneParameter := range parametersList {
				name := oneParameter.Name()
				value := oneParameter.Value()
				if value.IsReference() {
					reference := value.Reference()
					paramElementName := reference.Element().Name()
					paramElementIndex := reference.Index()
					if _, ok := tokensMap[paramElementName]; !ok {
						str := fmt.Sprintf("the func (name: %s) contains a param (name: %s, index: %d) that is not declared in the line", fnName, paramElementName, paramElementIndex)
						return nil, errors.New(str)
					}

					params[name] = tokensMap[paramElementName][int(paramElementIndex)]
					continue
				}

				params[name] = value.Bytes()
			}
		}

		if fn, ok := app.funcsMap[fnName]; ok {
			return fn(params)
		}
	}

	replacement := processor.Replacement()
	return app.writeElement(grammar, replacement)
}

func (app *composeAdapter) writeTokens(grammar Grammar, tokens tokens.Tokens) ([]byte, map[string][][]byte, error) {
	output := []byte{}
	mp := map[string][][]byte{}
	list := tokens.List()
	for _, oneToken := range list {
		name := oneToken.Name()
		retBytes, retMultiLine, err := app.writeToken(grammar, oneToken)
		if err != nil {
			return nil, nil, err
		}

		mp[name] = retMultiLine
		output = append(output, retBytes...)
	}

	return output, mp, nil
}

func (app *composeAdapter) writeToken(grammar Grammar, token tokens.Token) ([]byte, [][]byte, error) {
	name := token.Name()
	cardinality := token.Cardinality()
	if !cardinality.HasMax() {
		str := fmt.Sprintf("the cardinality, in the token (name: %s) must contain a max in order to be written", name)
		return nil, nil, errors.New(str)
	}

	pMax := cardinality.Max()
	min := cardinality.Min()
	if *pMax != min {
		str := fmt.Sprintf("the cardinality, in the token (name: %s) must contain a min (%d) and a max (%d) that are equal in order to be written", name, min, *pMax)
		return nil, nil, errors.New(str)
	}

	element := token.Element()
	elementBytes, err := app.writeElement(grammar, element)
	if err != nil {
		return nil, nil, err
	}

	output := []byte{}
	multiLine := [][]byte{}
	castedMin := int(min)
	for i := 0; i < castedMin; i++ {
		multiLine = append(multiLine, elementBytes)
		output = append(output, elementBytes...)
	}

	return output, multiLine, nil
}

func (app *composeAdapter) writeElement(grammar Grammar, element elements.Element) ([]byte, error) {
	if element.IsBlock() {
		blockName := element.Block()
		block, err := grammar.Blocks().Fetch(blockName)
		if err != nil {
			return nil, err
		}

		return app.writeBlock(grammar, block)
	}

	ruleName := element.Rule()
	rule, err := grammar.Rules().Fetch(ruleName)
	if err != nil {
		return nil, err
	}

	return rule.Bytes(), nil
}
