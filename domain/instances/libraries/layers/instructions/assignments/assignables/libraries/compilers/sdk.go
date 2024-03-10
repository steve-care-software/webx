package compilers

// Builder represents the compiler application
type Builder interface {
	Create() Builder
	WithCompile(compile string) Builder
	WithDecompile(decompile string) Builder
	Now() (Compiler, error)
}

// Compiler represents a compiler
type Compiler interface {
	IsCompile() bool
	Compile() string
	IsDecompile() bool
	Decompile() string
}
