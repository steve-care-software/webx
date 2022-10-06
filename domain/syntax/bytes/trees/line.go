package trees

import "github.com/steve-care-software/syntax/domain/syntax/bytes/grammars"

type line struct {
	index    uint
	grammar  grammars.Line
	elements Elements
}

func createLine(
	index uint,
	grammar grammars.Line,
) Line {
	return createLineInternally(index, grammar, nil)
}

func createLineWithElements(
	index uint,
	grammar grammars.Line,
	elements Elements,
) Line {
	return createLineInternally(index, grammar, elements)
}

func createLineInternally(
	index uint,
	grammar grammars.Line,
	elements Elements,
) Line {
	out := line{
		index:    index,
		grammar:  grammar,
		elements: elements,
	}

	return &out
}

// Index returns the index
func (obj *line) Index() uint {
	return obj.index
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
	return len(requested) == len(elements)
}
