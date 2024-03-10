package resources

type list struct {
	value     uint8
	delimiter string
}

func createList(
	value uint8,
	delimiter string,
) List {
	out := list{
		value:     value,
		delimiter: delimiter,
	}

	return &out
}

// Value returns the value
func (obj *list) Value() uint8 {
	return obj.value
}

// Delimiter returns the delimiter
func (obj *list) Delimiter() string {
	return obj.delimiter
}
