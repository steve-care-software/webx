package stacks

// Factory represents a stack factory
type Factory interface {
	Create() Stack
}

// Stack represents a stack instance
type Stack interface {
	Head() Frame
	Height() uint
	Push() Stack
	Pop() (Stack, error)
}

// Frame represents a frame
type Frame interface {
	Register(name string, value any, replaceIfExists bool) error
	Fetch(name string) (any, error)
}
