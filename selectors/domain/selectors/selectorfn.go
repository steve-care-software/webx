package selectors

type selectorFn struct {
	single SingleSelectorFn
	multi  MultiSelectorFn
}

func createSelectorFnWithSingle(
	single SingleSelectorFn,
) SelectorFn {
	return createSelectorFnInternally(single, nil)
}

func createSelectorFnWithMulti(
	multi MultiSelectorFn,
) SelectorFn {
	return createSelectorFnInternally(nil, multi)
}

func createSelectorFnInternally(
	single SingleSelectorFn,
	multi MultiSelectorFn,
) SelectorFn {
	out := selectorFn{
		single: single,
		multi:  multi,
	}

	return &out
}

// IsSingle returns true if single, false otherwise
func (obj *selectorFn) IsSingle() bool {
	return obj.single != nil
}

// Single returns the single selector func, if any
func (obj *selectorFn) Single() SingleSelectorFn {
	return obj.single
}

// IsMulti returns true if multi, false otherwise
func (obj *selectorFn) IsMulti() bool {
	return obj.multi != nil
}

// Multi returns the single multi func, if any
func (obj *selectorFn) Multi() MultiSelectorFn {
	return obj.multi
}
