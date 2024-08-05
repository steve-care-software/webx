package status

// Status represents the satus
type Status interface {
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
