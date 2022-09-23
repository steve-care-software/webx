package trees

import "github.com/steve-care-software/syntax/domain/bytes/grammars"

type line struct {
	grammar  grammars.Line
	elements Elements
}

func createLine(
	grammar grammars.Line,
	elements Elements,
) Line {
	out := line{
		grammar:  grammar,
		elements: elements,
	}

	return &out
}

// Grammar returns the grammar
func (obj *line) Grammar() grammars.Line {
	return obj.grammar
}

// Elements returns the elements
func (obj *line) Elements() Elements {
	return obj.elements
}

// IsSuccessful returns true if successful, false otherwise
func (obj *line) IsSuccessful() bool {
	requested := obj.grammar.Elements()
	elementsWithoutChannels := obj.elements.List(false)
	return len(requested) == len(elementsWithoutChannels)
}
