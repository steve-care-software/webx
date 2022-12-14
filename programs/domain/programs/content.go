package programs

type content struct {
	pInput    *uint
	value     Value
	constant  []byte
	execution Application
	program   Program
}

func createContentWithInput(
	pInput *uint,
) Content {
	return createContentInternally(pInput, nil, nil, nil, nil)
}

func createContentWithValue(
	value Value,
) Content {
	return createContentInternally(nil, value, nil, nil, nil)
}

func createContentWithConstant(
	constant []byte,
) Content {
	return createContentInternally(nil, nil, constant, nil, nil)
}

func createContentWithExecution(
	execution Application,
) Content {
	return createContentInternally(nil, nil, nil, execution, nil)
}

func createContentWithProgram(
	program Program,
) Content {
	return createContentInternally(nil, nil, nil, nil, program)
}

func createContentInternally(
	pInput *uint,
	value Value,
	constant []byte,
	execution Application,
	program Program,
) Content {
	out := content{
		pInput:    pInput,
		value:     value,
		constant:  constant,
		execution: execution,
		program:   program,
	}

	return &out
}

// IsInput returns true if input, false otherwise
func (obj *content) IsInput() bool {
	return obj.pInput != nil
}

// Input returns the input, if any
func (obj *content) Input() *uint {
	return obj.pInput
}

// IsValue returns true if value, false otherwise
func (obj *content) IsValue() bool {
	return obj.value != nil
}

// Value returns the value, if any
func (obj *content) Value() Value {
	return obj.value
}

// IsConstant returns true if []byte, false otherwise
func (obj *content) IsConstant() bool {
	return obj.constant != nil
}

// Constant returns the []byte, if any
func (obj *content) Constant() []byte {
	return obj.constant
}

// IsExecution returns true if execution, false otherwise
func (obj *content) IsExecution() bool {
	return obj.execution != nil
}

// Execution returns the execution, if any
func (obj *content) Execution() Application {
	return obj.execution
}

// IsProgram returns true if program, false otherwise
func (obj *content) IsProgram() bool {
	return obj.program != nil
}

// Program returns the program, if any
func (obj *content) Program() Program {
	return obj.program
}
