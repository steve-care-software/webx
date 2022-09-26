package modules

type module struct {
	name string
	fn   ExecuteFn
}

func createModule(
	name string,
	fn ExecuteFn,
) Module {
	out := module{
		name: name,
		fn:   fn,
	}

	return &out
}

// Name returns the name
func (obj *module) Name() string {
	return obj.name
}

// Func returns the execute fn
func (obj *module) Func() ExecuteFn {
	return obj.fn
}
