package trees

import (
	"errors"
	"fmt"

	"github.com/steve-care-software/syntax/domain/bytes/grammars"
)

type element struct {
	grammar grammars.Element
	content Content
	amount  uint
}

func createElement(
	grammar grammars.Element,
	content Content,
	amount uint,
) Element {
	out := element{
		grammar: grammar,
		content: content,
		amount:  amount,
	}

	return &out
}

// Fetch fetches a tree or value by name
func (obj *element) Fetch(name string, elementIndex uint) (Tree, Element, error) {
	if obj.Grammar().Name() == name {
		return nil, obj, nil
	}

	if obj.content.IsTree() {
		return obj.content.Tree().Fetch(name, elementIndex)
	}

	if obj.content.IsValue() {
		valueName := obj.content.Value().Content().Name()
		if valueName == name {
			return nil, obj, nil
		}
	}

	str := fmt.Sprintf("there is no Tree or Element associated to the given name: %s", name)
	return nil, nil, errors.New(str)
}

// Bytes returns the element's bytes
func (obj *element) Bytes(includeChannels bool) []byte {
	output := []byte{}
	castedAmount := int(obj.amount)
	for i := 0; i < castedAmount; i++ {
		if obj.content.IsValue() {
			value := obj.content.Value()
			if includeChannels && value.HasPrefix() {
				output = append(output, value.Prefix().Bytes(includeChannels)...)
			}

			output = append(output, value.Content().Number())
			continue
		}
	}
	return nil
}

// Grammar returns the grammar
func (obj *element) Grammar() grammars.Element {
	return obj.grammar
}

// Content returns the content
func (obj *element) Content() Content {
	return obj.content
}

// Amount returns the amount
func (obj *element) Amount() uint {
	return obj.amount
}
