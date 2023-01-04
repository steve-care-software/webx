package results

type result struct {
	isValid   bool
	values    []interface{}
	remaining []byte
}

func createResult(isValid bool) Result {
	return createResultInternally(isValid, nil, nil)
}

func createResultWithValues(
	isValid bool,
	values []interface{},
) Result {
	return createResultInternally(isValid, values, nil)
}

func createResultWithRemaining(
	isValid bool,
	remaining []byte,
) Result {
	return createResultInternally(isValid, nil, remaining)
}

func createResultWithValuesAndRemaining(
	isValid bool,
	values []interface{},
	remaining []byte,
) Result {
	return createResultInternally(isValid, values, remaining)
}

func createResultInternally(
	isValid bool,
	values []interface{},
	remaining []byte,
) Result {
	out := result{
		isValid:   isValid,
		values:    values,
		remaining: remaining,
	}

	return &out
}

// IsValid returns true if valid, false otherwise
func (obj *result) IsValid() bool {
	return obj.isValid
}

// HasValues returns true if there is values, false otherwise
func (obj *result) HasValues() bool {
	return obj.values != nil
}

// Values returns the values, if any
func (obj *result) Values() []interface{} {
	return obj.values
}

// HasRemaining returns true if there is remaining script, false otherwise
func (obj *result) HasRemaining() bool {
	return obj.remaining != nil
}

// Remaining returns the remaining script, if any
func (obj *result) Remaining() []byte {
	return obj.remaining
}
