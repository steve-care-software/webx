package connections

type field struct {
	name string
	path []string
}

func createField(
	name string,
	path []string,
) Field {
	out := field{
		name: name,
		path: path,
	}

	return &out
}

// Name returns the name
func (obj *field) Name() string {
	return obj.name
}

// Path returns the path
func (obj *field) Path() []string {
	return obj.path
}
