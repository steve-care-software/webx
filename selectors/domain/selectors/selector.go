package selectors

type selector struct {
	token  Token
	inside Inside
	fn     SelectorFn
}

func createSelector(
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
