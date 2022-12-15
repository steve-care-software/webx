package trees

import (
	"github.com/steve-care-software/webx/roots/domain/grammars/grammars"
)

type line struct {
	index     uint
	grammar   grammars.Line
	isReverse bool
	elements  Elements
}

func createLine(
	index uint,
	grammar grammars.Line,
	isReverse bool,
) Line {
	return createLineInternally(index, grammar, isReverse, nil)
}

func createLineWithElements(
	index uint,
	grammar grammars.Line,
	isReverse bool,
	elements Elements,
) Line {
	return createLineInternally(index, grammar, isReverse, elements)
}

func createLineInternally(
	index uint,
	grammar grammars.Line,
	isReverse bool,
	elements Elements,
) Line {
	out := line{
		index:     index,
		grammar:   grammar,
		isReverse: isReverse,
		elements:  elements,
	}

	return &out
}

// Index returns the index
func (obj *line) Index() uint {
	return obj.index
}

// IsReverse returns true if reverse, false otherwise
func (obj *line) IsReverse() bool {
	return obj.isReverse
}

// Grammar returns the grammar
func (obj *line) Grammar() grammars.Line {
	return obj.grammar
}

// HasElements returns true if there is elements, false otherwise
func (obj *line) HasElements() bool {
	return obj.elements != nil
}

// Elements returns the elements
func (obj *line) Elements() Elements {
	return obj.elements
}

// IsSuccessful returns true if successful, false otherwise
func (obj *line) IsSuccessful() bool {
	if !obj.HasElements() {
		return false
	}

	requested := obj.grammar.Elements()
	elements := obj.elements.List()
	for _, oneElement := range elements {
		if !oneElement.IsSuccessful() {
			return false
		}
	}

	if obj.IsReverse() {
		return true
	}

	return len(requested) == len(elements)
}
