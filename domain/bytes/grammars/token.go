package grammars

type token struct {
	name   string
	block  Block
	suites Suites
}

func createToken(
	name string,
	block Block,
) Token {
	return createTokenInternally(name, block, nil)
}

func createTokenWithSuites(
	name string,
	block Block,
	suites Suites,
) Token {
	return createTokenInternally(name, block, suites)
}

func createTokenInternally(
	name string,
	block Block,
	suites Suites,
) Token {
	out := token{
		name:   name,
		block:  block,
		suites: suites,
	}

	return &out
}

// Name returns the name
func (obj *token) Name() string {
	return obj.name
}

// Block returns the block
func (obj *token) Block() Block {
	return obj.block
}

// HasSuites returns true if there is suites, false otherwise
func (obj *token) HasSuites() bool {
	return obj.suites != nil
}

// Suites returns the suites, if any
func (obj *token) Suites() Suites {
	return obj.suites
}
