package grammars

import "github.com/steve-care-software/logics/domain/bytes/grammars/cardinalities"

type element struct {
	name        string
	content     ElementContent
	cardinality cardinalities.Cardinality
}

func createElement(
	name string,
	content ElementContent,
	cardinality cardinalities.Cardinality,
) Element {
	out := element{
		name:        name,
		content:     content,
		cardinality: cardinality,
	}

	return &out
}

// Name returns the name
func (obj *element) Name() string {
	return obj.name
}

// Content returns the content
func (obj *element) Content() ElementContent {
	return obj.content
}

// Cardinality returns the cardinality
func (obj *element) Cardinality() cardinalities.Cardinality {
	return obj.cardinality
}
