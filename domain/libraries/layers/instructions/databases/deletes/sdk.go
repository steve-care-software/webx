package deletes

// Delete represents a delete
type Delete interface {
	Context() string
	Path() string
	Identifier() string
}
