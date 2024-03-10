package inserts

// Insert represents an insert
type Insert interface {
	Context() string
	Instance() string
	Path() string
}
