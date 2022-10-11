package grammars

import "github.com/steve-care-software/syntax/domain/syntax/grammars/cardinalities"

type element struct {
	content     ElementContent
	cardinality cardinalities.Cardinality
}

func createElement(
	content ElementContent,
	cardinality cardinalities.Cardinality,
) Element {
	out := element{
		content:     content,
		cardinality: cardinality,
	}

	return &out
}

// Name returns the name
func (obj *element) Name() string {
	if obj.content.IsValue() {
		return obj.content.Value().Name()
	}

	if obj.content.IsExternal() {
		return obj.content.External().Name()
	}

	return obj.content.Instance().Name()
}

// Content returns the content
func (obj *element) Content() ElementContent {
	return obj.content
}

// Cardinality returns the cardinality
func (obj *element) Cardinality() cardinalities.Cardinality {
	return obj.cardinality
}
