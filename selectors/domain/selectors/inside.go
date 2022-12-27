package selectors

type inside struct {
	fn       ContentFn
	fetchers Fetchers
}

func createInsideWithFunc(
	fn ContentFn,
) Inside {
	return createInsideInternally(fn, nil)
}

func createInsideWithFetchers(
	fetchers Fetchers,
) Inside {
	return createInsideInternally(nil, fetchers)
}

func createInsideInternally(
	fn ContentFn,
	fetchers Fetchers,
) Inside {
	out := inside{
		fn:       fn,
		fetchers: fetchers,
	}

	return &out
}

// IsFn returns true if there is a func, false otherwise
func (obj *inside) IsFn() bool {
	return obj.fn != nil
}

// Fn returns the func, if any
func (obj *inside) Fn() ContentFn {
	return obj.fn
}

// IsFetchers returns true if there is fetchers, false otherwise
func (obj *inside) IsFetchers() bool {
	return obj.fetchers != nil
}

// Fetchers returns the fetchers, if any
func (obj *inside) Fetchers() Fetchers {
	return obj.fetchers
}
