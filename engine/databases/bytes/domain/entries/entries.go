package entries

type entries struct {
	list []Entry
}

func createEntries(
	list []Entry,
) Entries {
	out := entries{
		list: list,
	}

	return &out
}

// List returns the list of entries
func (obj *entries) List() []Entry {
	return obj.list
}
