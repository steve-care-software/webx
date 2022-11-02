package selectors

type contentFn struct {
	single SingleContentFn
	multi  MultiContentFn
}

func createContentFnWithSingle(
	single SingleContentFn,
) ContentFn {
	return createContentFnInternally(single, nil)
}

func createContentFnWithMulti(
	multi MultiContentFn,
) ContentFn {
	return createContentFnInternally(nil, multi)
}

func createContentFnInternally(
	single SingleContentFn,
	multi MultiContentFn,
) ContentFn {
	out := contentFn{
		single: single,
		multi:  multi,
	}

	return &out
}

// IsSingle returns true if single, false otherwise
func (obj *contentFn) IsSingle() bool {
	return obj.single != nil
}

// Single returns the single content func, if any
func (obj *contentFn) Single() SingleContentFn {
	return obj.single
}

// IsMulti returns true if multi, false otherwise
func (obj *contentFn) IsMulti() bool {
	return obj.multi != nil
}

// Multi returns the single multi func, if any
func (obj *contentFn) Multi() MultiContentFn {
	return obj.multi
}
