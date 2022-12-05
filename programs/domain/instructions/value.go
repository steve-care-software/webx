package instructions

type value struct {
	variable     []byte
	constant     []byte
	instructions Instructions
	execution    []byte
}

func createValueWithVariable(
	variable []byte,
) Value {
	return createValueInternally(variable, nil, nil, nil)
}

func createValueWithConstant(
	constant []byte,
) Value {
	return createValueInternally(nil, constant, nil, nil)
}

func createValueWithInstructions(
	instructions Instructions,
) Value {
	return createValueInternally(nil, nil, instructions, nil)
}

func createValueWithExecution(
	execution []byte,
) Value {
	return createValueInternally(nil, nil, nil, execution)
}

func createValueInternally(
	variable []byte,
	constant []byte,
	instructions Instructions,
	execution []byte,
) Value {
	out := value{
		variable:     variable,
		constant:     constant,
		instructions: instructions,
		execution:    execution,
	}

	return &out
}

// IsVariable returns true if there is a variable, false otherwise
func (obj *value) IsVariable() bool {
	return obj.variable != nil
}

// Variable returns the variable, if any
func (obj *value) Variable() []byte {
	return obj.variable
}

// IsConstant returns true if there is a constant, false otherwise
func (obj *value) IsConstant() bool {
	return obj.constant != nil
}

// Constant returns the constant, if any
func (obj *value) Constant() []byte {
	return obj.constant
}

// IsInstructions returns true if there is instructions, false otherwise
func (obj *value) IsInstructions() bool {
	return obj.instructions != nil
}

// Instructions returns the instructions, if any
func (obj *value) Instructions() Instructions {
	return obj.instructions
}

// IsExecution returns true if there is an execution, false otherwise
func (obj *value) IsExecution() bool {
	return obj.execution != nil
}

// Execution returns the execution, if any
func (obj *value) Execution() []byte {
	return obj.execution
}
