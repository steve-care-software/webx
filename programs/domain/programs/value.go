package programs

type value struct {
	input      []byte
	assignment Assignment
	constant   []byte
	execution  Application
	program    Program
}

func createValueWithInput(
	input []byte,
) Value {
	return createValueInternally(input, nil, nil, nil, nil)
}

func createValueWithAssignment(
	assignment Assignment,
) Value {
	return createValueInternally(nil, assignment, nil, nil, nil)
}

func createValueWithConstant(
	constant []byte,
) Value {
	return createValueInternally(nil, nil, constant, nil, nil)
}

func createValueWithExecution(
	execution Application,
) Value {
	return createValueInternally(nil, nil, nil, execution, nil)
}

func createValueWithProgram(
	program Program,
) Value {
	return createValueInternally(nil, nil, nil, nil, program)
}

func createValueInternally(
	input []byte,
	assignment Assignment,
	constant []byte,
	execution Application,
	program Program,
) Value {
	out := value{
		input:      input,
		assignment: assignment,
		constant:   constant,
		execution:  execution,
		program:    program,
	}

	return &out
}

// IsInput returns true if input, false otherwise
func (obj *value) IsInput() bool {
	return obj.input != nil
}

// Input returns the input, if any
func (obj *value) Input() []byte {
	return obj.input
}

// IsAssignment returns true if assignment, false otherwise
func (obj *value) IsAssignment() bool {
	return obj.assignment != nil
}

// Assignment returns the assignment, if any
func (obj *value) Assignment() Assignment {
	return obj.assignment
}

// IsConstant returns true if []byte, false otherwise
func (obj *value) IsConstant() bool {
	return obj.constant != nil
}

// Constant returns the []byte, if any
func (obj *value) Constant() []byte {
	return obj.constant
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
