package stacks

type factory struct {
}

func createFactory() Factory {
	out := factory{}
	return &out
}

// Create creates a new stack
func (app *factory) Create() Stack {
	return createStack()
}
