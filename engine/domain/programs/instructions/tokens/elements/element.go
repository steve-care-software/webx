package elements

import "github.com/steve-care-software/webx/engine/domain/programs/instructions/tokens/elements/syscalls"

type element struct {
	rule        string
	syscall     syscalls.Syscall
	instruction string
}

func createElementWithRule(rule string) Element {
	return createElementInternally(rule, nil, "")
}

func createElementWithSyscall(syscall syscalls.Syscall) Element {
	return createElementInternally("", syscall, "")
}

func createElementWithInstruction(instruction string) Element {
	return createElementInternally("", nil, instruction)
}

func createElementInternally(
	rule string,
	syscall syscalls.Syscall,
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
		return obj.syscall.Name()
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
	return obj.syscall != nil
}

// Syscall returns the syscall, if any
func (obj *element) Syscall() syscalls.Syscall {
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
