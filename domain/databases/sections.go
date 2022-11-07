package databases

type sections struct {
	list []Section
}

func createSections(
	list []Section,
) Sections {
	out := sections{
		list: list,
	}

	return &out
}

// List returns the list of sections
func (obj *sections) List() []Section {
	return obj.list
}
