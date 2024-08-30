package suites

import "github.com/steve-care-software/webx/engine/domain/programs/grammars/blocks/suites/validations"

type suite struct {
	name        string
	value       []byte
	isFail      bool
	validations validations.Validations
}

func createSuite(
	name string,
	value []byte,
	isFail bool,
) Suite {
	return createSuiteInternally(
		name,
		value,
		isFail,
		nil,
	)
}

func createSuiteWithValidations(
	name string,
	value []byte,
	isFail bool,
	validations validations.Validations,
) Suite {
	return createSuiteInternally(
		name,
		value,
		isFail,
		validations,
	)
}

func createSuiteInternally(
	name string,
	value []byte,
	isFail bool,
	validations validations.Validations,
) Suite {
	out := suite{
		name:        name,
		value:       value,
		isFail:      isFail,
		validations: validations,
	}

	return &out
}

// Name returns the name
func (obj *suite) Name() string {
	return obj.name
}

// Value returns the value
func (obj *suite) Value() []byte {
	return obj.value
}

// IsFail returns true if expected to fail, false otherwise
func (obj *suite) IsFail() bool {
	return obj.isFail
}

// HasValidations returns true if there is validations, false otherwise
func (obj *suite) HasValidations() bool {
	return obj.validations != nil
}

// Validations returns the validations, if any
func (obj *suite) Validations() validations.Validations {
	return obj.validations
}
