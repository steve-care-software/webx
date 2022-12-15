package selectors

import "github.com/steve-care-software/webx/roots/domain/grammars/grammars"

type selector struct {
	grammar grammars.Grammar
	token   Token
	inside  Inside
	fn      SelectorFn
}

func createSelector(
	grammar grammars.Grammar,
	token Token,
	inside Inside,
	fn SelectorFn,
) Selector {
	out := selector{
		token:  token,
		inside: inside,
		fn:     fn,
	}

	return &out
}

// Grammar returns the grammar
func (obj *selector) Grammar() grammars.Grammar {
	return obj.grammar
}

// Token returns the token
func (obj *selector) Token() Token {
	return obj.token
}

// Inside returns the inside
func (obj *selector) Inside() Inside {
	return obj.inside
}

// Fn returns the func
func (obj *selector) Fn() SelectorFn {
	return obj.fn
}
