package suites

type suites struct {
	list []Suite
}

func createSuites(
	list []Suite,
) Suites {
	out := suites{
		list: list,
	}

	return &out
}

// List returns the list of suite
func (obj *suites) List() []Suite {
	return obj.list
}
