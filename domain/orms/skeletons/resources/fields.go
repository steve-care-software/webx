package resources

type fields struct {
	list []Field
}

func createFields(
	list []Field,
) Fields {
	out := fields{
		list: list,
	}

	return &out
}

// List returns the list
func (obj *fields) List() []Field {
	return obj.list
}

// Names returns the field names
func (obj *fields) Names() []string {
	output := []string{}
	for _, oneField := range obj.list {
		output = append(output, oneField.Name())
	}

	return output
}
