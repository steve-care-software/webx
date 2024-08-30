package validations

type validations struct {
	list []Validation
}

func createValidations(
	list []Validation,
) Validations {
	out := validations{
		list: list,
	}

	return &out
}

// List returns the list of validations
func (obj *validations) List() []Validation {
	return obj.list
}
