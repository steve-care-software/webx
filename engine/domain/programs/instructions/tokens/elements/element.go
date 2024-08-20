package elements

type element struct {
	rule        string
	syscall     string
	instruction string
}

func createElementWithRule(rule string) Element {
	return createElementInternally(rule, "", "")
}

func createElementWithSyscall(syscall string) Element {
	return createElementInternally("", syscall, "")
}

func createElementWithInstruction(instruction string) Element {
	return createElementInternally("", "", instruction)
}

func createElementInternally(
	rule string,
	syscall string,
	instruction string,
) Element {
	out := element{
		rule:        rule,
		syscall:     syscall,
		instruction: instruction,
	}

	return &out
}

// Name returns the name
func (obj *element) Name() string {
	if obj.IsRule() {
		return obj.rule
	}

	if obj.IsSyscall() {
		return obj.syscall
	}

	return obj.instruction
}

// IsRule returns true if there is a rule, false otherwise
func (obj *element) IsRule() bool {
	return obj.rule != ""
}

// Rule returns the rule, if any
func (obj *element) Rule() string {
	return obj.rule
}

// IsSyscall returns true if there is a syscall, false otherwise
func (obj *element) IsSyscall() bool {
	return obj.syscall != ""
}

// Syscall returns the syscall, if any
func (obj *element) Syscall() string {
	return obj.syscall
}

// IsInstruction returns true if there is an instruction, false otherwise
func (obj *element) IsInstruction() bool {
	return obj.instruction != ""
}

// Instruction returns the instruction, if any
func (obj *element) Instruction() string {
	return obj.instruction
}
