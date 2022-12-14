package modules

type module struct {
	index uint
	fn    ExecuteFn
}

func createModule(
	index uint,
	fn ExecuteFn,
) Module {
	out := module{
		index: index,
		fn:    fn,
	}

	return &out
}

// Index returns the index
func (obj *module) Index() uint {
	return obj.index
}

// Func returns the execute fn
func (obj *module) Func() ExecuteFn {
	return obj.fn
}
