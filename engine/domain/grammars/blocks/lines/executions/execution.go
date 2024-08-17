package executions

type execution struct {
	tokens []string
	fnFLag uint16
}

func createExecution(
	tokens []string,
	fnFLag uint16,
) Execution {
	out := execution{
		tokens: tokens,
		fnFLag: fnFLag,
	}

	return &out
}

// Tokens returns the tokens
func (obj *execution) Tokens() []string {
	return obj.tokens
}

// FuncFlag returns the func flag
func (obj *execution) FuncFlag() uint16 {
	return obj.fnFLag
}
