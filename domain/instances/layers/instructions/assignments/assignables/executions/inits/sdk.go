package inits

// Init represents an init
type Init interface {
	DbPath() string
	Name() string
	Description() string
	Context() string
}
