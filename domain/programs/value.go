package programs

type value struct {
	input     string
	str       string
	execution Application
	program   Program
}

func createValueWithInput(
	input string,
) Value {
	return createValueInternally(input, "", nil, nil)
}

func createValueWithString(
	str string,
) Value {
	return createValueInternally("", str, nil, nil)
}

func createValueWithExecution(
	execution Application,
) Value {
	return createValueInternally("", "", execution, nil)
}

func createValueWithProgram(
	program Program,
) Value {
	return createValueInternally("", "", nil, program)
}

func createValueInternally(
	input string,
	str string,
	execution Application,
	program Program,
) Value {
	out := value{
		input:     input,
		str:       str,
		execution: execution,
		program:   program,
	}

	return &out
}

// IsInput returns true if input, false otherwise
func (obj *value) IsInput() bool {
	return obj.input != ""
}

// Input returns the input, if any
func (obj *value) Input() string {
	return obj.input
}

// IsString returns true if string, false otherwise
func (obj *value) IsString() bool {
	return obj.str != ""
}

// String returns the string, if any
func (obj *value) String() string {
	return obj.str
}

// IsExecution returns true if execution, false otherwise
func (obj *value) IsExecution() bool {
	return obj.execution != nil
}

// Execution returns the execution, if any
func (obj *value) Execution() Application {
	return obj.execution
}

// IsProgram returns true if program, false otherwise
func (obj *value) IsProgram() bool {
	return obj.program != nil
}

// Program returns the program, if any
func (obj *value) Program() Program {
	return obj.program
}
