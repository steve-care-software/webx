package modules

type modules struct {
	list []Module
}

func createModules(
	list []Module,
) Modules {
	out := modules{
		list: list,
	}

	return &out
}

// List returns the modules
func (obj *modules) List() []Module {
	return obj.list
}
