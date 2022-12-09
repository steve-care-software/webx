package compilers

import "github.com/steve-care-software/webx/selectors/domain/selectors"

type value struct {
	constant string
	selector selectors.Selector
}

func createValueWithNil() Value {
	return createValueInternally("", nil)
}

func createValueWithConstant(
	constant string,
) Value {
	return createValueInternally(constant, nil)
}

func createValueWithSelector(
	selector selectors.Selector,
) Value {
	return createValueInternally("", selector)
}

func createValueInternally(
	constant string,
	selector selectors.Selector,
) Value {
	out := value{
		constant: constant,
		selector: selector,
	}

	return &out
}

// IsConstant returns true if there is a constant, false otherwise
func (obj *value) IsConstant() bool {
	// if the value is nil:
	if obj.selector == nil && obj.constant == "" {
		return true
	}

	return obj.constant != ""
}

// Constant returns the constant, if any
func (obj *value) Constant() string {
	return obj.constant
}

// IsSelector returns true if there is a selector, false otherwise
func (obj *value) IsSelector() bool {
	return obj.selector != nil
}

// Selector returns the criteria, if any
func (obj *value) Selector() selectors.Selector {
	return obj.selector
}
