package instructions

type value struct {
	input        string
	str          string
	instructions Instructions
	execution    string
}

func createValueWithInput(
	input string,
) Value {
	return createValueInternally(input, "", nil, "")
}

func createValueWithString(
	str string,
) Value {
	return createValueInternally("", str, nil, "")
}

func createValueWithInstructions(
	instructions Instructions,
) Value {
	return createValueInternally("", "", instructions, "")
}

func createValueWithExecution(
	execution string,
) Value {
	return createValueInternally("", "", nil, execution)
}

func createValueInternally(
	input string,
	str string,
	instructions Instructions,
	execution string,
) Value {
	out := value{
		input:        input,
		str:          str,
		instructions: instructions,
		execution:    execution,
	}

	return &out
}

// IsInput returns true if there is an input, false otherwise
func (obj *value) IsInput() bool {
	return obj.input != ""
}

// Input returns the input, if any
func (obj *value) Input() string {
	return obj.input
}

// IsString returns true if there is a string, false otherwise
func (obj *value) IsString() bool {
	return obj.str != ""
}

// String returns the string, if any
func (obj *value) String() string {
	return obj.str
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
	return obj.execution != ""
}

// Execution returns the execution, if any
func (obj *value) Execution() string {
	return obj.execution
}
