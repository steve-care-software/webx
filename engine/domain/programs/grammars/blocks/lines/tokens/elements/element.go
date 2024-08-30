package elements

type element struct {
	rule     string
	block    string
	spacer   string
	constant string
}

func createElementWithRule(rule string) Element {
	return createElementInternally(rule, "", "", "")
}

func createElementWithBlock(block string) Element {
	return createElementInternally("", block, "", "")
}

func createElementWithSpacer(spacer string) Element {
	return createElementInternally("", "", spacer, "")
}

func createElementWithConstant(constant string) Element {
	return createElementInternally("", "", "", constant)
}

func createElementInternally(
	rule string,
	block string,
	spacer string,
	constant string,
) Element {
	out := element{
		rule:     rule,
		block:    block,
		spacer:   spacer,
		constant: constant,
	}

	return &out
}

// Name returns the rule or block name
func (obj *element) Name() string {
	if obj.IsBlock() {
		return obj.block
	}

	return obj.rule
}

// IsRule returns true if there is a rule, false otherwise
func (obj *element) IsRule() bool {
	return obj.rule != ""
}

// Rule returns the rule, if any
func (obj *element) Rule() string {
	return obj.rule
}

// IsBlock returns true if there is a block, false otherwise
func (obj *element) IsBlock() bool {
	return obj.block != ""
}

// Block returns the block, if any
func (obj *element) Block() string {
	return obj.block
}

// IsSpacer returns true if there is a spacer, false otherwise
func (obj *element) IsSpacer() bool {
	return obj.spacer != ""
}

// Spacer returns the spacer, if any
func (obj *element) Spacer() string {
	return obj.spacer
}

// IsConstant returns true if there is a constant, false otherwise
func (obj *element) IsConstant() bool {
	return obj.constant != ""
}

// Constant returns the constant, if any
func (obj *element) Constant() string {
	return obj.constant
}
