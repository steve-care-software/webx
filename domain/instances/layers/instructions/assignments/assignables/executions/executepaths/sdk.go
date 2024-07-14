package executepaths

// ExecutePath represents an execute path
type ExecutePath interface {
	Context() string
	InputPath() string
	Return() string
}
