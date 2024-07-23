package compilers

// NewCompilerWithCompileForTests creates a new compiler with compile for tests
func NewCompilerWithCompileForTests(compile string) Compiler {
	ins, err := NewBuilder().Create().WithCompile(compile).Now()
	if err != nil {
		panic(err)
	}

	return ins
}

// NewCompilerWithDecompileForTests creates a new compiler with decompile for tests
func NewCompilerWithDecompileForTests(decompile string) Compiler {
	ins, err := NewBuilder().Create().WithDecompile(decompile).Now()
	if err != nil {
		panic(err)
	}

	return ins
}
