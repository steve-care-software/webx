package retrieves

// Retrieve represents a retrieve instruction
type Retrieve interface {
	Context() string
	Index() string
	Return() string
	HasLength() bool
	Length() string
}
