package conditions

// NewConditionForTests creates a new condition for tests
func NewConditionForTests(pointer Pointer, operator Operator, element Element) Condition {
	ins, err := NewBuilder().Create().
		WithPointer(pointer).
		WithOperator(operator).
		WithElement(element).
		Now()

	if err != nil {
		panic(err)
	}

	return ins
}

// NewElementWithResourceForTests creates a new element with resource for tests
func NewElementWithResourceForTests(resource Resource) Element {
	ins, err := NewElementBuilder().Create().WithResource(resource).Now()
	if err != nil {
		panic(err)
	}

	return ins
}

// NewElementWithConditionForTests creates a new element with condition for tests
func NewElementWithConditionForTests(condition Condition) Element {
	ins, err := NewElementBuilder().Create().WithCondition(condition).Now()
	if err != nil {
		panic(err)
	}

	return ins
}

// NewResourceWithValueForTests creates a new resource with value for tests
func NewResourceWithValueForTests(value interface{}) Resource {
	ins, err := NewResourceBuilder().Create().WithValue(value).Now()
	if err != nil {
		panic(err)
	}

	return ins
}

// NewResourceWithFieldForTests creates a new resource with field for tests
func NewResourceWithFieldForTests(field Pointer) Resource {
	ins, err := NewResourceBuilder().Create().WithField(field).Now()
	if err != nil {
		panic(err)
	}

	return ins
}

// NewPointerForTests creates a new pointer for tests
func NewPointerForTests(entity string, field string) Pointer {
	ins, err := NewPointerBuilder().Create().WithEntity(entity).WithField(field).Now()
	if err != nil {
		panic(err)
	}

	return ins
}

// NewOperatorWithIntegerForTests creates a new operator with integer for tests
func NewOperatorWithIntegerForTests(integer IntegerOperator) Operator {
	ins, err := NewOperatorBuilder().Create().WithInteger(integer).Now()
	if err != nil {
		panic(err)
	}

	return ins
}

// NewOperatorWithRelationalForTests creates a new operator with relational for tests
func NewOperatorWithRelationalForTests(relational RelationalOperator) Operator {
	ins, err := NewOperatorBuilder().Create().WithRelational(relational).Now()
	if err != nil {
		panic(err)
	}

	return ins
}

// NewOperatorWithEqualForTests creates a new operator with equal for tests
func NewOperatorWithEqualForTests() Operator {
	ins, err := NewOperatorBuilder().Create().IsEqual().Now()
	if err != nil {
		panic(err)
	}

	return ins
}

// NewRelationalOperatorWithOrForTests creates a relational operator with or for tests
func NewRelationalOperatorWithOrForTests() RelationalOperator {
	ins, err := NewRelationalOperatorBuilder().Create().IsOr().Now()
	if err != nil {
		panic(err)
	}

	return ins
}

// NewRelationalOperatorWithAndForTests creates a relational operator with and for tests
func NewRelationalOperatorWithAndForTests() RelationalOperator {
	ins, err := NewRelationalOperatorBuilder().Create().IsAnd().Now()
	if err != nil {
		panic(err)
	}

	return ins
}

// NewIntegerOperatorWithIsEqualForTests creates a new integer operator with IsEqual for tests
func NewIntegerOperatorWithIsEqualForTests() IntegerOperator {
	ins, err := NewIntegerOperatorBuilder().Create().IsEqual().Now()
	if err != nil {
		panic(err)
	}

	return ins
}

// NewIntegerOperatorWithIsBiggerThanForTests creates a new integer operator with IsBiggerThan for tests
func NewIntegerOperatorWithIsBiggerThanForTests() IntegerOperator {
	ins, err := NewIntegerOperatorBuilder().Create().IsBiggerThan().Now()
	if err != nil {
		panic(err)
	}

	return ins
}

// NewIntegerOperatorWithIsBiggerThanAndIsEqualForTests creates a new integer operator with IsBiggerThan and isEqual for tests
func NewIntegerOperatorWithIsBiggerThanAndIsEqualForTests() IntegerOperator {
	ins, err := NewIntegerOperatorBuilder().Create().IsBiggerThan().IsEqual().Now()
	if err != nil {
		panic(err)
	}

	return ins
}

// NewIntegerOperatorWithIsSmallerThanForTests creates a new integer operator with IsSmallerThan for tests
func NewIntegerOperatorWithIsSmallerThanForTests() IntegerOperator {
	ins, err := NewIntegerOperatorBuilder().Create().IsSmallerThan().Now()
	if err != nil {
		panic(err)
	}

	return ins
}

// NewIntegerOperatorWithIsSmallerThanAndIsEqualForTests creates a new integer operator with IsSmallerThan and isEqual for tests
func NewIntegerOperatorWithIsSmallerThanAndIsEqualForTests() IntegerOperator {
	ins, err := NewIntegerOperatorBuilder().Create().IsSmallerThan().IsEqual().Now()
	if err != nil {
		panic(err)
	}

	return ins
}
