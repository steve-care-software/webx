package values

type value struct {
	name   string
	number uint8
}

func createValue(
	name string,
	number uint8,
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
func (obj *value) Number() uint8 {
	return obj.number
}
