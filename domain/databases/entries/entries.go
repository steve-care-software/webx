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

// List returns the list
func (obj *entries) List() []Entry {
	return obj.list
}
