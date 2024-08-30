package elements

type element struct {
	rule     string
	constant string
}

func createElementWithRule(
	rule string,
) Element {
	return createElementInternally(rule, "")
}

func createElementWithConstant(
	constant string,
) Element {
	return createElementInternally("", constant)
}

func createElementInternally(
	rule string,
	constant string,
) Element {
	out := element{
		rule:     rule,
		constant: constant,
	}

	return &out
}

// IsRule returns true if there is a rule, false otherwise
func (obj *element) IsRule() bool {
	return obj.rule != ""
}

// Rule returns the rule, if any
func (obj *element) Rule() string {
	return obj.rule
}

// IsConstant returns true if there is a constant, false otherwise
func (obj *element) IsConstant() bool {
	return obj.constant != ""
}

// Constant returns the constant, if any
func (obj *element) Constant() string {
	return obj.constant
}
