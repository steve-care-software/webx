package repositories

// NewRepositoryWithSkeletonForTests creates a new repository with skeleton for tests
func NewRepositoryWithSkeletonForTests() Repository {
	ins, err := NewBuilder().Create().IsSkeleton().Now()
	if err != nil {
		panic(err)
	}

	return ins
}

// NewRepositoryWithHeightForTests creates a new repository with height for tests
func NewRepositoryWithHeightForTests() Repository {
	ins, err := NewBuilder().Create().IsHeight().Now()
	if err != nil {
		panic(err)
	}

	return ins
}

// NewRepositoryWithListForTests creates a new repository with list for tests
func NewRepositoryWithListForTests(list string) Repository {
	ins, err := NewBuilder().Create().WithList(list).Now()
	if err != nil {
		panic(err)
	}

	return ins
}

// NewRepositoryWithRetrieveForTests creates a new repository with retrieve for tests
func NewRepositoryWithRetrieveForTests(retrieve string) Repository {
	ins, err := NewBuilder().Create().WithRetrieve(retrieve).Now()
	if err != nil {
		panic(err)
	}

	return ins
}
