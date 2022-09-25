package compilers

import "github.com/steve-care-software/syntax/domain/bytes/grammars"

type element struct {
	grammar     grammars.Grammar
	composition Composition
}

func createElement(
	grammar grammars.Grammar,
	composition Composition,
) Element {
	out := element{
		grammar:     grammar,
		composition: composition,
	}

	return &out
}

// Grammar returns the grammar
func (obj *element) Grammar() grammars.Grammar {
	return obj.grammar
}

// Composition returns the composition
func (obj *element) Composition() Composition {
	return obj.composition
}
