package trees

import (
	"errors"
	"fmt"

	"github.com/steve-care-software/syntax/domain/syntax/bytes/grammars"
)

type element struct {
	grammar  grammars.Element
	contents Contents
}

func createElement(
	grammar grammars.Element,
	contents Contents,
) Element {
	out := element{
		grammar:  grammar,
		contents: contents,
	}

	return &out
}

// Fetch fetches a tree or value by name
func (obj *element) Fetch(name string, elementIndex uint) (Tree, Element, error) {
	if obj.grammar.Name() == name {
		return nil, obj, nil
	}

	list := obj.contents.List()
	for _, oneContent := range list {
		if !oneContent.IsTree() {
			continue
		}

		tree, element, err := oneContent.Tree().Fetch(name, elementIndex)
		if err != nil {
			continue
		}

		if tree != nil {
			return tree, nil, nil
		}

		if element != nil {
			return nil, element, nil
		}
	}

	str := fmt.Sprintf("there is no Tree or Element associated to the given name: %s", name)
	return nil, nil, errors.New(str)
}

// Bytes returns the element's bytes
func (obj *element) Bytes(includeChannels bool) []byte {
	output := []byte{}
	list := obj.contents.List()
	for _, oneContent := range list {
		if oneContent.IsValue() {
			value := oneContent.Value()
			if includeChannels && value.HasPrefix() {
				output = append(output, value.Prefix().Bytes(includeChannels)...)
			}

			output = append(output, value.Content())
			continue
		}

		output = append(output, oneContent.Tree().Bytes(includeChannels)...)
	}

	return output
}

// Grammar returns the grammar
func (obj *element) Grammar() grammars.Element {
	return obj.grammar
}

// Contents returns the contents
func (obj *element) Contents() Contents {
	return obj.contents
}

// Amount returns the amount
func (obj *element) Amount() uint {
	return uint(len(obj.contents.List()))
}
