package modules

type module struct {
	index uint
	name  []byte
}

func createModule(
	index uint,
	name []byte,
) Module {
	out := module{
		index: index,
		name:  name,
	}

	return &out
}

// Index returns the index
func (obj *module) Index() uint {
	return obj.index
}

// Name returns the name
func (obj *module) Name() []byte {
	return obj.name
}
