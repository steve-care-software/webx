package commits

// NewCommitForTests creates a new commit for tests
func NewCommitForTests(description string, actions string) Commit {
	ins, err := NewBuilder().Create().
		WithDescription(description).
		WithActions(actions).
		Now()

	if err != nil {
		panic(err)
	}

	return ins
}

// NewCommitWithParentForTests creates a new commit with parent for tests
func NewCommitWithParentForTests(description string, actions string, parent string) Commit {
	ins, err := NewBuilder().Create().
		WithDescription(description).
		WithActions(actions).
		WithParent(parent).
		Now()

	if err != nil {
		panic(err)
	}

	return ins
}
