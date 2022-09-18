package modules

// ExecuteFn represents the execute func
type ExecuteFn func(input map[string]interface{}) (interface{}, error)

// Module represents a module
type Module interface {
	Name() string
	Func() ExecuteFn
}
