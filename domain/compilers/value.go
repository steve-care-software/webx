package compilers

import "github.com/steve-care-software/webx/domain/criterias"

type value struct {
	constant interface{}
	criteria criterias.Criteria
}

func createValueWithNil() Value {
	return createValueInternally(nil, nil)
}

func createValueWithConstant(
	constant interface{},
) Value {
	return createValueInternally(constant, nil)
}

func createValueWithCriteria(
	criteria criterias.Criteria,
) Value {
	return createValueInternally(nil, criteria)
}

func createValueInternally(
	constant interface{},
	criteria criterias.Criteria,
) Value {
	out := value{
		constant: constant,
		criteria: criteria,
	}

	return &out
}

// IsConstant returns true if there is a constant, false otherwise
func (obj *value) IsConstant() bool {
	// if the value is nil:
	if obj.criteria == nil && obj.constant == nil {
		return true
	}

	return obj.constant != nil
}

// Constant returns the constant, if any
func (obj *value) Constant() interface{} {
	return obj.constant
}

// IsCriteria returns true if there is a criteria, false otherwise
func (obj *value) IsCriteria() bool {
	return obj.criteria != nil
}

// Criteria returns the criteria, if any
func (obj *value) Criteria() criterias.Criteria {
	return obj.criteria
}
