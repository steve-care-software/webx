package programs

import "github.com/steve-care-software/webx/domain/cryptography/hash"

type program struct {
	hash     hash.Hash
	engine   string
	compiler []byte
	script   []byte
}

func createProgram(
	hash hash.Hash,
	engine string,
	compiler []byte,
	script []byte,
) Program {
	out := program{
		hash:     hash,
		engine:   engine,
		compiler: compiler,
		script:   script,
	}

	return &out
}

// Hash returns the hash
func (obj *program) Hash() hash.Hash {
	return obj.hash
}

// Engine returns the engine
func (obj *program) Engine() string {
	return obj.engine
}

// Compiler returns the compiler
func (obj *program) Compiler() []byte {
	return obj.compiler
}

// Script returns the script
func (obj *program) Script() []byte {
	return obj.script
}
