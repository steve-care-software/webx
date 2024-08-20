package elements

type element struct {
	rule    string
	block   string
	syscall string
}

func createElementWithRule(rule string) Element {
	return createElementInternally(rule, "", "")
}

func createElementWithBlock(block string) Element {
	return createElementInternally("", block, "")
}

func createElementWithSyscall(syscall string) Element {
	return createElementInternally("", "", syscall)
}

func createElementInternally(
	rule string,
	block string,
	syscall string,
) Element {
	out := element{
		rule:    rule,
		block:   block,
		syscall: syscall,
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

// IsSyscall returns true if there is a syscall, false otherwise
func (obj *element) IsSyscall() bool {
	return obj.syscall != ""
}

// Syscall returns the syscall, if any
func (obj *element) Syscall() string {
	return obj.syscall
}
