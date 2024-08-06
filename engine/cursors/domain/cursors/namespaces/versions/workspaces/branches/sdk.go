package branches

// Branch represents a branch
type Branch interface {
	Name() string
	IsState() bool
	State() string
	IsChild() bool
	Child() Branch
}
