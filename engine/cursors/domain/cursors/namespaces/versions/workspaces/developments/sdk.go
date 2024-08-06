package developments

// Development represents a development workspace
type Development interface {
	Name() string
	HasBranch() bool
}
