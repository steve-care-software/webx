package elements

type element struct {
	rule  string
	block string
}

func createElementWithRule(rule string) Element {
	return createElementInternally(rule, "")
}

func createElementWithBlock(block string) Element {
	return createElementInternally("", block)
}

func createElementInternally(
	rule string,
	block string,
) Element {
	out := element{
		rule:  rule,
		block: block,
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
