package instructions

import "github.com/steve-care-software/webx/engine/domain/programs/grammars/rules"

type element struct {
	rule        rules.Rule
	instruction Instruction
}

func createElementWithRule(rule rules.Rule) Element {
	return createElementInternally(rule, nil)
}

func createElementWithInstruction(instruction Instruction) Element {
	return createElementInternally(nil, instruction)
}

func createElementInternally(
	rule rules.Rule,
	instruction Instruction,
) Element {
	out := element{
		rule:        rule,
		instruction: instruction,
	}

	return &out
}

// Name returns the name
func (obj *element) Name() string {
	if obj.IsRule() {
		return obj.rule.Name()
	}

	return obj.instruction.Block()
}

// IsRule returns true if there is a rule, false otherwise
func (obj *element) IsRule() bool {
	return obj.rule != nil
}

// Rule returns the rule, if any
func (obj *element) Rule() rules.Rule {
	return obj.rule
}

// IsInstruction returns true if there is an instruction, false otherwise
func (obj *element) IsInstruction() bool {
	return obj.instruction != nil
}

// Instruction returns the instruction, if any
func (obj *element) Instruction() Instruction {
	return obj.instruction
}
