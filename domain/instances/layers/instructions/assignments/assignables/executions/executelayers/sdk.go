package executelayers

// ExecuteLayer represents an execute layer
type ExecuteLayer interface {
	Context() string
	Input() string
	LayerPath() string
	Return() string
}
