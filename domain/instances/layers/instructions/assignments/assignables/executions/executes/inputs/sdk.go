package inputs

// Input represents an input
type Input interface {
	IsValue() bool
	Value() string
	IsPath() bool
	Path() string
}
