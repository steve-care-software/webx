package outputs

// NewOutputForTests creates a new output for tests
func NewOutputForTests(input []byte) Output {
	ins, err := NewBuilder().Create().WithInput(input).Now()
	if err != nil {
		panic(err)
	}

	return ins
}

// NewOutputWithExecuteForTests creates a new output with execute for tests
func NewOutputWithExecuteForTests(input []byte, execute []byte) Output {
	ins, err := NewBuilder().Create().WithInput(input).WithExecute(execute).Now()
	if err != nil {
		panic(err)
	}

	return ins
}
