package trees

import (
	"errors"
	"fmt"

	"github.com/steve-care-software/webx/domain/grammars"
)

type tree struct {
	grammar   grammars.Token
	block     Block
	suffix    Trees
	remaining []byte
}

func createTree(
	grammar grammars.Token,
	block Block,
) Tree {
	return createTreeInternally(grammar, block, nil, nil)
}

func createTreeWithRemaining(
	grammar grammars.Token,
	block Block,
	remaining []byte,
) Tree {
	return createTreeInternally(grammar, block, nil, remaining)
}

func createTreeWithSuffix(
	grammar grammars.Token,
	block Block,
	suffix Trees,
) Tree {
	return createTreeInternally(grammar, block, suffix, nil)
}

func createTreeWithSuffixAndRemaining(
	grammar grammars.Token,
	block Block,
	suffix Trees,
	remaining []byte,
) Tree {
	return createTreeInternally(grammar, block, suffix, remaining)
}

func createTreeInternally(
	grammar grammars.Token,
	block Block,
	suffix Trees,
	remaining []byte,
) Tree {
	out := tree{
		grammar:   grammar,
		block:     block,
		suffix:    suffix,
		remaining: remaining,
	}

	return &out
}

// Fetch fetches a tree or value by name
func (obj *tree) Fetch(name string, elementIndex uint) (Tree, Element, error) {
	if obj.Grammar().Name() == name {
		return obj, nil, nil
	}

	str := fmt.Sprintf("there is no Tree or Element associated to the given name: %s,at element's index: %d", name, elementIndex)
	if !obj.Block().HasSuccessful() {
		return nil, nil, errors.New(str)
	}

	cpt := uint(0)
	elementsList := obj.Block().Successful().Elements().List()
	for _, oneElement := range elementsList {
		tree, element, err := oneElement.Fetch(name, elementIndex)
		if err != nil {
			continue
		}

		isReady := cpt >= elementIndex
		if tree != nil && isReady {
			return tree, nil, nil
		}

		if element != nil && isReady {
			return nil, element, nil
		}

		if tree != nil || element != nil {
			cpt++
		}
	}

	return nil, nil, errors.New(str)
}

// Bytes returns the tree's bytes
func (obj *tree) Bytes(includeChannels bool) []byte {
	output := []byte{}
	if !obj.block.HasSuccessful() {
		return output
	}

	elements := obj.block.Successful().Elements().List()
	for _, oneElement := range elements {
		output = append(output, oneElement.Bytes(includeChannels)...)
	}

	if includeChannels && obj.HasSuffix() {
		output = append(output, obj.Suffix().Bytes(includeChannels)...)
	}

	return output
}

// Grammar returns the grammar
func (obj *tree) Grammar() grammars.Token {
	return obj.grammar
}

// Block returns the block
func (obj *tree) Block() Block {
	return obj.block
}

// HasSuffix returns true if there is suffix, false otherwise
func (obj *tree) HasSuffix() bool {
	return obj.suffix != nil
}

// Suffix returns the block
func (obj *tree) Suffix() Trees {
	return obj.suffix
}

// HasRemaining returns true if there is remaining, false otherwise
func (obj *tree) HasRemaining() bool {
	return obj.remaining != nil
}

// Remaining returns remaining, if any
func (obj *tree) Remaining() []byte {
	return obj.remaining
}
