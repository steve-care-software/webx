package compilers

type value struct {
	constant string
	//criteria criterias.Criteria
}

func createValueWithNil() Value {
	return createValueInternally("")
}

func createValueWithConstant(
	constant string,
) Value {
	return createValueInternally(constant)
}

func createValueWithCriteria(
//criteria criterias.Criteria,
) Value {
	return createValueInternally("")
}

func createValueInternally(
	constant string,
	//criteria criterias.Criteria,
) Value {
	out := value{
		constant: constant,
		//criteria: criteria,
	}

	return &out
}

// IsConstant returns true if there is a constant, false otherwise
func (obj *value) IsConstant() bool {
	// if the value is nil:
	/*if obj.criteria == nil && obj.constant == "" {
		return true
	}*/

	return obj.constant != ""
}

// Constant returns the constant, if any
func (obj *value) Constant() string {
	return obj.constant
}

// IsCriteria returns true if there is a criteria, false otherwise
/*func (obj *value) IsCriteria() bool {
	return obj.criteria != nil
}

// Criteria returns the criteria, if any
func (obj *value) Criteria() criterias.Criteria {
	return obj.criteria
}*/
