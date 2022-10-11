package values

type value struct {
	name   string
	number byte
}

func createValue(
	name string,
	number byte,
) Value {
	out := value{
		name:   name,
		number: number,
	}

	return &out
}

// Name returns the name
func (obj *value) Name() string {
	return obj.name
}

// Number returns the number
func (obj *value) Number() byte {
	return obj.number
}
