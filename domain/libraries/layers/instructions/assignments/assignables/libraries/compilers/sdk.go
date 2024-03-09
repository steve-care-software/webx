package compilers

// Compiler represents a compiler
type Compiler interface {
	IsCompile() bool
	Compile() string
	IsDecompile() bool
	Decompile() string
}
