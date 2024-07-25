package routes

import (
	"github.com/steve-care-software/webx/engine/states/domain/hash"
	"github.com/steve-care-software/webx/engine/vms/domain/instances/routes/cardinalities"
)

// NewRouteWithGobalAndTokenForTests creates a new route with global and token omission for tests
func NewRouteWithGobalAndTokenForTests(tokens Tokens, global Omission, token Omission) Route {
	ins, err := NewBuilder().Create().WithTokens(tokens).WithGlobal(global).WithToken(token).Now()
	if err != nil {
		panic(err)
	}

	return ins
}

// NewRouteWithGobalForTests creates a new route with global omission for tests
func NewRouteWithGobalForTests(tokens Tokens, global Omission) Route {
	ins, err := NewBuilder().Create().WithTokens(tokens).WithGlobal(global).Now()
	if err != nil {
		panic(err)
	}

	return ins
}

// NewRouteWithTokenForTests creates a new route with token omission for tests
func NewRouteWithTokenForTests(tokens Tokens, token Omission) Route {
	ins, err := NewBuilder().Create().WithTokens(tokens).WithToken(token).Now()
	if err != nil {
		panic(err)
	}

	return ins
}

// NewRouteForTests creates a new route for tests
func NewRouteForTests(tokens Tokens) Route {
	ins, err := NewBuilder().Create().WithTokens(tokens).Now()
	if err != nil {
		panic(err)
	}

	return ins
}

// NewTokensForTests creates tokens for tests
func NewTokensForTests(list []Token) Tokens {
	ins, err := NewTokensBuilder().Create().WithList(list).Now()
	if err != nil {
		panic(err)
	}

	return ins
}

// NewTokenWithOmissionForTests creates token with omission for tests
func NewTokenWithOmissionForTests(elements Elements, cardinality cardinalities.Cardinality, omission Omission) Token {
	ins, err := NewTokenBuilder().Create().WithElements(elements).WithCardinality(cardinality).WithOmission(omission).Now()
	if err != nil {
		panic(err)
	}

	return ins
}

// NewTokenForTests creates token for tests
func NewTokenForTests(elements Elements, cardinality cardinalities.Cardinality) Token {
	ins, err := NewTokenBuilder().Create().WithElements(elements).WithCardinality(cardinality).Now()
	if err != nil {
		panic(err)
	}

	return ins
}

// NewOmissionWithPrefixAndSuffixForTests creates omission with prefix and suffix for tests
func NewOmissionWithPrefixAndSuffixForTests(prefix Element, suffix Element) Omission {
	ins, err := NewOmissionBuilder().Create().WithPrefix(prefix).WithSuffix(suffix).Now()
	if err != nil {
		panic(err)
	}

	return ins
}

// NewOmissionWithPrefixForTests creates omission with prefix for tests
func NewOmissionWithPrefixForTests(prefix Element) Omission {
	ins, err := NewOmissionBuilder().Create().WithPrefix(prefix).Now()
	if err != nil {
		panic(err)
	}

	return ins
}

// NewOmissionWithSuffixForTests creates omission with suffix for tests
func NewOmissionWithSuffixForTests(suffix Element) Omission {
	ins, err := NewOmissionBuilder().Create().WithSuffix(suffix).Now()
	if err != nil {
		panic(err)
	}

	return ins
}

// NewElementsForTests creates elements for tests
func NewElementsForTests(list []Element) Elements {
	ins, err := NewElementsBuilder().Create().WithList(list).Now()
	if err != nil {
		panic(err)
	}

	return ins
}

// NewElementWithLayerForTests creates element with layer for tests
func NewElementWithLayerForTests(layer hash.Hash) Element {
	ins, err := NewElementBuilder().Create().WithLayer(layer).Now()
	if err != nil {
		panic(err)
	}

	return ins
}

// NewElementWithBytesForTests creates element with bytes for tests
func NewElementWithBytesForTests(bytes []byte) Element {
	ins, err := NewElementBuilder().Create().WithBytes(bytes).Now()
	if err != nil {
		panic(err)
	}

	return ins
}

// NewElementWithStringForTests creates element with string for tests
func NewElementWithStringForTests(str string) Element {
	ins, err := NewElementBuilder().Create().WithString(str).Now()
	if err != nil {
		panic(err)
	}

	return ins
}
