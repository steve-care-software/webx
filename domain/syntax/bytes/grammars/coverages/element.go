package coverages

type element struct {
	name  string
	value []byte
}

func createElement(
	name string,
) Element {
	return createElementInternally(name, nil)
}

func createElementWithValue(
	name string,
	value []byte,
) Element {
	return createElementInternally(name, value)
}

func createElementInternally(
	name string,
	value []byte,
) Element {
	out := element{
		name:  name,
		value: value,
	}

	return &out
}

// Name returns the name
func (obj *element) Name() string {
	return obj.name
}

// HasValue returns true if there is a value, false otherwise
func (obj *element) HasValue() bool {
	return obj.value != nil
}

// Value returns the value, if any
func (obj *element) Value() []byte {
	return obj.value
}
