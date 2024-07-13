package compilers

import "github.com/steve-care-software/historydb/domain/hash"

type compiler struct {
	hash      hash.Hash
	compile   string
	decompile string
}

func createCompilerWithCompile(
	hash hash.Hash,
	compile string,
) Compiler {
	return createCompilerInternally(hash, compile, "")
}

func createCompilerWithDecompile(
	hash hash.Hash,
	decompile string,
) Compiler {
	return createCompilerInternally(hash, "", decompile)
}

func createCompilerInternally(
	hash hash.Hash,
	compile string,
	decompile string,
) Compiler {
	out := compiler{
		hash:      hash,
		compile:   compile,
		decompile: decompile,
	}

	return &out
}

// Hash returns the hash
func (obj *compiler) Hash() hash.Hash {
	return obj.hash
}

// IsCompile retruns true if compile, false otherwise
func (obj *compiler) IsCompile() bool {
	return obj.compile != ""
}

// Compile returns decompile, if any
func (obj *compiler) Compile() string {
	return obj.compile
}

// IsDecompile retruns true if decompile, false otherwise
func (obj *compiler) IsDecompile() bool {
	return obj.decompile != ""
}

// Decompile returns decompile, if any
func (obj *compiler) Decompile() string {
	return obj.decompile
}
