package updates

type update struct {
	name        string
	description string
}

func createUpdate(
	name string,
	description string,
) Update {
	out := update{
		name:        name,
		description: description,
	}

	return &out
}

// Name returns the name
func (obj *update) Name() string {
	return obj.name
}

// Description returns the description
func (obj *update) Description() string {
	return obj.description
}
