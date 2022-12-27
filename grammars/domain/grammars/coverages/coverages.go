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

// ContainsError returns true if it contains an error, false otherwise
func (obj *coverages) ContainsError() bool {
	for _, oneCoverage := range obj.list {
		if !oneCoverage.Executions().ContainsError() {
			continue
		}

		return true
	}

	return false
}
