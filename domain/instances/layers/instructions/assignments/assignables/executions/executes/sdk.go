package executes

// Execute represents an execute
type Execute interface {
	Context() string
	Input() string
	Return() string
}
