package executelayerpaths

// ExecuteLayerPath represents a layer path
type ExecuteLayerPath interface {
	Context() string
	InputPath() string
	LayerPath() string
	Return() string
}
