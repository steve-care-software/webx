package originals

type original struct {
	name        string
	description string
	isDeleted   bool
}

func createOriginal(
	name string,
	description string,
	isDeleted bool,
) Original {
	out := original{
		name:        name,
		description: description,
		isDeleted:   isDeleted,
	}

	return &out
}

// IsDeleted returns true if deleted, false otherwise
func (obj *original) IsDeleted() bool {
	return obj.isDeleted
}

// Name returns the name
func (obj *original) Name() string {
	return obj.name
}

// Description returns the description
func (obj *original) Description() string {
	return obj.description
}
