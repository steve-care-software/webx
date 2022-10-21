package coverages

type coverages struct {
	list []Coverage
}

func createCoverages(
	list []Coverage,
) Coverages {
	out := coverages{
		list: list,
	}

	return &out
}

// List returns the coverages
func (obj *coverages) List() []Coverage {
	return obj.list
}
