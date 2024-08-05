package originals

type original struct {
	name        string
	description string
}

func createOriginal(
	name string,
	description string,
) Original {
	out := original{
		name:        name,
		description: description,
	}

	return &out
}

// Name returns the name
func (obj *original) Name() string {
	return obj.name
}

// Description returns the description
func (obj *original) Description() string {
	return obj.description
}
