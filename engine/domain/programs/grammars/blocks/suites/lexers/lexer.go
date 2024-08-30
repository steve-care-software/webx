package lexers

type lexer struct {
	output []byte
	isFail bool
}

func createLexer(
	output []byte,
	isFail bool,
) Lexer {
	out := lexer{
		output: output,
		isFail: isFail,
	}

	return &out
}

// Output returns the output
func (obj *lexer) Output() []byte {
	return obj.output
}

// IsFail returns true if fail, false otherwise
func (obj *lexer) IsFail() bool {
	return obj.isFail
}
