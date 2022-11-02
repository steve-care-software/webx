package modules

type module struct {
	name []byte
	fn   ExecuteFn
}

func createModule(
	name []byte,
	fn ExecuteFn,
) Module {
	out := module{
		name: name,
		fn:   fn,
	}

	return &out
}

// Name returns the name
func (obj *module) Name() []byte {
	return obj.name
}

// Func returns the execute fn
func (obj *module) Func() ExecuteFn {
	return obj.fn
}
