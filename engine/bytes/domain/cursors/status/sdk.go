package status

// Statuses represents statuses
type Statuses interface {
	List() []Status
}

// Status represents the satus
type Status interface {
	Name() string
	HasNamespace() bool
	Namespace() string
	HasVersion() bool
	Version() string
	HasIteration() bool
	Iteration() string
	HasWorkspace() bool
	Workspace() string
	HasBranch() bool
	Branch() string
	HasState() bool
	State() string
}
