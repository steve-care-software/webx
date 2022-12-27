package programs

type content struct {
	pInput    *uint
	constant  []byte
	execution Application
	program   Program
}

func createContentWithInput(
	pInput *uint,
) Content {
	return createContentInternally(pInput, nil, nil, nil)
}

func createContentWithConstant(
	constant []byte,
) Content {
	return createContentInternally(nil, constant, nil, nil)
}

func createContentWithExecution(
	execution Application,
) Content {
	return createContentInternally(nil, nil, execution, nil)
}

func createContentWithProgram(
	program Program,
) Content {
	return createContentInternally(nil, nil, nil, program)
}

func createContentInternally(
	pInput *uint,
	constant []byte,
	execution Application,
	program Program,
) Content {
	out := content{
		pInput:    pInput,
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
